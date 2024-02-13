package db

import (
	"sessionmanagement/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DBConnection() (conn *pgx.Conn, err error) {
	DATABASE_URL := "postgresql://" + config.DatabaseConfig.DATABASE_USERNAME + ":" + config.DatabaseConfig.DATABASE_PASSWORD + "@127.0.0.1:" + config.DatabaseConfig.DATABASE_PORT + "/" + config.DatabaseConfig.DATABASE_NAME + "?sslmode=" + config.DatabaseConfig.DATABASE_SSLMODE
	config, err := pgx.ParseConfig(DATABASE_URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	conn, err = pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return conn, nil
}
