package database

import (
	"fmt"
	"log/slog"

	"github.com/AhmedZeyad/Akalni/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(config config.Config) *sqlx.DB {
	connectionsString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sqlx.Connect("pgx", connectionsString)
	if err != nil {
		slog.Error("failed to connect to db", "error", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		slog.Error("failed to ping db", "error", err)
		panic(err)
	}
	slog.Info("db connected")

	return db
}
