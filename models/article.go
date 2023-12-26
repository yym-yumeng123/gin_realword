package models

import "time"

type Article struct {
	Id             int `db:"id"`
	AuthorUsername string
	Title          string
	Slug           string
	Body           string
	Description    string
	TagList        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (Article) TableName() string {
	return "article"
}
