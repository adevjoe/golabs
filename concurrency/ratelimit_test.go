package concurrency

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"sync"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	limit := rate.NewLimiter(rate.Every(500*time.Millisecond), 1)
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			if err := limit.WaitN(ctx, 1); err != nil {
				fmt.Println(err.Error())
			} else {
				log.Print("done", i)
			}
		}(i)
	}
	wg.Wait()
}
