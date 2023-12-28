package response

import (
	"gin_realword/models"
	"time"
)

type ListArticlesResponse struct {
	ArticlesCount int64       `json:"articlesCount"`
	Articles      []*Articles `json:"articles"`
}

type Articles struct {
	Author          *ArticleAuthor `json:"author"`
	Title           string         `json:"title"`
	Slug            string         `json:"slug"`
	Body            string         `json:"body"`
	Description     string         `json:"description"`
	TagList         []string       `json:"tag_list"`
	Favorited       bool           `json:"favorited"`
	FavoritersCount int            `json:"favoritersCount"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func (a *Articles) FromDB(dbArticle *models.Article) {
	a.Author = &ArticleAuthor{
		Bio:       dbArticle.AuthorUserBio,
		Following: false,
		Image:     dbArticle.AuthorUserImage,
		Username:  dbArticle.AuthorUsername,
	}
	a.Title = dbArticle.Title
	a.Slug = dbArticle.Slug
	a.Body = dbArticle.Body
	a.Description = dbArticle.Description
	a.TagList = dbArticle.TagList
	a.Favorited = false
	a.FavoritersCount = 0
	a.CreatedAt = dbArticle.CreatedAt
	a.UpdatedAt = dbArticle.UpdatedAt
}

type ArticleAuthor struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}
