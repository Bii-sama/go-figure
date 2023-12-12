package main

import(
	"github.com/Bii-sama/go-figure.git/routes"
	"os"
	"github.com/gin-gonic/gin"

)


func main()  {
	port := os.Getenv("PORT")

	if port == ""{
		port=  "8000"
	}

	router :=  gin.New()
	router.Use(gin.Logger())

	routes.Authroutes(router)
	routes.UserRoutes(router)

	router.GET("api-1", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"success": "Access granted for Api-1"})
	})

	router.GET("api-2", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"success": "Access granted for Api-2"})
	})

	router.Run(":" + port)
}