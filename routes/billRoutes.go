package routes

import (
	"github.com/Bii-sama/go-figure.git/controllers"
	"github.com/Bii-sama/go-figure.git/middlewares"
	"github.com/gin-gonic/gin"
)


func BillRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.Use(middlewares.Auth())
	incomingRoutes.GET("/bills",controllers.GetAllBills())
	incomingRoutes.GET("/bills/bill_id", controllers.GetABill())
	incomingRoutes.POST("/createbill", controllers.CreateBill())
	incomingRoutes.PATCH("/updatebill", controllers.UpdateBill())
	incomingRoutes.DELETE("/deletebill", controllers.DeleteBill())
}