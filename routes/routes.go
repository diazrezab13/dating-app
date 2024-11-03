package routes

import (
	"dating-app/controllers"
	"dating-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	protected := r.Group("/profile")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/", controllers.GetProfiles)
		protected.POST("/swipe", controllers.Swipe)
		protected.POST("/upgrade", controllers.UpgradePremium)
	}
}
