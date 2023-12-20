package utils

import (
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

var userCollection *mongo.Collection = database.NewCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("JWT_SECRET")

func GenerateAllTokens(email string, firstname string, lastname string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignUpInfo{
		Email: email,
		Firstname: firstname,
		Lastname: lastname,
		User_Type: userType,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignUpInfo{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil{
		log.Panicln(err)
		return
	}

	return token, refreshToken, err
}