package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextTimeOut(t *testing.T) {
	// 自动取消(定时取消)
	{
		timeout := 2 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)

		fmt.Println("A 执行完成，返回：", A(ctx))
		select {
		case <-ctx.Done():
			fmt.Println("context Done")
			break
		}
	}
	time.Sleep(5 * time.Second)
}
