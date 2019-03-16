package main

import "fmt"

type Week int

var names = [...]string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}

const (
	Monday Week = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (w Week) String() string {
	if w > 0 && w < 8 {
		return names[w-1]
	}
	return "非法的星期名"
}

func main() {
	fmt.Println(Wednesday)
}
