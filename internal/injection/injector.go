package injection

import (
	"os"

	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/database"
)

func InjectDependencies() {
	db_user, isErr := os.LookupEnv("DB_USER")
	if !isErr {
		db_user = "postgres"
	}

	db_password, isErr := os.LookupEnv("DB_PASSWORD")
	if !isErr {
		db_password = "password"
	}

	db_name, isErr := os.LookupEnv("DB_NAME")
	if !isErr {
		db_name = "ecommerce_db"
	}

	db_host, isErr := os.LookupEnv("DB_HOST")
	if !isErr {
		db_host = "127.0.0.1"
	}

	db_port, isErr := os.LookupEnv("DB_PORT")
	if !isErr {
		db_port = "5432"
	}

	dbConfig := database.NewDbConfig(
		db_user,
		db_password,
		db_name,
		db_host,
		db_port,
	)
	database.InitPool(dbConfig)
}
