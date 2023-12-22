package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Bii-sama/go-figure.git/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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



func UpdateTokens(signedToken string, signedRefreshToken string, uid string)  {
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

	var updateObject primitive.D

	updateObject = append(updateObject, primitive.E{Key:"token", Value: signedToken})
	updateObject = append(updateObject, primitive.E{Key:"refresh_token", Value: signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObject = append(updateObject, primitive.E{Key: "updated_at", Value: Updated_at})

	upsert := true
	filter := bson.M{"user_id": uid}

	opt:=options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObject}}, &opt,)


	defer cancel()

	if err != nil{

		log.Panic(err)
		return

	}

	return
}