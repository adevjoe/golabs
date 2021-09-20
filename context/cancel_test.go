package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	value := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Millisecond)
		value <- struct{}{}
		select {
		case <-ctx.Done():
			fmt.Println("handle cancel")
		}
	}()
	<-value
	fmt.Println("receive")
	cancel()
	fmt.Println("cancel")
}
