package response

import "time"

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

type ArticleAuthor struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}
