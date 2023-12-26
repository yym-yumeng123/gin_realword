package storage

import (
	"context"
	"gin_realword/models"
)

func CreateArticle(ctx context.Context, article *models.Article) error {
	return gormDB.WithContext(ctx).Create(article).Error
}
