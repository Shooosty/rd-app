package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shooosty/rd-app/pkg/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://rd-cabinet-7mds4.ondigitalocean.app",
			"https://rd-cabinet-7mds4.ondigitalocean.app/", "https://lk.rhinodesign.ru", "https://lk.rhinodesign.ru/"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE", "HEAD"},
		ExposeHeaders: []string{"X-Next-Page", "X-Page", "X-Per-Page", "X-Prev-Page", "X-Total", "X-Total-Pages"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization",
			"Access-Control-Allow-Headers", "X-Requested-With", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-up-employee", h.signUpEmployee)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/reset-password", h.resetPassword)
	}

	users := router.Group("/users")
	{
		users.GET("/", h.getAllUsers)
		users.GET("/:id", h.getUserById)
		users.GET("/:id/orders", h.getAllForUserOrders)
		users.PUT("/:id", h.updateUser)
		users.PUT("/:id/change-password", h.changePassword)
		users.DELETE("/:id", h.deleteUser)
	}

	orders := router.Group("/orders")
	{
		orders.GET("/", h.getAllOrders)
		orders.GET("/:id", h.getOrderById)
		orders.GET("/:id/persons", h.getAllByOrderId)
		orders.POST("/", h.createOrder)
		orders.PUT("/:id", h.updateOrder)
		orders.DELETE("/:id", h.deleteOrder)
	}

	persons := router.Group("/persons")
	{
		persons.GET("/", h.getAllPersons)
		persons.POST("/", h.createPerson)
		persons.PUT("/:id", h.updatePerson)
		persons.DELETE("/:id", h.deletePerson)
	}

	photos := router.Group("/photos")
	{
		photos.GET("/", h.getAllPhotos)
		photos.POST("/", h.createPhoto)
		photos.DELETE("/:id", h.deletePhoto)
	}

	photographers := router.Group("/photographers")
	{
		photographers.GET("/:id/orders", h.getAllForPhotographerOrders)
	}

	designers := router.Group("/designers")
	{
		designers.GET("/:id/orders", h.getAllForDesignerOrders)
	}

	mail := router.Group("/mail")
	{
		mail.POST("/new-order", h.sendNewOrderMessage)
	}

	return router
}
