package handler

import (
	"gin_realword/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddTagsHandler(r *gin.Engine) {
	tagsGroup := r.Group("/api/tags")
	tagsGroup.GET("", listPopularTags)
}

func listPopularTags(ctx *gin.Context) {
	tag, err := storage.ListPopularTag(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"tags": tag,
	})
}
