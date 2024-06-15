package injection

import (
	"os"

	"gituh.com/adi-kmt/ecommerce-ixl-go/internal/database"
	admin_repositories "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/repositories"
	admin_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/admin/services"
	user_repositories "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/repositories"
	user_services "gituh.com/adi-kmt/ecommerce-ixl-go/pkg/customer/services"
)

func InjectDependencies() (*user_services.UserService, *admin_services.AdminService) {
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
	db := database.InitPool(dbConfig)

	userRepository := user_repositories.NewUserRepository(db.DbPool, db.DBQueries)
	adminRepository := admin_repositories.NewAdminRepository(db.DbPool, db.DBQueries)

	userServices := user_services.NewUserService(userRepository)
	adminServices := admin_services.NewAdminService(adminRepository)

	return userServices, adminServices
}
