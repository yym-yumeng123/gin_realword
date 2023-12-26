package storage

import (
	"context"
	"gin_realword/models"
	"testing"
	"time"
)

func TestCreateArticle(t *testing.T) {
	ctx := context.TODO()

	CreateArticle(ctx, &models.Article{
		AuthorUsername: "zs",
		Title:          "你好",
		Slug:           "xxx",
		Body:           "212",
		Description:    "21212",
		TagList:        "323",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	})
}
