package main

import (
	"github.com/eabiao/goutils/logger"
	"net"
	"os"
	"time"
)

var (
	log = logger.GetLogger()
)

func main() {
	tray := NewSysTray()

	// 点击图标退出程序
	tray.OnClick(func() {
		os.Exit(0)
	})

	go func() {
		trayOn := false
		connectSuccess := false

		for {
			connectSuccess = checkConnection()
			if connectSuccess && !trayOn {
				t := time.Now().Format("2006-01-02 15:04:05")
				trayOn = true
				tray.ShowIcon(iconOnData, "vpn on at "+t)
				//tray.Show("icon/on.ico", "vpn on at "+t)

			} else if !connectSuccess && trayOn {
				t := time.Now().Format("2006-01-02 15:04:05")
				trayOn = false
				//tray.Show("icon/on.ico", "vpn off at "+t)
				tray.ShowIcon(iconOffData, "vpn on at "+t)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	err := tray.Run()
	if err != nil {
		log.Error(err.Error())
	}
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
