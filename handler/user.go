package handler

import (
	"fmt"
	"gin_realword/logger"
	"gin_realword/middleware"
	"gin_realword/models"
	"gin_realword/params/request"
	"gin_realword/params/response"
	"gin_realword/security"
	"gin_realword/storage"
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
	r.GET("api/profiles/:username", userProfile)
	r.Use(middleware.AuthMiddleware()).PUT("/api/user", editUser)
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

	hashPassword, err := security.HashPassword(body.User.Password)
	if err != nil {
		log.WithError(err).Errorf("hashPassword failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := storage.CreateUser(ctx, &models.User{
		Username: body.User.Username,
		Password: hashPassword,
		Email:    body.User.Email,
		Image:    "https://api.realworld.io/images/smiley-cyrus.jpeg",
		Bio:      "",
	}); err != nil {
		log.WithError(err).Errorf("create user failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

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

	dbUser, err := storage.GetUserByEmail(ctx, body.User.Email)
	if err != nil {
		log.WithError(err).Errorf("get user failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 密码对比, 密文
	if !security.CheckPassword(body.User.Password, dbUser.Password) {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := security.GeneratorJWT(dbUser.Username, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			Username: dbUser.Username,
			Bio:      "",
			Image:    "https://api.realworld.io/images/smiley-cyrus.jpeg",
		},
	})
}

func userProfile(ctx *gin.Context) {
	log := logger.New(ctx)
	userName := ctx.Param("username")
	log = log.WithField("username", userName)
	log.Infof("user Profile called, userName: %v\n", userName)
	user, err := storage.GetUserByUserName(ctx, userName)
	if err != nil {
		log.WithError(err).Infoln("get user by username failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, response.UserProfileResponse{
		UserProfile: response.UserProfile{
			Username:  user.Username,
			Bio:       user.Bio,
			Image:     user.Image,
			Following: "false",
		},
	})
}

func editUser(ctx *gin.Context) {
	log := logger.New(ctx)
	log.Infof("user: %v", security.GetCurrentUserName(ctx))
	var body request.EditUserRequest

	// 绑定参数
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// username email 不能为空
	if body.EditUserBody.Username == "" || body.EditUserBody.Email == "" {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 密码加密
	if body.EditUserBody.Password != "" {
		var err error
		body.EditUserBody.Password, err = security.HashPassword(body.EditUserBody.Password)
		if err != nil {
			log.WithError(err).Errorf("hash password failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	dbUser := &models.User{
		Username: body.EditUserBody.Username,
		Password: body.EditUserBody.Password,
		Email:    body.EditUserBody.Email,
		Image:    body.EditUserBody.Image,
		Bio:      body.EditUserBody.Bio,
	}

	// 数据库更新数据
	if err := storage.UpdateUserByUserName(ctx, security.GetCurrentUserName(ctx), dbUser); err != nil {
		log.WithError(err).Errorf("updated failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 生成 jwt
	token, err := security.GeneratorJWT(dbUser.Username, dbUser.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    dbUser.Email,
			Token:    token,
			Username: dbUser.Username,
			Bio:      dbUser.Bio,
			Image:    dbUser.Image,
		},
	})

}
