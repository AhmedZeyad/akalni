package routes

import (
	"github.com/ba7rIbrahim/Akalni/config"
	"github.com/ba7rIbrahim/Akalni/handlers"
	"github.com/ba7rIbrahim/Akalni/middlewares"
	"github.com/ba7rIbrahim/Akalni/repo"
	"github.com/ba7rIbrahim/Akalni/services"
	"github.com/jmoiron/sqlx"
)

func LoadRoutes(conf *config.Config, db *sqlx.DB) {
	// Load routes here
	jwtService := middlewares.NewJWTService(conf.JWTExpire, conf.RefreshJWTExpire, conf.JWTSecret)
	clientRepo := repo.NewAuthRepo(db)
	clientService := services.NewClientService(clientRepo, jwtService)
	clientHandler := handlers.NewClientHandler(*clientService)
	AddNonAuthRoutes("POST", "/singup", clientHandler.Create)
}
