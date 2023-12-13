package routes

import (
	"github.com/Bii-sama/go-figure.git/controllers"
	"github.com/Bii-sama/go-figure.git/middlewares"
	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.Use(middlewares.Auth()) 
	incomingRoutes.GET("/users")
	incomingRoutes.GET("users/user_id")
	
	
}
