package api

import (
	"cycling-tracker-server/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpApi struct {
	router      *gin.Engine
	userService service.UserService
}

func (api *HttpApi) Init() *HttpApi {
	api.router = gin.Default()
	api.userService = service.UserService{}

	api.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	api.bindPing()
	api.bindEndpoints()

	return api
}

func (api *HttpApi) StartServer(addr string) {
	api.router.Run(addr)
}

func (api *HttpApi) GetRouter() *gin.Engine {
	return api.router
}

func (api *HttpApi) bindEndpoints() {
	api.router.POST("/api/v1/login", api.userService.LoginUser)
}

func (api *HttpApi) bindPing() {
	api.router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
