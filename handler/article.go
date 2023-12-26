package handler

import (
	"gin_realword/logger"
	"gin_realword/params/response"
	"gin_realword/storage"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func AddArticleHandler(r *gin.Engine) {
	articlesGroup := r.Group("/api/articles")
	articlesGroup.GET("", listArticles)
}

func listArticles(ctx *gin.Context) {
	log := logger.New(ctx)
	limit, offset := cast.ToInt(ctx.Query("limit")),
		cast.ToInt(ctx.Query("offset"))

	log.Infof("list articles, limit: %v, offset: %v\n", limit, offset)

	articles, err := storage.ListArticles(ctx, limit, offset)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	total, err := storage.CountArticle(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var resp response.ListArticlesResponse
	resp.ArticlesCount = total
	for _, article := range articles {
		resp.Articles = append(resp.Articles, &response.Articles{
			Author: &response.ArticleAuthor{
				Bio:       article.AuthorUserBio,
				Following: false,
				Image:     article.AuthorUserImage,
				Username:  article.AuthorUsername,
			},
			Title:           article.Title,
			Slug:            article.Slug,
			Body:            article.Body,
			Description:     article.Description,
			TagList:         article.TagList,
			Favorited:       false,
			FavoritersCount: 0,
			CreatedAt:       article.CreatedAt,
			UpdatedAt:       article.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, resp)

}
