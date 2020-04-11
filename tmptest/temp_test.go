package tmptest

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	a := struct {
		Name string
	}{Name: "joe"}
	v, err := json.Marshal(a)
	fmt.Println(string(v), err)
	b := &struct {
		Name string
	}{Name: "joe"}
	bb, err := json.Marshal(b)
	fmt.Println(string(bb), err)
}
