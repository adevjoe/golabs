package channel

import (
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	pipelineFn(nil)
	normalPipeline(nil)

}

func normalPipeline(fn func()) {
	add := func(s []int, value int) []int {
		for i := range s {
			if fn != nil {
				fn()
			}
			s[i] += value
		}
		return s
	}
	multiply := func(s []int, value int) []int {
		for i := range s {
			if fn != nil {
				fn()
			}
			s[i] *= value
		}
		return s
	}

	a := []int{1, 2, 3, 4}
	multiply(add(multiply(a, 2), 1), 2)
}

func pipelineFn(fn func()) {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				if fn != nil {
					fn()
				}
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				if fn != nil {
					fn()
				}
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	s := []int{1, 2, 3, 4}
	intStream := generator(done, s...)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	b := make([]int, len(s))
	for v := range pipeline {
		b = append(b, v)
	}
}

func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pipelineFn(func() {
			time.Sleep(10 * time.Millisecond)
		})
	}
}

func BenchmarkNormalPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalPipeline(func() {
			time.Sleep(10 * time.Millisecond)
		})
	}
}
