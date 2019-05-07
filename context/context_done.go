package context

import (
	"context"
	"log"
	"time"
)

var logs *log.Logger

func doClearn(ctx context.Context) {
	// for 循环来每1秒work一下，判断ctx是否被取消了，如果是就退出
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logs.Println("doClearn:收到Cancel，做好收尾工作后马上退出。")
			return
		default:
			logs.Println("doClearn:每隔1秒观察信号，继续观察...")
		}
	}
}

func doNothing(ctx context.Context) {
	for {
		time.Sleep(3 * time.Second)
		select {
		case <-ctx.Done():
			logs.Println("doNothing:收到Cancel，但不退出......")

			// 注释return可以观察到，ctx.Done()信号是可以一直接收到的，return不注释意味退出协程
			//return
		default:
			logs.Println("doNothing:每隔3秒观察信号，一直运行")
		}
	}
}
