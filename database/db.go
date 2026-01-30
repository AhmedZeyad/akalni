package database

import (
	"fmt"
	"log"

	"github.com/ba7rIbrahim/Akalni/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(config config.Config) *sqlx.DB {
	connectionsString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sqlx.Connect("pgx", connectionsString)
	if err != nil {
		log.Println(err)
	}
	// err = db.Ping()
	// if err != nil {
	// 	log.Println(err)
	// }
	log.Println("db connectedxx")
	return db
}
