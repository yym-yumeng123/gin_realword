package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Article struct {
	Id             int64 `db:"id"`
	AuthorUsername string
	Title          string
	Slug           string
	Body           string
	Description    string
	TagList        TagList `gorm:"type:string"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	AuthorUserEmail string `gorm:"->"`
	AuthorUserImage string `gorm:"->"`
	AuthorUserBio   string `gorm:"->"`
}

func (Article) TableName() string {
	return "article"
}

type TagList []string

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *TagList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, j)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j TagList) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}
