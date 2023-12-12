package routes

import (
	"github.com/Bii-sama/go-figure.git/controllers"
	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/signup", controllers.Login())
}
