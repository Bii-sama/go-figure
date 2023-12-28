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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}
	}
}