package handler

import (
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
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:3000"},
	//	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE", "HEAD"},
	//	ExposeHeaders:    []string{"X-Next-Page", "X-Page", "X-Per-Page", "X-Prev-Page", "X-Total", "X-Total-Pages"},
	//	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

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
