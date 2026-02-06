package routes

import (
	"github.com/ba7rIbrahim/Akalni/auth"
	"github.com/ba7rIbrahim/Akalni/config"
	"github.com/jmoiron/sqlx"
)

func LoadRoutes(conf *config.Config, db *sqlx.DB) {
	// Load routes here
	jwtService := auth.NewJWTService(conf.JWTExpire, conf.RefreshJWTExpire, conf.JWTSecret)
	clientRepo := auth.NewAuthRepo(db)
	clientService := auth.NewClientService(clientRepo, jwtService)
	clientHandler := auth.NewClientHandler(*clientService)
	AddNonAuthRoutes("POST", "/singup", clientHandler.Create)
	AddNonAuthRoutes("POST", "/login", clientHandler.Login)
}
