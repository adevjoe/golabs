package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestCondBroadcast(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	value := int32(0)
	wg := sync.WaitGroup{}
	wg.Add(2)
	dec(c, func() {
		defer wg.Done()
		atomic.AddInt32(&value, -1)
		fmt.Println("A got it. -1")
	})
	dec(c, func() {
		defer wg.Done()
		atomic.AddInt32(&value, -2)
		fmt.Println("B got it. -2")
	})
	c.Broadcast()
	wg.Wait()
}

func dec(c *sync.Cond, fn func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fn()
	}()
	wg.Wait()
}

func TestCondSignal(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	value := int32(0)
	done := make(chan bool)
	dec(c, func() {
		atomic.AddInt32(&value, -1)
		fmt.Println("A got it. -1")
		done <- true
	})
	dec(c, func() {
		atomic.AddInt32(&value, -2)
		fmt.Println("B got it. -2")
		done <- true
	})
	c.Signal()
	<-done
}

func TestButton(t *testing.T) {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
