package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	lock := sync.Mutex{}
	value := 0
	increment := func() {
		defer lock.Unlock()
		lock.Lock()
		value++
		fmt.Printf("incrementing, value is %d\n", value)
	}
	decrement := func() {
		defer lock.Unlock()
		lock.Lock()
		value--
		fmt.Printf("decrementing, value is %d\n", value)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
}
