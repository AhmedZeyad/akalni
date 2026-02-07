package routes

import (
	"log/slog"
	"net/http"

	"github.com/ba7rIbrahim/Akalni/auth"
	"github.com/ba7rIbrahim/Akalni/config"
	"github.com/ba7rIbrahim/Akalni/logger"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
	Auth    bool
	Roles   []string
}

var routes = []Routes{}

func InitRouter(conf *config.Config, jwtService *auth.JTWSevice) {
	// engin:=
	var mode string
	if conf.ISDev == "true" && conf.ISLocal == "true" {
		mode = gin.DebugMode
	} else {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	engine := gin.New()
	engine.Use(func(c *gin.Context) {
		logger.Log.Debug("Request info ", "method", c.Request.Method, "path", c.Request.URL.Path)
		c.Next()
	})
	engine.GET("/kaithheathcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	group := engine.Group("/api")

	group.HEAD("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "im good ",
		})

	})

	RegeserRoutes(group, jwtService)

	engine.Run("0.0.0.0:" + conf.Port)

}
func AddAuthRoutes(Method, path string, handler gin.HandlerFunc, Roles ...string) {
	routes = append(routes, Routes{Method, path, handler, true, Roles})
}
func AddNonAuthRoutes(Method, path string, handler gin.HandlerFunc) {
	routes = append(routes, Routes{Method, path, handler, false, nil})
}
func RegeserRoutes(engine *gin.RouterGroup, jwtService *auth.JTWSevice) {
	for _, route := range routes {
		// TODO: Implement route registration logic
		// TODO: add middleware
		if route.Auth {
			engine.Handle(route.Method, route.Path, JWTMiddleware(jwtService), route.Handler)

		} else {

			engine.Handle(route.Method, route.Path, route.Handler)
		}
	}

}

// todo implement jwt middleware
func JWTMiddleware(jwtService *auth.JTWSevice) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token
		token := ctx.Request.Header.Get("Authorization")

		claims, err := jwtService.TokenVerify(token)
		if err != nil {
			slog.Error("error on verify token", "error", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("client", claims)
		ctx.Next()
		// TokenVerify
		// next or abort
	}

}
