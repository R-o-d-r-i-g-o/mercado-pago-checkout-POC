package controllers

import (
	"code-space-backend-api/api/response"
	"code-space-backend-api/app/User/interfaces/dto/input"
	"code-space-backend-api/app/User/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	createUser usecases.CreateUser
	loginUser  usecases.LoginUser
}

func NewUserController(createUser usecases.CreateUser, loginUser usecases.LoginUser) UserController {
	return &userController{
		createUser: createUser,
		loginUser:  loginUser,
	}
}

func (u *userController) Create(ctx *gin.Context) {
	var user input.CreateUserInputDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	if err := u.createUser.Execute(ctx, user); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}

func (u *userController) Login(ctx *gin.Context) {
	var user input.LoginUserInputDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(response.ComposeError(err))
		return
	}

	token, err := u.loginUser.Execute(ctx, user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, response.ComposeError(err))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, response.Data(token))
}
