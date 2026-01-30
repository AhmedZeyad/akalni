package routes

import (
	"github.com/ba7rIbrahim/Akalni/config"
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

func InitRouter(conf *config.Config) {
	// engin:=
	var mode string
	if conf.ISDev == "true" && conf.ISLocal == "true" {
		mode = gin.DebugMode
	} else {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	engine := gin.New()
	engine.GET("/kaithheathcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	group := engine.Group("/api")
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello leapcell",
		})
	})
	RegeserRoutes(group)
	engine.Run("0.0.0.0:" + conf.Port)

}
func AddAuthRoutes(Method, path string, handler gin.HandlerFunc, Roles []string) {
	routes = append(routes, Routes{Method, path, handler, true, Roles})
}
func AddNonAuthRoutes(Method, path string, handler gin.HandlerFunc) {
	routes = append(routes, Routes{Method, path, handler, false, nil})
}
func RegeserRoutes(engine *gin.RouterGroup) {
	for _, route := range routes {
		// TODO: Implement route registration logic
		// TODO: add middleware
		engine.Handle(route.Method, route.Path, route.Handler)
	}

}
