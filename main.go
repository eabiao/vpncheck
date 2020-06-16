package main

import (
	"github.com/eabiao/goutils/logger"
	"github.com/eabiao/goutils/systray"
	"net"
	"os"
	"time"
)

var (
	log = logger.GetLogger()
)

func main() {
	tray := systray.NewSysTray()

	tray.OnRClick(func() {
		os.Exit(0)
	})

	// 内嵌图标数据转换为临时文件路径
	iconOnFile, _ := tray.IconBytesToFilePath(iconOnData)
	iconOffFile, _ := tray.IconBytesToFilePath(iconOffData)

	go func() {
		trayOn := false
		connectSuccess := false

		for {
			connectSuccess = checkConnection()
			if connectSuccess && !trayOn {
				trayOn = true
				tray.Show(iconOnFile, "vpn on at "+time.Now().Format("2006-01-02 15:04:05"))
			} else if !connectSuccess && trayOn {
				trayOn = false
				tray.Show(iconOffFile, "vpn on at "+time.Now().Format("2006-01-02 15:04:05"))
			}
			time.Sleep(1 * time.Second)
		}
	}()

	tray.Show(iconOffFile, "")
	tray.Run()
}

// 检查连接
func checkConnection() bool {
	conn, err := net.DialTimeout("tcp", "192.168.138.123:22", 1*time.Second)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}
