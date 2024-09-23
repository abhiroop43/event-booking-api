package routes

import (
	"abhiroopsanta.dev/event-booking-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.GET("/my-events", getRegisteredEvents)
	authenticated.POST("/events/:id/register", registerForEvents)
	authenticated.DELETE("/events/:id/cancel", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
