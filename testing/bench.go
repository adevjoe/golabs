package testing

import (
	"reflect"
)

func PrintSize(typeName string, t interface{}) {
	reflect.ValueOf(t).Kind()
	//fmt.Println(fmt.Sprintf("type: %s, size: %v bytes", typeName, unsafe.Sizeof(reflect.ValueOf(t).Kind())))
}
