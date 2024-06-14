db-migrate:
	migrate -database postgres://postgres:password@127.0.0.1:5432/ecommerce_db?sslmode=disable -path=db/migrations/ up 1

inspect-db:
	 docker exec -it ecommerce_postgres psql -U postgres -W ecommerce_db
	#  password is password