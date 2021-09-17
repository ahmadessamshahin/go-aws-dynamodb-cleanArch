package handler

import (
	"AwsServerLessCleanCodeArchitecture/api/helper"
	"AwsServerLessCleanCodeArchitecture/api/middleware"
	"AwsServerLessCleanCodeArchitecture/api/model"
	"AwsServerLessCleanCodeArchitecture/entity"
	"AwsServerLessCleanCodeArchitecture/usecase/users"
	"fmt"
	"github.com/gin-gonic/gin"
)

// This is the Adaptor layer that convert the data from one format to another one
//([From & to] The format most convenient to useCase and entity && convenient format to external agency)

//handler handle http request and response and validation

type GinHandler struct {
	UseCase users.UseCase
}

func NewGinHandler(useCase users.UseCase, jwtSecret string) (r *gin.Engine) {
	h := &GinHandler{
		useCase,
	}
	r = gin.Default()
	r.POST("/users/signup", h.SignUp)
	r.POST("/users/signin", h.SignIn)
	r.GET("/hello", middleware.TokenAuthMiddleware(jwtSecret), h.SayHello)
	return r
}
func (h *GinHandler) SignUp(c *gin.Context) {
	//	serializable
	var l entity.User
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	err = l.Validate()

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	err = h.UseCase.Signup(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.SignupOutput{Message: fmt.Sprintf("user %s created successfully", l.Username)}
	c.JSON(201, res)
}
func (h *GinHandler) SignIn(c *gin.Context) {
	var l model.SigninInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	var user entity.User
	user.Username, user.Password = l.Username, l.Password

	token, err := h.UseCase.SignIn(user)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}
	res := model.SigninOutput{Token: token}

	c.JSON(200, res)
}
func (h *GinHandler) SayHello(c *gin.Context) {
	var l entity.User
	l.Username = c.GetString("username")

	message, err := h.UseCase.SayHello(l)
	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.HelloOutput{Message: message}

	c.JSON(200, res)
}
