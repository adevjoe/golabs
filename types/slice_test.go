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

func TestSliceIter(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(s)
	SliceIter(s)
	fmt.Println(s)
}

func TestSliceDelete(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(s)
	SliceDelete(s)
	fmt.Println(s)
}
