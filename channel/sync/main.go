package sync

import (
	"fmt"
	"time"
)

func worker(done chan int) {
	time.Sleep(time.Second)
	// 通知任务已完成
	fmt.Println(2, "go start")
	done <- 1
	fmt.Println(5, "go done")
}

func main() {
	done := make(chan int) // 无缓冲，阻塞
	go worker(done)
	// 等待任务完成
	fmt.Println(1, "sleep 3 sec")
	time.Sleep(3 * time.Second)
	fmt.Println(3, <-done)
	fmt.Println(4, "sleep 2 sec")
	time.Sleep(2 * time.Second)
}
