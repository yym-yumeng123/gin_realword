package request

type ListArticleQuery struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Tag    string `json:"tag"`
}

type CreateArticleRequest struct {
	Article *CreateArticleBody `json:"article"`
}

type CreateArticleBody struct {
	Title       string   `json:"title"`
	Body        string   `json:"body"`
	Description string   `json:"description"`
	TagList     []string `json:"tagList"`
}
