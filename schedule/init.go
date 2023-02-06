// 后台定时任务初始化
package schedule

import (
	"github.com/robfig/cron/v3"
)

// 后台任务初始化
func InitBackJobs() {
	c := cron.New()

	c.AddFunc("@every 10s", ServerStatusMonitorTask) // 每一秒获取一次负载信息

	c.Start()
}
