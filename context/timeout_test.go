package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("handle", ctx.Err())
		case <-time.After(100 * time.Millisecond):
			fmt.Println("process done")
		}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}
