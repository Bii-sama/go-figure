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
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.NewCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()  {
	
}

func PasswordVerification()  {
	
}

func SignUp()  {
	
}

func Login()  {
	
}

func GetUsers()  {
	
}

func GetAUser () gin.HandlerFunc  {

	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := utils.CheckTypeEqualsUserId(ctx, userId); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}