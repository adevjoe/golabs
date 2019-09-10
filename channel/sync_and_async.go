package channel

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("start")
	c := make(chan int)
	go func(c chan int) {
		fmt.Println("in goroutine")
		time.Sleep(4 * time.Second)
		c <- 2
		fmt.Println("out goroutine")
	}(c)
	fmt.Println("do something")
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	a := <-c
	fmt.Println("receive value", a)
}
