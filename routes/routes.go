package routes

import (
	"dating-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	// protected := r.Group("/")
	// protected.Use(middlewares.AuthMiddleware())
	// {
	// 	protected.GET("/profiles", controllers.GetProfiles)
	// 	protected.POST("/swipe", controllers.Swipe)
	// }
}
