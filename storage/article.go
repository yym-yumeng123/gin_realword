package storage

import (
	"context"
	"gin_realword/models"
)

func CreateArticle(ctx context.Context, article *models.Article) error {
	return gormDB.WithContext(ctx).Create(article).Error
}

func UpdateArticle(ctx context.Context, oldSlug string, article *models.Article) error {
	return gormDB.WithContext(ctx).Model(&models.Article{}).Where("slug = ?", oldSlug).Updates(article).Error
}

func DeleteArticle(ctx context.Context, slug string) error {
	return gormDB.WithContext(ctx).Where("slug=?", slug).Delete(&models.Article{}).Error
}

func GetArticleBySlug(ctx context.Context, slug string) (*models.Article, error) {
	var article *models.Article
	tx := gormDB.WithContext(ctx).Model(&models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Where("slug = ?", slug).First(&article)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return article, nil
}

func ListArticles(ctx context.Context, limit, offset int) ([]*models.Article, error) {
	var articles []*models.Article
	err := gormDB.WithContext(ctx).Model(models.Article{}).
		Select("article.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article.author_username = user.username").
		Order("created_at desc").
		Offset(offset).Limit(limit).Find(&articles).Error

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func CountArticle(ctx context.Context) (int64, error) {
	var count int64
	err := gormDB.WithContext(ctx).Model(models.Article{}).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
