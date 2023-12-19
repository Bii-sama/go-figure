package utils

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"context"
	"os"
	"time"
	"github.com/Bii-sama/go-figure.git/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type SignUpInfo struct{
	Email string
	Firstname string
	Lastname string
	User_Type string
	Uid string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("JWT_SECRET")

func GenerateAllTokens(email string, firstname string, lastname string, user_type string, uid string)  {
	
}