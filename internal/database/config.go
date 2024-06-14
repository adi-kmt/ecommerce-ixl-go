package database

type DbConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
}

// Create a new db config type
func NewDbConfig(user string, password string, name string, host string, port string) *DbConfig {
	return &DbConfig{
		dbUser:     user,
		dbPassword: password,
		dbName:     name,
		dbHost:     host,
		dbPort:     port,
	}
}
