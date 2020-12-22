package types

import "fmt"

func Covert(s1, s2 []int) {
	for _, value := range s1 {
		for i := range s2 {
			s2[i] = value
		}
	}
}

func SliceIter(s []int) {
	for _, v := range s {
		fmt.Println(v)
		s = append(s, v)
	}
	fmt.Println(s)
}

func SliceDelete(s []int) {
	for i, v := range s {
		fmt.Println(v)
		s[i] = 0
		fmt.Println(s)
	}
	s[1] = -1
}

func AppendSlice(s []int) {
	s = append(s, 1)
}

func ChangeSlice(s []int) {
	if len(s) > 0 {
		s[0] = 10
	}
}
