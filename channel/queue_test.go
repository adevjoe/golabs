package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	var pics []string
	for i := 0; i < 1000000; i++ {
		pics = append(pics, fmt.Sprintf("pic_%d", i))
	}
	picChan := make(chan string)
	stopChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		produce(pics, picChan, stopChan)
		wg.Done()
	}()
	go func() {
		consume(picChan, stopChan)
		wg.Done()
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(stopChan)
		wg.Done()
	}()
	wg.Wait()
}

func produce(pics []string, ch chan string, stop chan struct{}) {
	defer close(ch)
	for _, pic := range pics {
		select {
		case _, ok := <-stop:
			if !ok {
				return
			}
		default:
			ch <- pic
		}
	}
}

func consume(ch chan string, stop chan struct{}) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case pic, ok := <-ch:
					if !ok {
						fmt.Printf("work %d done\n", i)
						return
					}
					fmt.Printf("work %d get pic: %s\n", i, pic)
				case _, ok := <-stop:
					if !ok {
						fmt.Printf("work %d stop\n", i)
						return
					}
				}
			}
		}(i)
	}
	wg.Wait()
}
