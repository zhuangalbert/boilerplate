package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuangalbert/boilerplate/src/api/v1/middlewares"
	"github.com/zhuangalbert/boilerplate/src/api/v1/objects"
	"github.com/zhuangalbert/boilerplate/src/api/v1/services"
)

type UserController struct {
	userService *services.UserService
}

func UserControllerHandler(router *gin.Engine) {
	handler := &UserController{
		userService: services.UserServiceHandler(),
	}

	defaultMiddleware := middlewares.DefaultMiddleware{}
	route := router.Group("v1/users")
	route.POST("login", handler.Login)
	route.POST("register", handler.Register)

	route.Use(defaultMiddleware.AuthenticationMiddleware())
	{
		route.GET(":id", handler.GetUser)
		route.GET(":id/update", handler.Update)
	}
}

func (handler *UserController) GetUser(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	userData, err := handler.userService.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	context.JSON(http.StatusOK, userData)
}

func (handler *UserController) Login(context *gin.Context) {
	jsonRequest := &objects.UserLoginObjectRequest{
		Username: context.Query("username"),
		Password: context.Query("password"),
	}

	loginData, err := handler.userService.Login(jsonRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	context.JSON(http.StatusOK, loginData)
}

func (handler *UserController) Update(context *gin.Context) {
	jsonRequest := &objects.UserUpdateObjectRequest{
		Email: context.Query("email"),
		Phone: context.Query("phone"),
	}

	userId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	error := handler.userService.Update(userId, jsonRequest)

	if error != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	context.JSON(http.StatusOK, nil)
}

func (handler *UserController) Delete(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	deletedUser := handler.userService.Delete(userId)

	if deletedUser != nil {
		context.JSON(http.StatusBadRequest, nil)
	}

	context.JSON(http.StatusOK, nil)
}

func (handler *UserController) Register(context *gin.Context) {
	jsonRequest := &objects.UserStoreObjectRequest{
		Username: context.Query("username"),
		Password: context.Query("password"),
		Phone:    context.Query("phone"),
		Email:    context.Query("email"),
	}

	userId, err := handler.userService.Store(jsonRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, userId)
	}

	context.JSON(http.StatusOK, userId)
}
