package middleware

import (
	"blog1/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JwtAuthMiddleware is middle ware handler for authorizing the routes.
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := auth.ExtractTokenID(c)
		if err != nil {
			fmt.Printf("%v", err)
			c.String(http.StatusUnauthorized, "user is not Authorized")
			c.Abort()
			return

		}

		c.Next()
	}

}
