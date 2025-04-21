package connections

import (
	"fmt"
	"subscriptions/internal/app/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connections holds the database connection and HTTP client.
type Connections struct {
	DB         *gorm.DB
	HTTPClient *echo.Echo
}

// Close provides a cleanup method for any active connections.
func (c *Connections) Close() {
	// TODO: Implement cleanup if needed.
}

// BuildPostgresDSN creates a DSN string for the PostgreSQL connection.
func BuildPostgresDSN(cfg config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Almaty",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	)
}

// newHTTPClient initializes and configures the Echo HTTP client.
func newHTTPClient() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	return e
}

// New initializes the Connections by setting up the database connection and HTTP client.
func New(cfg *config.Config) (*Connections, error) {
	dsn := BuildPostgresDSN(cfg.DB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Connections{
		DB:         db,
		HTTPClient: newHTTPClient(),
	}, nil
}
