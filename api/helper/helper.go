package helper

import (
	"AwsServerLessCleanCodeArchitecture/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Unmarshal(c *gin.Context, t interface{}) error {
	x, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(x, &t)

	if err != nil {
		return entity.ErrInvalidInput
	}
	return nil
}

func ErrHandler(err error, c *gin.Context) {
	switch err {
	case entity.ErrUserAlreadyExists:
		c.AbortWithStatusJSON(409, gin.H{"error": err.Error()})
	case entity.ErrInvalidPassword:
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
	case entity.ErrUserDoesNotExist:
		c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
	case entity.ErrInvalidInput:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	default:
		c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
	}
}
