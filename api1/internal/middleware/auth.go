package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mremperor-atwork/rpg1/api1/internal/supabase"
)

type AuthUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Supabase JWT claims structure
type SupabaseClaims struct {
	Aud         string `json:"aud"`
	Exp         int64  `json:"exp"`
	Sub         string `json:"sub"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	AppMetadata struct {
		Provider string `json:"provider"`
	} `json:"app_metadata"`
	UserMetadata map[string]interface{} `json:"user_metadata"`
	Iat          int64                  `json:"iat"`
	Iss          string                 `json:"iss"`
}

// Implement jwt.Claims interface
func (c *SupabaseClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Exp, 0)), nil
}

func (c *SupabaseClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (c *SupabaseClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Iat, 0)), nil
}

func (c *SupabaseClaims) GetIssuer() (string, error) {
	return c.Iss, nil
}

func (c *SupabaseClaims) GetSubject() (string, error) {
	return c.Sub, nil
}

func (c *SupabaseClaims) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{c.Aud}, nil
}

// AutoRefreshAuthMiddleware automatically refreshes tokens when they're close to expiring
func AutoRefreshAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "bearer token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse token to check expiration
		claims, err := parseTokenClaims(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
			c.Abort()
			return
		}

		// Check if token is expired
		if claims.Exp < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			c.Abort()
			return
		}

		// Check if token will expire soon (within 5 minutes)
		expiresIn := claims.Exp - time.Now().Unix()
		if expiresIn < 300 { // 5 minutes = 300 seconds
			// Try to refresh the token
			refreshToken := c.GetHeader("X-Refresh-Token")
			if refreshToken != "" {
				newTokens, err := refreshAccessToken(refreshToken)
				if err == nil {
					// Set new tokens in response headers
					c.Header("X-New-Access-Token", newTokens.AccessToken)
					c.Header("X-New-Refresh-Token", newTokens.RefreshToken)
					c.Header("X-Token-Expires-In", fmt.Sprintf("%d", newTokens.ExpiresIn))

					// Update the token for this request
					tokenString = newTokens.AccessToken

					// Re-parse claims with new token
					claims, _ = parseTokenClaims(tokenString)
				}
			}
		}

		// Validate the token (either original or refreshed)
		user, err := validateSupabaseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// AuthMiddleware is the standard auth middleware without auto-refresh
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "bearer token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		user, err := validateSupabaseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// TokenResponse represents the response from a token refresh
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// refreshAccessToken attempts to refresh an access token using the refresh token
func refreshAccessToken(refreshToken string) (*TokenResponse, error) {
	supabaseClient := supabase.GetClient()
	if supabaseClient == nil {
		return nil, fmt.Errorf("supabase client not initialized")
	}

	authResponse, err := supabaseClient.Auth.RefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return &TokenResponse{
		AccessToken:  authResponse.AccessToken,
		RefreshToken: authResponse.RefreshToken,
		ExpiresIn:    int64(authResponse.ExpiresIn),
	}, nil
}

// parseTokenClaims parses JWT token and returns claims without validation
func parseTokenClaims(tokenString string) (*SupabaseClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &SupabaseClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*SupabaseClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func validateSupabaseToken(tokenString string) (*AuthUser, error) {
	claims, err := parseTokenClaims(tokenString)
	if err != nil {
		return nil, err
	}

	// Check if token is expired
	if claims.Exp < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	// Validate required fields
	if claims.Sub == "" {
		return nil, fmt.Errorf("missing user ID in token")
	}

	if claims.Email == "" {
		return nil, fmt.Errorf("missing email in token")
	}

	// Validate issuer (should be your Supabase project)
	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL != "" {
		expectedIssuer := supabaseURL + "/auth/v1"
		if claims.Iss != expectedIssuer {
			return nil, fmt.Errorf("invalid token issuer")
		}
	}

	return &AuthUser{
		ID:    claims.Sub,
		Email: claims.Email,
	}, nil
}

// validateTokenWithSignature validates JWT token with signature verification
// This is the proper way to validate tokens in production
func validateTokenWithSignature(tokenString string) (*AuthUser, error) {
	// Get Supabase JWT secret from environment
	jwtSecret := os.Getenv("SUPABASE_JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("supabase JWT secret not configured")
	}

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &SupabaseClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*SupabaseClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return &AuthUser{
		ID:    claims.Sub,
		Email: claims.Email,
	}, nil
}

func GetUserFromContext(c *gin.Context) (*AuthUser, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	authUser, ok := user.(*AuthUser)
	return authUser, ok
}

func RequireAuth() gin.HandlerFunc {
	return AuthMiddleware()
}

// RequireAuthWithAutoRefresh uses the auto-refresh middleware
func RequireAuthWithAutoRefresh() gin.HandlerFunc {
	return AutoRefreshAuthMiddleware()
}

func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if user, err := validateSupabaseToken(tokenString); err == nil {
				c.Set("user", user)
			}
		}
		c.Next()
	}
}
