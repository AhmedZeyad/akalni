package routes

import (
	"github.com/AhmedZeyad/Akalni/auth"
	"github.com/AhmedZeyad/Akalni/client"
	"github.com/AhmedZeyad/Akalni/config"
	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/order"
	"github.com/AhmedZeyad/Akalni/users"
	"github.com/jmoiron/sqlx"
)

func LoadRoutes(conf *config.Config, db *sqlx.DB, jwtService *middleware.JWTService) {
	// Load routes here

	authRepo := auth.NewAuthRepo(db)
	authService := auth.NewAuthService(authRepo, jwtService)
	authHandler := auth.NewAuthHandler(*authService)
	AddNonAuthRoutes("POST", "/singup", authHandler.Create)
	AddNonAuthRoutes("POST", "/login", authHandler.Login)
	AddNonAuthRoutes("POST", "/Refresh", authHandler.Refresh)
	AddNonAuthRoutes("POST", "/clients", authHandler.Create)
	clientRepo := client.NewClientRepo(db)
	clientService := client.NewClientService(clientRepo)
	clientHandler := client.NewClientHandler(*clientService)
	AddAuthRoutes("GET", "/profile", clientHandler.GetProfile)

	//INFO Admin routes

	userRepo := users.NewUserRepo(db)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(*userService, jwtService)
	AddAdminNonAuthRoutes("POST", "/users", userHandler.CreateUser)
	AddAdminNonAuthRoutes("POST", "/users/login", userHandler.Login)
	AddAdminNonAuthRoutes("POST", "/users/reset-password", userHandler.ResetPassword)
	// users.InitUserRoutes(db)
	// AddAdminNonAuthRoutes("POST", "/signup")
	// INFO orders LoadRoutes
	orderRepo := order.NewOrderRepo(db)
	orderService := order.NewOrderService(orderRepo)
	orderHandler := order.NewOrderHandler(orderService)
	// TODO:fix token check and make it auth routes
	AddAdminNonAuthRoutes("GET", "/orders", orderHandler.GetOrder)
	AddAdminNonAuthRoutes("PUT", "/orders/status", orderHandler.UpdateOrderStatus)
}
