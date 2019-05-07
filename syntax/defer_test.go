package syntax

import (
	"fmt"
	"testing"
)

func TestDeferOrder(t *testing.T) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recover from main:", x)
		}
	}()
	f1()
}
