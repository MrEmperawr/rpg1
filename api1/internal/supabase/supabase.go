package supabase

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	supa "github.com/supabase-community/supabase-go"
)

var (
	client *supa.Client
	pool   *pgxpool.Pool
)

func Connect(supabaseURL, supabaseKey string) error {
	if supabaseURL == "" || supabaseKey == "" {
		return fmt.Errorf("supabase URL and supabase anon key are required")
	}

	var err error
	client, err = supa.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return fmt.Errorf("failed to create Supabase client: %w", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("database URL environment variable is required")
	}

	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("failed to create database pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func GetClient() *supa.Client {
	return client
}

func GetPool() *pgxpool.Pool {
	return pool
}

func Close() {
	if pool != nil {
		pool.Close()
	}
}

type AuthUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func ValidateToken(ctx context.Context, token string) (*AuthUser, error) {
	if client == nil {
		return nil, fmt.Errorf("supabase client not initialized")
	}

	return nil, fmt.Errorf("use middleware.ValidateToken instead")
}

func GetCurrentUser(ctx context.Context) (*AuthUser, error) {
	if client == nil {
		return nil, fmt.Errorf("supabase client not initialized")
	}

	user, err := client.Auth.GetUser()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	return &AuthUser{
		ID:    user.ID.String(),
		Email: user.Email,
	}, nil
}
