package context

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func TestContextDone(t *testing.T) {
	logs = log.New(os.Stdout, "", log.Ltime)

	// 新建一个ctx
	ctx, cancel := context.WithCancel(context.Background())

	// 传递ctx
	go doClearn(ctx)
	go doNothing(ctx)

	// 主程序阻塞20秒，留给协程来演示
	time.Sleep(10 * time.Second)
	logs.Println("cancel")

	// 调用cancel：context.WithCancel 返回的CancelFunc
	cancel()

	// 发出cancel 命令后，主程序阻塞10秒，再看协程的运行情况
	time.Sleep(5 * time.Second)
}
