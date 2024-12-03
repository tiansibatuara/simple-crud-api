package config

import (
	"database/sql"
	"fmt"
	"simple-crud-api/helper"

	"github.com/rs/zerolog/log"
)

const (
	host = "127.0.0.1"
	port = "5432"
	user = "tiansibatuara"
	password = "dian123"
	dbName = "postgres"
)

func DatabaseConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database!")

	return db
}