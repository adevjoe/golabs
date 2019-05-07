package types

import (
	"fmt"
	"testing"
)

func TestRemoveSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	for key, value := range a {
		if value == 3 {
			a = append(a[:key], a[key+1:]...)
		}
	}
	fmt.Println(a)
}
