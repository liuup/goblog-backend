package models

import (
	"time"

	"gorm.io/gorm"
)

// blog评论区数据表

type Comment struct {
	CommentId int
	ArticleID int // 与之关联的文章id

	Content  string // 评论的具体内容
	NickName string // 用于显示的用户名
	Email    string // 邮箱

	IPAddress  string // 发布评论时的ip地址
	IPLocation string // ip解析后的具体国家/省份

	CreateTime  time.Time // 发布评论时的时间
	UpdateTime  time.Time // 评论修改的时间
	DeletedTime time.Time // 评论删除的时间
	IsDeleted   bool

	Avatar  string // 头像链接
	Website string // 个人网站
}

// 获取文章评论数据表名
func (Comment) TableName(*gorm.DB) string {
	return "comments"
}

// 创建hook函数
func (cm *Comment) BeforeCreate(db *gorm.DB) error {
	cm.CreateTime = time.Now()
	cm.IsDeleted = false
	return nil
}

// 更新hook函数
func (cm *Comment) BeforeSave(db *gorm.DB) error {
	cm.UpdateTime = time.Now()
	return nil
}
