package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}

func TestOnceCycle(t *testing.T) {
	// Deadlock
	//var onceA, onceB sync.Once
	//var initB func()
	//initA := func() { onceB.Do(initB) }
	//initB = func() { onceA.Do(initA) }
	//onceA.Do(initA)
}
