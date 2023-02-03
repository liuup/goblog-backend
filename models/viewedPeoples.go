package models

import (
	"time"

	"gorm.io/gorm"
)

// 网站观看过的人数
// 以ip来进行访问量统计
type ViewedPeoples struct {
	Id          int
	IPAddress   string    // ip地址
	CreateTime  time.Time // 第一次打开网页的时间
	UpdateTime  time.Time // 最后一次打开网页的时间
	DeletedTime time.Time // 记录删除的时间
	IsDeleted   bool      // 假删除，true表示删除，false表示未删除
}

func (ViewedPeoples) TableName() string {
	return "viewed_peoples"
}

// 创建hook函数
func (vp *ViewedPeoples) BeforeCreate(db *gorm.DB) error {
	vp.CreateTime = time.Now()
	vp.IsDeleted = false
	return nil
}

// 更新hook函数
func (vp *ViewedPeoples) BeforeSave(db *gorm.DB) error {
	vp.UpdateTime = time.Now()
	return nil
}
