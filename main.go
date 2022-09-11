package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type FuncType func(argc ...any) any

func Add(argc ...any) any {
	ch := 0
	for _, key := range argc {
		ch += int(key.(int))
	}
	return ch
}

func addblock() {

}
func demo(f FuncType) {

	//c := f(1, 2)
	v := reflect.ValueOf(f)
	//k := v.Type()
	size := unsafe.Sizeof(f)
	fmt.Println(v.Kind() == reflect.Ptr, size)
}

func main() {
	demo(Add)
	fmt.Println("Hello World")
}
