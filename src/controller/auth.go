package controller

import (
	"fmt"
	"net/http"

	"auth/src/constanta/response"
	"auth/src/models"
	"auth/src/service"
	"auth/src/service/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authUC  service.Auth
	tokenUC service.Token
}

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func NewAuthController(authUC service.Auth, tokenUC service.Token) AuthController {
	return &authController{
		authUC:  authUC,
		tokenUC: tokenUC,
	}
}

func (c *authController) Register(ctx *gin.Context) {

	var (
		register auth.RegisterRequest
		err      error
	)

	err = ctx.ShouldBind(&register)
	if err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.authUC.Register(ctx, register)
	if err != nil {
		res := response.BuildErrorResponse("Register Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := response.BuildResponse("Register Success!", nil)
	ctx.JSON(http.StatusCreated, res)
}

func (c *authController) Login(ctx *gin.Context) {

	var (
		result *response.ResultResponse
		login  auth.LoginRequest
		user   *models.User
		err    error
	)

	err = ctx.ShouldBind(&login)
	if err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user, err = c.authUC.Login(ctx, login)
	if err != nil {
		res := response.BuildErrorResponse("Login Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, _ = c.tokenUC.GenerateToken(ctx, user)

	res := response.BuildResponse("Login Success!", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) Logout(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	token, err := c.tokenUC.ValidateToken(authHeader)
	if err != nil {
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	email := fmt.Sprintf("%v", claims["email"])
	err = c.authUC.Logout(ctx, email)
	if err != nil {
		res := response.BuildErrorResponse("Login Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponse("Logout Success!", nil)
	ctx.JSON(http.StatusOK, res)
}
