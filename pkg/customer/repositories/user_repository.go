package user_repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "gituh.com/adi-kmt/ecommerce-ixl-go/db/sqlc"
)

type UserRepository struct {
	q    *db.Queries
	conn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q:    db.New(conn),
		conn: conn,
	}
}
