package models

import (
	// "strconv"
	"time"

	"gorm.io/gorm"
)

// 服务器负载数据库模型
type ServerStatus struct {
	SSID int

	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
	Free        uint64

	ModelName string
	MHZ       float64
	Percent   float64

	CreateTime time.Time
}

// 服务器负载信息返回体
type ServerStatusInfo struct {
	Memory MemoryInfo
	CPU    CPUInfo
}

// 内存占用信息
type MemoryInfo struct {
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
}

// CPU占用信息
type CPUInfo struct {
	ModelName string
	MHZ       float64
	Percent   []float64
}

func (ServerStatus) TableName() string {
	return "server_status"
}

// 创建hook函数
func (ss *ServerStatus) BeforeCreate(db *gorm.DB) error {
	ss.CreateTime = time.Now()
	return nil
}
