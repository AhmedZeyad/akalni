package routes

import (
	"github.com/AhmedZeyad/Akalni/categories"
	"github.com/AhmedZeyad/Akalni/client"
	"github.com/AhmedZeyad/Akalni/config"
	"github.com/AhmedZeyad/Akalni/middleware"
	"github.com/AhmedZeyad/Akalni/order"
	"github.com/AhmedZeyad/Akalni/products"
	"github.com/AhmedZeyad/Akalni/restaurant"
	"github.com/AhmedZeyad/Akalni/users"
	"github.com/jmoiron/sqlx"
)

func LoadRoutes(conf *config.Config, db *sqlx.DB, jwtService *middleware.JWTService) {
	// Load routes here

	clientRepo := client.NewClientRepo(db)
	clientService := client.NewClientService(clientRepo, jwtService)
	clientHandler := client.NewClientHandler(*clientService, *conf)
	AddNonAuthRoutes("POST", "/singup", clientHandler.Create)
	AddNonAuthRoutes("POST", "/login", clientHandler.Login)
	AddNonAuthRoutes("POST", "/refresh", clientHandler.Refresh)
	// AddNonAuthRoutes("POST", "/clients", clientHandler.Create)
	AddAuthRoutes("GET", "/profile", clientHandler.GetProfile)
	AddAuthRoutes("POST", "/email/verification/send", clientHandler.SendOtp)
	AddAuthRoutes("POST", "/email/verification/resend", clientHandler.ResendOtp)
	AddAuthRoutes("POST", "/email/verification/verify", clientHandler.VerifyOtp)

	AddAuthRoutes("POST", "/email/update/send/otp", clientHandler.ResendOtpForUpdateEmail)
	AddAuthRoutes("POST", "/email/update/resend/otp", clientHandler.ResendOtpForUpdateEmail)
	AddAuthRoutes("POST", "/email/update/verify/otp", clientHandler.VerifyUpdateEmail)

	//INFO Admin routes

	userRepo := users.NewUserRepo(db)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(*userService, jwtService)
	AddAdminNonAuthRoutes("POST", "/users", userHandler.CreateUser)
	AddAdminNonAuthRoutes("POST", "/login", userHandler.Login)
	AddAdminNonAuthRoutes("POST", "/resetpassword", userHandler.ResetPassword)
	// users.InitUserRoutes(db)
	// AddAdminNonAuthRoutes("POST", "/signup")
	// INFO orders LoadRoutes
	orderRepo := order.NewOrderRepo(db)
	orderService := order.NewOrderService(orderRepo)
	orderHandler := order.NewOrderHandler(orderService)
	// TODO:fix token check and make it auth routes
	AddAdminNonAuthRoutes("GET", "/orders", orderHandler.GetOrder)
	AddAdminNonAuthRoutes("PUT", "/orders/status", orderHandler.UpdateOrderStatus)

	restaurantRepo := restaurant.NewRestaurantRepo(db)
	restaurantService := restaurant.NewRestaurantService(restaurantRepo)
	restaurantHandler := restaurant.NewRestaurantHandler(restaurantService)
	// NOTE: Admin routes
	AddAdminRoutes("GET", "/restaurants", restaurantHandler.SearchRestaurant)
	AddAdminRoutes("GET", "/restaurantsbyid", restaurantHandler.GetRestaurantById)
	AddAdminRoutes("POST", "/restaurants", restaurantHandler.CreateRestaurant)
	AddAdminRoutes("PUT", "/restaurants", restaurantHandler.UpdateRestaurant)
	AddAdminRoutes("PUT", "/restaurants/status", restaurantHandler.UpdateRestaurantStatus)

	// NOTE: client routes
	AddNonAuthRoutes("GET", "/restaurantsbyid", restaurantHandler.GetRestaurantById) //TODO check if it match the todo
	AddNonAuthRoutes("GET", "/restaurants", restaurantHandler.GetActiveRestaurant)

	// TODO
	// AddAuthRoutes("POST", "orders",)
	// AddAuthRoutes("GET", "orders",)
	// AddAuthRoutes("GET", "orders/active",)
	// product CRUD
	categoryRepo := categories.NewRepo(db)
	categoryService := categories.NewCategoriesService(categoryRepo)
	categoryHandler := categories.NewCategoriesHandler(categoryService)
	AddAdminRoutes("GET", "categories", categoryHandler.GetCategories)
	AddAdminRoutes("POST", "categories", categoryHandler.CreateCategory)
	AddAdminRoutes("PUT", "categories", categoryHandler.UpdateCategory)
	AddAdminRoutes("DELETE", "categories", categoryHandler.DeleteCategory)
	AddAdminRoutes("PUT", "categories/status", categoryHandler.UpdateCategoryStatus)

	productRepo := products.NewRepo(db)
	productService := products.NewProductsService(productRepo)
	productHandler := products.NewProductHandler(productService)
	AddAdminRoutes("GET", "products", productHandler.GetProduct)
	AddAdminRoutes("POST", "products", productHandler.AddProduct)
	AddAdminRoutes("PUT", "products", productHandler.UpdateProduct)
	AddAdminRoutes("DELETE", "products", productHandler.DeleteProduct)
	AddAdminRoutes("PUT", "products/status", productHandler.UpdateProductStatus)

}
