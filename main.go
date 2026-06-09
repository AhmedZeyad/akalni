package main

import (
	"github.com/AhmedZeyad/Akalni/auth"
	"github.com/AhmedZeyad/Akalni/config"
	"github.com/AhmedZeyad/Akalni/database"
	"github.com/AhmedZeyad/Akalni/logger"
	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/routes"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/jmoiron/sqlx"
)

func main() {
	conf := config.LoadConfig()
	shared.Conf = conf
	db := database.Connect(*conf)
	logger.Init(conf)
	logger.Log.Info("server start ", "port", conf.Port)
	// jwtService := auth.NewJWTService(conf.JWTExpire, conf.RefreshJWTExpire, conf.JWTSecret)
	jwtService := middleware.NewJwtService(conf.JWTExpire, conf.RefreshJWTExpire, conf.JWTSecret)
	auth.SendOTP(conf, "AhmedZeyad.AZ@proton.me", "شكرا استاذ بكر")
	routes.LoadRoutes(conf, db, jwtService)
	routes.InitRouter(conf, jwtService)
	// routes.RegeserRoutes()
	defer onDistroy(conf, db)

}

func onDistroy(conf *config.Config, db *sqlx.DB) {
	db.Close()
	logger.Log.Info("server stop ", "port", conf.Port)
}
