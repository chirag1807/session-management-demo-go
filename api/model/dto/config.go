package dto

type Config struct {
	DATABASE Database `json:"database"`
}

type Database struct {
	DATABASE_USERNAME string `json:"db_username"`
	DATABASE_PASSWORD string `json:"db_password"`
	DATABASE_PORT     string `json:"db_port"`
	DATABASE_NAME     string `json:"db_name"`
	DATABASE_SSLMODE  string `json:"db_sslmode"`
}
