package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// OSシグナルを受け取るためのチャンネルを作成
	signalChannel := make(chan os.Signal, 1)

	// sigChannelに取り込むOSシグナルの種別を指定
	signal.Notify(signalChannel, syscall.SIGKILL, syscall.SIGTERM)

	for {
		// OSシグナルを受け取ったら処理をする
		sig := <-signalChannel
		fmt.Printf("シグナルを受け取った: %s\n", sig)

		if sig == syscall.SIGTERM {
			fmt.Printf("SIGTERMを受け取ったのでここでグレースフルなシャットダウンのための準備をする\n")
		}
	}
}
