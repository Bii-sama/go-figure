package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"


	"strconv"
	"time"

	"github.com/Bii-sama/go-figure.git/database"
	"github.com/Bii-sama/go-figure.git/models"
	"github.com/Bii-sama/go-figure.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


var billCollection *mongo.Collection = database.NewCollection(database.Client, "bills")
var validateBill = validator.New()

func GetAllBills()  {
	
}

func GetABill()  {
	
}

func CreateBill() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 50 * time.Second)
		defer cancel()
		var bill models.Bill

		if err := c.BindJSON(&bill); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		isBillValid := validateBill.Struct(bill)

		if isBillValid != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": isBillValid.Error()})
		}
	}
}

func UpdateBill()  {
	
}

func DeleteBillBill()  {
	
}
