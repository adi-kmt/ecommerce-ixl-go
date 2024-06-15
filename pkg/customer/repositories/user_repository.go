package user_repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
)

type UserRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool, q *db.Queries) *UserRepository {
	return &UserRepository{
		q:    q,
		conn: conn,
	}
}
