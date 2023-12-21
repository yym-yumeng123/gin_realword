package handler

import (
	"fmt"
	"gin_realword/logger"
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
	log := logger.New(ctx)
	body := request.UserRegistrationRequest{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println()
	log.WithField("user", utils.JsonMarshal(body)).Infof("user registration called")
}

func userLogin(ctx *gin.Context) {
	log := logger.New(ctx)
	body := request.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.WithField("user", utils.JsonMarshal(body)).Infof("user login called")
}
