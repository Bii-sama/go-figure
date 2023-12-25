package utils


import (
	"github.com/gin-gonic/gin"
	"errors"
)

func CheckUserType(c *gin.Context, role string) (err error){
       userType := c.GetString("user_type")
	   err = nil

	   if userType != role{
		err = errors.New("Unauthorized to access this information")
		return err
	   }

	   return err
}

func CheckTypeEqualsUserId(ctx *gin.Context, userId string) (err error) {

	userType := ctx.GetString("user_type")
	uId := ctx.GetString("uid")

	err = nil

	if userType == "USER" && uId != userId{
		err = errors.New("Unauthorized to access this information")
		return err
	}

	err = CheckUserType(ctx, userType)

	return err
	
}