package server

import (
	"gin_realword/handler"
	"github.com/gin-gonic/gin"
)

func RunHTTPServer() {
	r := gin.Default()
	handler.AddUserHandler(r)
	handler.AddArticleHandler(r)
	handler.AddTagsHandler(r)
	handler.AddCommentsHandler(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
