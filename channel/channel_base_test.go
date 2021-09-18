package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

}

func TestChannelPoint(t *testing.T) {
	// new channel and assign a new channel
	a := new(chan int)
	fmt.Printf("chan a point: %p\n", a)
	b := make(chan int)
	a = &b
	fmt.Printf("chan a point: %p\n", a)
	fmt.Printf("chan b point: %p\n", b)
	go func() {
		*a <- 1
	}()
	fmt.Printf("chan a point: %p\n", a)
	fmt.Println(<-*a)

	// new channel
	c := new(chan int)
	fmt.Printf("chan c point: %p\n", c)
	*c = make(chan int)
	fmt.Printf("chan c point: %p\n", c)
	go func() {
		*c <- 1
	}()
	fmt.Printf("chan c point: %p\n", c)
	fmt.Println(<-*c)

	// make channel
	d := make(chan int)
	fmt.Printf("chan d point: %p\n", d)
	go func() {
		d <- 1
	}()
	fmt.Printf("chan d point: %p\n", d)
	fmt.Println(<-d)
}

// 没初始化 channel，会死锁
func test1() {
	var a chan int
	go func() {
		a <- 1
	}()
	fmt.Println(<-a)
}

// 测试单向 channel
func TestUnidirectionalChannel(t *testing.T) {
	value := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		send(value)
	}()
	go func() {
		defer wg.Done()
		receive(value)
	}()
	wg.Wait()
}

func send(value chan<- int) {
	// 不能从只写 channel 读取，下面这样写会编译报错
	// <-value
	fmt.Println("send value")
	value <- 1
}

func receive(value <-chan int) {
	// 不能往只读 channel 写入，下面这样写会编译报错
	// value <- 1
	fmt.Println("receive", <-value)
}

func TestChannelBlocking(t *testing.T) {
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		sendMany(c, 10)
	}()
	go func() {
		defer wg.Done()
		receiveSlow(c)
	}()
	wg.Wait()
}

func sendMany(c chan int, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("i'm sending %d\n", i)
		c <- i
	}
	fmt.Println("send done")
	close(c)
}

func receiveSlow(c chan int) {
	for {
		if v, ok := <-c; ok {
			fmt.Printf("received %d\n", v)
			time.Sleep(10 * time.Millisecond)
		} else {
			fmt.Println("receive done")
			break
		}
	}
}

func TestForRange(t *testing.T) {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()
	for v := range c {
		fmt.Println(v)
	}
}
