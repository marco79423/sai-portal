package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/marco79423/sai-portal/service/app"
)

func main() {
	// 初始化
	application, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	// 啟動
	if err := application.Start(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		application.Stop()
	}()

	// 等待關閉訊號
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
