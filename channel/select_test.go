package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectSend(t *testing.T) {
	a := make(chan struct{})
	go func() {
		<-a
	}()
	select {
	case <-a:
		fmt.Println("receive")
	case a <- struct{}{}:
		fmt.Println("send")
	}
}

func TestSelectReceive(t *testing.T) {
	a := make(chan struct{})
	go func() {
		a <- struct{}{}
	}()
	select {
	case <-a:
		fmt.Println("receive")
	case a <- struct{}{}:
		fmt.Println("send")
	}
}

func TestSelectBlock(t *testing.T) {
	a := make(chan struct{})
	go func() {
		a <- struct{}{}
	}()
	select {
	case <-a:
		fmt.Println("receive")
	case a <- struct{}{}:
		fmt.Println("send")
	}
	fmt.Println("done")
}

func TestSelectDefault(t *testing.T) {
	a := make(chan struct{})
	start := time.Now()
	select {
	case <-a:
		fmt.Println("receive")
	case a <- struct{}{}:
		fmt.Println("send")
	default:
		fmt.Printf("default after %v\n", time.Since(start))
	}
}

func TestChannelClose(t *testing.T) {
	a := make(chan struct{})
	close(a)
	select {
	case <-a:
		fmt.Println("receive")
	}
	fmt.Println("done")
}

// 测试多个 case select 执行情况
// 由于是随机的，每个 case 得到执行的机会是一样的
func TestMultiCase(t *testing.T) {
	a := make(chan struct{})
	b := make(chan struct{})
	close(a)
	close(b)
	countA := 0
	countB := 0
	for i := 0; i < 1000; i++ {
		select {
		case <-a:
			countA++
		case <-b:
			countB++
		}
	}
	fmt.Printf("a: %d, b: %d\n", countA, countB)
}

func TestSelectTimeout(t *testing.T) {
	a := make(chan struct{})
	select {
	case <-a:
		fmt.Println("receive")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}
