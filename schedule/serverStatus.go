package schedule

import (
	"goblog/database"
	"goblog/models"
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// 系统负载监控任务
func ServerStatusMonitorTask() {
	sstatus, err := getServerStatus()
	if err != nil {
		log.Println(err)
		return
	}

	// TODO: 将负载信息写入数据库
	ss := models.ServerStatus{
		Total:       sstatus.Memory.Total / 1024,
		Available:   sstatus.Memory.Available / 1024,
		Used:        sstatus.Memory.Used / 1024,
		UsedPercent: sstatus.Memory.UsedPercent,
		Free:        sstatus.Memory.Free / 1024,

		ModelName: sstatus.CPU.ModelName,
		MHZ:       sstatus.CPU.MHZ,
		Percent:   sstatus.CPU.Percent[0],
	}

	result := database.DB.Create(&ss)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
}

func getServerStatus() (s models.ServerStatusInfo, err error) {
	// 获取服务器负载信息
	meminfo, err := mem.VirtualMemory()
	if err != nil {
		return s, err
	}
	cpuinfo, err := cpu.Info()
	if err != nil {
		return s, err
	}
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return s, err
	}

	s.Memory = models.MemoryInfo{
		Total:       meminfo.Total,
		Available:   meminfo.Available,
		Used:        meminfo.Used,
		UsedPercent: meminfo.UsedPercent,
		Free:        meminfo.Free,
	}
	s.CPU = models.CPUInfo{
		ModelName: cpuinfo[0].ModelName,
		MHZ:       cpuinfo[0].Mhz,
		Percent:   cpuPercent,
	}

	return s, nil
}
