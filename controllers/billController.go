package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Bii-sama/go-figure.git/database"
	"github.com/Bii-sama/go-figure.git/models"
	"github.com/Bii-sama/go-figure.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/bson/primitive"
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

	var bill models.Bill
if err := utils.CheckCreatedBy(c, *bill.Created_by); err != nil{
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
}

var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
defer cancel()

var bills[] *models.Bill

cursor, err := billCollection.Find(ctx, bson.M{"created_by": bill.Created_by})

if err != nil{
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

defer cursor.Close(ctx)

for cursor.Next(ctx){
	var aBill *models.Bill

	if err := cursor.Decode(&aBill); err != nil{
		log.Println(err)
		continue
	}

	bills = append(bills, aBill)
}
if err := cursor.Err(); err != nil {
   log.Println(err)
}
  c.JSON(http.StatusOK, bills)
	}
}

func GetABill() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 50 * time.Second)
		defer cancel()

	var bill *models.Bill
    billId := c.Param("bill_id")

	if err := utils.CheckCreatedBy(c, *bill.Created_by); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	if billId != bill.Bill_ID {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bill does not exist"})
		return
	}

	err:= billCollection.FindOne(ctx, bson.M{"bill_id": billId}).Decode(&bill)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bill)
	}

}

func CreateBill() gin.HandlerFunc  {
	return func(c *gin.Context) {
		
		var bill models.Bill
		var user models.User

		if err := c.BindJSON(&bill); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		isBillValid := validateBill.Struct(bill)

		if isBillValid != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": isBillValid.Error()})
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 50 * time.Second)
		defer cancel()

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

func UpdateBill() gin.HandlerFunc {
	return func(c *gin.Context) {
		billId := c.Param("bill_id")

		// Fetch the existing bill from the database
		var existingBill models.Bill
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()

		opts := options.FindOne().SetProjection(bson.M{"_id": 0}) // Exclude _id field from the retrieved document

		if err := billCollection.FindOne(ctx, bson.M{"bill_id": billId}, opts).Decode(&existingBill); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if existingBill.Bill_ID != billId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bill does not exist"})
			return
		}

	
	updatedFields := models.Bill{} 

		updatedFieldsMap := bson.M{
			"$set": bson.M{
				"customer_name": updatedFields.Customername,
                "email": updatedFields.Email,
				"items": updatedFields.Items,
			},
		}

		updateOpts := options.FindOneAndUpdate().SetReturnDocument(options.After)

		// Perform the update operation
		if err := billCollection.FindOneAndUpdate(ctx, bson.M{"bill_id": billId}, updatedFieldsMap, updateOpts).Decode(&existingBill); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, existingBill)
	}
}



func DeleteBill() gin.HandlerFunc  {
	return func(ctx *gin.Context) {}
}
