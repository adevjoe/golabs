package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)
	s := "abcdefghijklmnopqrstuvwxyz"
	go func() {
		for {
			if i, ok := <-ch1; ok {
				if i > len(s)-1 {
					done <- true
					return
				}
				fmt.Print(s[i : i+1])
				ch2 <- i + 1
			}
		}
	}()
	go func() {
		for {
			if i, ok := <-ch2; ok {
				if i > len(s)-1 {
					done <- true
					return
				}
				fmt.Print(s[i : i+1])
				ch1 <- i + 1
			}
		}
	}()
	ch1 <- 0
	<-done
}
