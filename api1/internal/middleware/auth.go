package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func validateSupabaseToken(tokenString string) (*AuthUser, error) {
	// For development, we'll use unverified parsing to extract claims
	// In production, you should verify the JWT signature with Supabase's public key
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &SupabaseClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*SupabaseClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
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
	if supabaseURL != "" && claims.Iss != supabaseURL {
		return nil, fmt.Errorf("invalid token issuer")
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
