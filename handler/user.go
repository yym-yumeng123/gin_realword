package handler

import "github.com/gin-gonic/gin"

func AddUserHandler(r *gin.Engine) {
	usersGroup := r.Group("/api/users")
	{
		usersGroup.POST("", userRegistration)
		usersGroup.POST("/login", userLogin)
	}
}

func userRegistration(ctx *gin.Context) {

}

func userLogin(ctx *gin.Context) {

}
