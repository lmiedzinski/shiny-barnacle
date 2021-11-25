package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lmiedzinski/shiny-barnacle/internal/entity"
	"github.com/lmiedzinski/shiny-barnacle/internal/usecase"
	"github.com/lmiedzinski/shiny-barnacle/pkg/logger"
)

type helloWorldRoutes struct {
	huc    usecase.HelloWorld
	logger logger.Interface
}

func newHelloWorldRoutes(handler *gin.RouterGroup, huc usecase.HelloWorld, logger logger.Interface) {
	routes := &helloWorldRoutes{huc, logger}

	h := handler.Group("/hello-world")
	{
		h.GET("/", routes.getHelloWorld)
	}
}

type getHelloWorldsResponse struct {
	HelloWorldMessages []entity.HelloWorldDto `json:"helloWorldMessages"`
}

// @Summary     Show hello world messages
// @Description Show all known hello world messages
// @ID          getHelloWorlds
// @Tags  	    helloworld
// @Accept      json
// @Produce     json
// @Success     200 {object} getHelloWorldsResponse
// @Failure     500 {object} response
// @Router      /hello-world [get]
func (r *helloWorldRoutes) getHelloWorld(c *gin.Context) {
	helloWorldMessages, err := r.huc.GetHelloWorld(c.Request.Context())
	if err != nil {
		r.logger.Error(err)
		errorResponse(c, http.StatusInternalServerError, "There was a problem while processing your request")

		return
	}

	c.JSON(http.StatusOK, getHelloWorldsResponse{helloWorldMessages})
}
