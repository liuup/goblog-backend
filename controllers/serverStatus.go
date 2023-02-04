package controllers

import (
	"log"
	"net/http"
	"time"

	"goblog/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// 允许跨域
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// websocket获取服务器负载情况
func HandleServerStatus(c *gin.Context) {

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// websocket读取线程，负责快速关闭连接
	go readFromWS(ws)

	// websocket发送线程
	go sendServerStatus(ws)

	select {}
}

// websocket服务器负载读取线程，负责快速关闭连接
func readFromWS(ws *websocket.Conn) {
	defer ws.Close()

	for {
		_, _, err := ws.ReadMessage() // 读信息失败就关闭ws连接
		if err != nil {
			break
		}
	}
}

// websocket服务器负载发送线程
func sendServerStatus(ws *websocket.Conn) {
	defer ws.Close()

	sstatus := models.ServerStatus{} // 返回体信息

	for {
		time.Sleep(1 * time.Second)

		// 获取服务器负载信息
		meminfo, err := mem.VirtualMemory()
		if err != nil {
			break
		}
		cpuinfo, err := cpu.Info()
		if err != nil {
			break
		}
		cpuPercent, err := cpu.Percent(0, true)
		if err != nil {
			break
		}

		sstatus.Memory = models.MemoryInfo{
			Total:       meminfo.Total,
			Available:   meminfo.Available,
			Used:        meminfo.Used,
			UsedPercent: meminfo.UsedPercent,
			Free:        meminfo.Free,
		}
		sstatus.CPU = models.CPUInfo{
			ModelName: cpuinfo[0].ModelName,
			MHZ:       cpuinfo[0].Mhz,
			Percent:   cpuPercent,
		}

		// 发送服务器负载信息
		err = ws.WriteJSON(sstatus)
		if err != nil {
			break
		}
	}
}
