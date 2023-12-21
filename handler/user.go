package handler

import (
	"fmt"
	"gin_realword/params/request"
	"gin_realword/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUserHandler(r *gin.Engine) {
	usersGroup := r.Group("/api/users")
	{
		usersGroup.POST("", userRegistration)
		usersGroup.POST("/login", userLogin)
	}
}

func userRegistration(ctx *gin.Context) {
	body := request.UserRegistrationRequest{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(utils.JsonMarshal(body))
}

func userLogin(ctx *gin.Context) {
	body := request.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(utils.JsonMarshal(body), "login")

}
