package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestReplicated(t *testing.T) {
	doWork := func(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan<- int) {

		started := time.Now()
		defer wg.Done()

		// 模拟随机加载
		simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
		select {
		case <-done:
		case <-time.After(simulatedLoadTime):
		}

		select {
		case <-done:
		case result <- id:
		}

		took := time.Since(started)
		// 显示处理程序将花费多长时间
		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}
		fmt.Printf("%v took %v\n", id, took)
	}

	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ { //1
		go doWork(done, i, &wg, result)
	}

	firstReturned := <-result //2
	close(done)               //3
	wg.Wait()

	fmt.Printf("Received an answer from #%v\n", firstReturned)
}
