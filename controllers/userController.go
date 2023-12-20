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

func HashPassword(password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil{
      log.Panicln(err)
	}
	return string(hashedPassword)
}

func PasswordVerification(userPassword string, enteredPassword string)(bool, string)  {
	err := bcrypt.CompareHashAndPassword([]byte(enteredPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil{
		msg = fmt.Sprintf("Email/Password is invalid")
		check = false
	}
	return check, msg
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
		
		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()


		token, refreshToken, _ := utils.GenerateAllTokens(*user.Email, *user.Firstname, *user.Lastname, *user.User_Type, *&user.User_ID)
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

func Login() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

		var user *models.User

		var checkUser *models.User

		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&checkUser)
		defer cancel()

		if err != nil{
          c.JSON(http.StatusInternalServerError, gin.H{"error": "Email/Password invalid"})
		  return
		}

		passwordCheck, msg := PasswordVerification(*user.Password, *checkUser.Password)
		defer cancel()

		if passwordCheck != true{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if checkUser.Email == nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}

		token, refreshToken, _:= utils.GenerateAllTokens(*checkUser.Email, *checkUser.Firstname, *checkUser.Lastname, *checkUser.User_Type, checkUser.User_ID)

	}
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