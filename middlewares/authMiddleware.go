package middlewares

import(
	"fmt"
	"net/http"
	"github.com/Bii-sama/go-figure.git/utils"
	"github.com/gin-gonic/gin"
)


func Auth()gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintln("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err:= utils.ValidateToken(clientToken)

		if err != "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("firstname", claims.Firstname)
		c.Set("lastname", claims.Lastname)
		c.Set("userType", claims.User_Type)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}