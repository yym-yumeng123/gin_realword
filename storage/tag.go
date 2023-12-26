package storage

import (
	"context"
)

func ListPopularTag(ctx context.Context) ([]string, error) {
	var res []string
	err := gormDB.WithContext(ctx).Table("popular_tags").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
