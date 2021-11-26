package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lmiedzinski/shiny-barnacle/docs"
	"github.com/lmiedzinski/shiny-barnacle/internal/usecase"
	"github.com/lmiedzinski/shiny-barnacle/pkg/logger"
)

// @title       Shiny Barnacle API
// @description Golang microservice example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, logger logger.Interface, uc usecase.HelloWorld) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	handler.GET("/swagger", redirectPath(handler, "/swagger/index.html"))
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Healthcheck
	handler.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	// Routers
	h := handler.Group("/v1")
	{
		newHelloWorldRoutes(h, uc, logger)
	}
}

func redirectPath(r *gin.Engine, newPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Path = newPath
		r.HandleContext(c)
	}
}
