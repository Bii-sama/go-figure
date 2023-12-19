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

var userCollection *mongo.Collection = database.NewCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()  {
	
}

func PasswordVerification()  {
	
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(user)

		if validateErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()

		if err != nil{
			log.Panicln(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		}

		if count > 0{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email already exists"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()


		token, refreshToken, _ := utils.GenerateAllTokens(*user.Email, *user.Firstname, *user.Lastname, *&user.User_Type, *&user.User_ID)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertNo, insertErr := userCollection.InsertOne(ctx, user)

		if insertErr != nil{
			msg := fmt.Sprintf("User Item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertNo)
	}
}

func Login()  {
	
}

func GetUsers()  {
	
}

func GetAUser () gin.HandlerFunc  {

	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := utils.CheckTypeEqualsUserId(c, userId); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, user)
	}
}