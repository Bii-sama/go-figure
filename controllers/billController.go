package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

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
)


var billCollection *mongo.Collection = database.NewCollection(database.Client, "bills")
var validateBill = validator.New()

type BillResponse struct{
	Count int64
	Bill  models.Bill
}

func GetAllBills() gin.HandlerFunc  {
return func(c *gin.Context) {

	var bill *models.Bill
if err := utils.CheckCreatedBy(c, *bill.Created_by); err != nil{
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
}

var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
defer cancel()

var bills []bson.M

cursor, err := billCollection.Find(ctx, bson.M{"created_by": bill.Created_by})

if err != nil{
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
	}
}

func GetABill()  {
	
}

func CreateBill() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 50 * time.Second)
		defer cancel()
		var bill models.Bill
		var user models.User

		if err := c.BindJSON(&bill); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		isBillValid := validateBill.Struct(bill)

		if isBillValid != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": isBillValid.Error()})
		}

		count, err := billCollection.CountDocuments(ctx, bson.M{"created_by": bill.Created_by})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
				 }
		
		
		 response := BillResponse{
					Count: count,
					Bill:  bill,
		  }
		  
		  c.JSON(http.StatusOK, response)

		  bill.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		  bill.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		  createdBy :=  strings.Split(*user.Email, "@")

		  if len(createdBy) > 0{
			bill.Created_by = &createdBy[0]
		  } else {
			log.Panicln("Invalid Email")
		  }

		  newBill, billErr := billCollection.InsertOne(ctx, bill)

         if billErr != nil {
			msg := fmt.Sprintf("New Bill was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		 }
		 defer cancel()
		 c.JSON(http.StatusOK, newBill)

	}
}

func UpdateBill()  {
	
}

func DeleteBillBill()  {
	
}
