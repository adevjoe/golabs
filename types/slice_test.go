package types

import (
	"fmt"
	"testing"
)

func TestCovert(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	fmt.Println("s1:", s1, "s2:", s2)
	Covert(s1, s2)
	fmt.Println("s1:", s1, "s2:", s2)
}
