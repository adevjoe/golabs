package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	time1 := time.Now()
	task := make([]int, 0)
	for i := 0; i < 5; i++ {
		task = append(task, i)
	}
	wg := &sync.WaitGroup{}
	done := make(chan string)

	for key, value := range task {
		wg.Add(1)
		go func(k, v int, wg *sync.WaitGroup, done chan string) {
			defer wg.Done()
			if k > 3 {
				return
			}
			fmt.Printf("log: task %d start.\n", k)
			time.Sleep(time.Duration(v) * time.Second)
			fmt.Printf("log: %d done, value is %d.\n", k, v)
			done <- fmt.Sprintf("%d done, value is %d.", k, v)
		}(key, value, wg, done)
	}
	go func(wg *sync.WaitGroup, done chan string) {
		wg.Wait()
		close(done)
	}(wg, done)
	for value := range done {
		fmt.Println(value)
	}
	fmt.Println(time.Since(time1))
}

