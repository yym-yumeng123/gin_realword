package handler

import (
	"fmt"
	"gin_realword/logger"
	"gin_realword/params/request"
	"gin_realword/params/response"
	"gin_realword/security"
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

	// TODO: insert data to db

	token, err := security.GeneratorJWT(body.User.Username, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			Username: body.User.Username,
			Bio:      "",
			Image:    "https://api.realworld.io/images/smiley-cyrus.jpeg",
		},
	})
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

	// TODO: get username from db

	userName := "yym"
	token, err := security.GeneratorJWT(userName, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			Username: userName,
			Bio:      "",
			Image:    "https://api.realworld.io/images/smiley-cyrus.jpeg",
		},
	})
}
