package syntax

import "fmt"

type Person interface {
	Say()
}

type Doctor struct {
	Name string
}

func (s Doctor) Say() {
	fmt.Println("Hi! I'm a doctor.")
}

type Student struct {
	Name string
}

func (s Student) Say() {
	fmt.Println("Hi! I'm a student.")
}

func PrintPerson(i Person) {
	i.Say()
}
