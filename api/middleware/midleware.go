package middleware

import (
	"AwsServerLessCleanCodeArchitecture/pkg/jwtparser"
	"errors"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const prefix = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithError(400, errors.New("Authorization header missing or empty "))
		}
		token := header[len(prefix):]
		claim, err := jwtparser.Validate(token, secret)
		if err != nil {
			c.AbortWithError(400, err)
		}
		c.Set("username", claim.Username)
		c.Next()
	}
}
