package models

import "time"

// 网站观看过的人数
// 以ip来进行访问量统计
type ViwedPeoples struct {
	Id          int
	IPAddress   string    // ip地址
	CreateTime  time.Time // 第一次打开网页的时间
	EditedTime  time.Time // 最后一次打开网页的时间
	DeletedTime time.Time // 记录删除的时间
	IsDeleted   bool      // 假删除，true表示删除，false表示未删除
}
