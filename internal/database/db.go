package database

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
)

type DbConn struct {
	DbPool    *pgxpool.Pool
	DBQueries *db.Queries
}

// InitPool initializes the database connection pool and runs migrations.
func InitPool(config *DbConfig) *DbConn {
	connectionString := fmt.Sprintf("postgresql://%s:%s@127.0.0.1:%s/%s?sslmode=disable",
		config.dbUser, config.dbPassword, config.dbPort, config.dbName)

	ctx := context.Background()

	// Initialize the connection pool
	d, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Errorf("error opening database: %v", err)
	}

	// Verify the connection
	if err := d.Ping(ctx); err != nil {
		log.Errorf("error pinging database: %v", err)
	}

	// Return the connection pool pointer singleton connection
	return &DbConn{
		DbPool:    d,
		DBQueries: db.New(d),
	}
}

// ClosePool closes the database connection pool.
func (db *DbConn) ClosePool() error {
	if db.DbPool != nil {
		return db.ClosePool()
	}
	return nil
}
