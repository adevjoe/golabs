package types

import (
	"fmt"
	"testing"
	"unsafe"
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
func TestCopy(t *testing.T) {
	a := []int64{4, 2, 1}

	b := make([]int64, len(a))
	copy(b, a)

	c := append([]int64{}, a...)

	d := [3]int64{4, 2, 1}

	e := []byte("hello")
	copy(e[2:], e)
	fmt.Printf("a: %v, point: %p, size: %v\n", a, a, unsafe.Sizeof(a))
	fmt.Printf("b: %v, point: %p, size: %v\n", b, b, unsafe.Sizeof(b))
	fmt.Printf("c: %v, point: %p, size: %v\n", c, c, unsafe.Sizeof(c))
	fmt.Printf("d: %v, point: %p, size: %v\n", d, &d, unsafe.Sizeof(d))
	fmt.Printf("e: %s, point: %p, size: %v\n", e, e, unsafe.Sizeof(e))
}

func TestAppendSlice(t *testing.T) {
	s := []int{2}
	AppendSlice(s)
	fmt.Println(s) // Output: [2]
}

func TestChangeSlice(t *testing.T) {
	s := []int{1}
	ChangeSlice(s) // set s[0] = 10
	fmt.Println(s) // Output: [10]
}

func TestGrow(t *testing.T) {
	a := []int{1, 2, 3, 4}
	fmt.Printf("%p\n", a)
	a = append(a, []int{5, 6, 7, 8}...)
	fmt.Printf("%p\n", a)
	b := make([]int, 5, 10)
	b[1] = 1
	fmt.Println(b, cap(b))
	fmt.Printf("%p\n", b)
	b = append(b, 5)
	fmt.Println(b, cap(b))
	fmt.Printf("%p\n", b)
}
