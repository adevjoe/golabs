package channel

import (
	"fmt"
	"testing"
	"time"
)

// 输入一组 channel，任意 channel ready 了，就执行完成。
func TestOrChannel(t *testing.T) {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		done := make(chan interface{})
		go func() {
			defer close(done)
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-or(append(channels[2:], done)...):
			}
		}()
		return done
	}

	start := time.Now()
	do := func(t time.Duration) <-chan interface{} {
		done := make(chan interface{})
		go func() {
			time.Sleep(t)
			done <- struct{}{}
		}()
		return done
	}
	<-or(
		do(time.Second),
		do(time.Hour),
		do(time.Millisecond),
		do(24*time.Hour),
	)
	fmt.Printf("done, cost %v\n", time.Since(start))
}
