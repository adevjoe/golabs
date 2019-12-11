package option

import "testing"

func TestNewFoo(t *testing.T) {
	foo := NewFoo("joe")
	t.Log(foo)
	foo1 := NewFoo("joe", func(option *Option) {
		option.num = 200
	})
	t.Log(foo1)
	foo2 := NewFoo("joe", func(option *Option) {
		option.num = 300
	}, WithStr("abc"))
	t.Log(foo2)
	foo3 := NewFoo("joe", func(option *Option) {
		option.num = 300
	}, WithStr("abc"), WithNum(400))
	t.Log(foo3)
}
