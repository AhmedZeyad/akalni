package routes

import (
	"github.com/ba7rIbrahim/Akalni/auth"
	"github.com/ba7rIbrahim/Akalni/client"
	"github.com/ba7rIbrahim/Akalni/config"
	"github.com/jmoiron/sqlx"
)

func LoadRoutes(conf *config.Config, db *sqlx.DB, jwtService *auth.JTWSevice) {
	// Load routes here

	authRepo := auth.NewAuthRepo(db)
	authService := auth.NewAuthService(authRepo, jwtService)
	authHandler := auth.NewAuthHandler(*authService)
	AddNonAuthRoutes("POST", "/singup", authHandler.Create)
	AddNonAuthRoutes("POST", "/login", authHandler.Login)
	AddNonAuthRoutes("POST", "/Refresh", authHandler.Refresh)
	clientRepo := client.NewClientRepo(db)
	clientService := client.NewClientService(clientRepo)
	clientHandler := client.NewClientHandler(*clientService)
	AddAuthRoutes("GET", "/profile", clientHandler.GetProfile)

}
