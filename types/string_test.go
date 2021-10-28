package types

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestString(t *testing.T) {
	a := "world"
	fmt.Printf("%s, %p\n", a, &a)
	a = randString()
	fmt.Printf("%s, %p\n", a, &a)
}

func randString() string {
	return fmt.Sprintf("%d", rand.Int())
}
