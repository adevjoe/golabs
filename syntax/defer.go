package syntax

import "fmt"

func f1() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recover from f1:", x)
			panic("triggered from f1")
		}
	}()
	f2()
}
func f2() {
	panic("triggered from f2")
}
