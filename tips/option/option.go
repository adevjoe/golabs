/**
option: 初始化结构体时，传递默认参数或按需传递参数
*/
package option

type Foo struct {
	Name   string
	option Option

	// ...
}

type Option struct {
	num int
	str string
}

type ModOption func(option *Option)

func WithNum(num int) ModOption {
	return func(option *Option) {
		option.num = num
	}
}

func WithStr(str string) ModOption {
	return func(option *Option) {
		option.str = str
	}
}

func NewFoo(name string, modOption ...ModOption) *Foo {
	// 默认值
	option := Option{
		num: 100,
		str: "default str",
	}

	for _, o := range modOption {
		o(&option)
	}

	return &Foo{
		Name:   name,
		option: option,
	}
}
