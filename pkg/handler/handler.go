package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/pkg/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	users := router.Group("/users")
	{
		users.GET("/", h.getAllUsers)
		users.GET("/:id", h.getUserById)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
	}

	orders := router.Group("/orders")
	{
		orders.GET("/", h.getAllOrders)
		orders.GET("/:id", h.getOrderById)
		orders.POST("/", h.createOrder)
		orders.PUT("/:id", h.updateOrder)
		orders.DELETE("/:id", h.deleteOrder)
	}

	return router
}
