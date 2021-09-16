package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestWG(t *testing.T) {
	fn := func(wg *sync.WaitGroup, num int) {
		defer wg.Done()
		fmt.Println("Hi, i'm ", num)
	}

	count := 5
	wg := &sync.WaitGroup{}
	wg.Add(count)
	for i := 1; i <= count; i++ {
		go fn(wg, i)
	}
	wg.Wait()
}
