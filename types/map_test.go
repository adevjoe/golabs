package types

import (
	"fmt"
	"testing"
)

func TestAppendMap(t *testing.T) {
	m := map[string]int{"foo": 1}
	AppendMap(m)   // set m["bar"] = 2
	fmt.Println(m) // map[bar:2, foo:1]
}

func TestChangeMap(t *testing.T) {
	m := map[string]int{"foo": 1}
	ChangeMap(m)   // set m["foo"] = 2
	fmt.Println(m) // map[foo:2]
}
