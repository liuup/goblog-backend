package models

import (
	"time"

	"gorm.io/gorm"
)

// 具体发布的文章
type Article struct {
	ArticleId int

	Content string // 文章的内容

	CreateTime  time.Time // 发布文章的时间
	UpdateTime  time.Time
	DeletedTime time.Time
	IsDeleted   bool
}

// 获取文章表名
func (Article) TableName(*gorm.DB) string {
	return "articles"
}

// 创建hook函数
func (at *Article) BeforeCreate(db *gorm.DB) error {
	at.CreateTime = time.Now()
	at.IsDeleted = false
	return nil
}

// 更新hook函数
func (at *Article) BeforeSave(db *gorm.DB) error {
	at.UpdateTime = time.Now()
	return nil
}
