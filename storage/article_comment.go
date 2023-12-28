package storage

import (
	"context"
	"gin_realword/models"
)

func CreateArticleComment(ctx context.Context, articleComment *models.ArticleComment) error {
	return gormDB.WithContext(ctx).Create(articleComment).Error
}

func DeleteArticleCommentById(ctx context.Context, articleCommentId int64) error {
	return gormDB.Where("id = ?", articleCommentId).Delete(models.ArticleComment{}).Error
}

func GetArticleCommentsByArticleId(ctx context.Context, articleSlug string) ([]*models.ArticleComment, error) {
	comments := make([]*models.ArticleComment, 0)
	if err := gormDB.WithContext(ctx).
		Select("article_comment.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article_comment.author_username = user.username").
		Joins("INNER JOIN article ON article.id = article_comment.article_id").
		Where("article.slug = ?", articleSlug).
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func GetArticleCommentById(ctx context.Context, commentId int64) (*models.ArticleComment, error) {
	comment := models.ArticleComment{}
	if err := gormDB.WithContext(ctx).
		Select("article_comment.*, user.email as author_user_email, user.bio as author_user_bio, user.image as author_user_image").
		Joins("LEFT JOIN user ON article_comment.author_username = user.username").
		Where("article_comment.id = ?", commentId).
		Find(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
