package main

import (
	"github.com/ba7rIbrahim/Akalni/auth"
	"github.com/ba7rIbrahim/Akalni/config"
	"github.com/ba7rIbrahim/Akalni/database"
	"github.com/ba7rIbrahim/Akalni/logger"
	"github.com/ba7rIbrahim/Akalni/routes"
	"github.com/jmoiron/sqlx"
)

func main() {
	conf := config.LoadConfig()
	db := database.Connect(*conf)
	logger.Init(conf)
	logger.Log.Info("server start ", "port", conf.Port)
	jwtService := auth.NewJWTService(conf.JWTExpire, conf.RefreshJWTExpire, conf.JWTSecret)

	routes.LoadRoutes(conf, db, jwtService)
	routes.InitRouter(conf, jwtService)
	// routes.RegeserRoutes()
	defer onDistroy(conf, db)

}

func onDistroy(conf *config.Config, db *sqlx.DB) {
	db.Close()
	logger.Log.Info("server stop ", "port", conf.Port)
}
