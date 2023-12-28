package handler

import (
	"gin_realword/logger"
	"gin_realword/middleware"
	"gin_realword/models"
	"gin_realword/params/request"
	"gin_realword/params/response"
	"gin_realword/security"
	"gin_realword/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

func AddArticleHandler(r *gin.Engine) {
	articlesGroup := r.Group("/api/articles")
	{
		articlesGroup.GET("", listArticles)
		articlesGroup.GET("/:slug", getArticle)

		articlesGroup.Use(middleware.AuthMiddleware())
		articlesGroup.POST("", createArticles)
		articlesGroup.PUT("/:slug", editArticles)
		articlesGroup.DELETE("/:slug", deleteArticles)
	}

}

func getArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")
	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusOK, gin.H{
		"article": respArticle,
	})
}

func createArticles(ctx *gin.Context) {
	var req request.CreateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.CreateArticle(ctx, &models.Article{
		AuthorUsername: security.GetCurrentUserName(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusCreated, gin.H{
		"article": respArticle,
	})
}

func editArticles(ctx *gin.Context) {
	oldSlug := ctx.Param("slug")
	var req request.CreateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 是否是当前用户
	oldArticle, err := storage.GetArticleBySlug(ctx, oldSlug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if oldArticle.AuthorUsername != security.GetCurrentUserName(ctx) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	slug := strings.ReplaceAll(req.Article.Title, " ", "-") + "-" + uuid.NewString()
	if err := storage.UpdateArticle(ctx, oldSlug, &models.Article{
		AuthorUsername: security.GetCurrentUserName(ctx),
		Title:          req.Article.Title,
		Slug:           slug,
		Body:           req.Article.Body,
		Description:    req.Article.Description,
		TagList:        req.Article.TagList,
	}); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	article, err := storage.GetArticleBySlug(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	respArticle := &response.Article{}
	respArticle.FromDB(article)
	ctx.JSON(http.StatusCreated, gin.H{
		"article": respArticle,
	})
}

func deleteArticles(ctx *gin.Context) {
	slug := ctx.Param("slug")
	err := storage.DeleteArticle(ctx, slug)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusNoContent)
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
		resp.Articles = append(resp.Articles, &response.Article{
			Author: &response.Author{
				Bio:       article.AuthorUserBio,
				Following: false,
				Image:     article.AuthorUserImage,
				Username:  article.AuthorUsername,
			},
			Title:          article.Title,
			Slug:           article.Slug,
			Body:           article.Body,
			Description:    article.Description,
			TagList:        article.TagList,
			Favorited:      false,
			FavoritesCount: 0,
			CreatedAt:      article.CreatedAt,
			UpdatedAt:      article.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, resp)

}
