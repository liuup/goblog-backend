package models

// 服务器负载信息返回体
type ServerStatus struct {
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
