package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 反射对象转换为指针对象
func Test_Rule2(t *testing.T) {
	p := &Person{
		Name: "zhangsan",
		Age:  18,
		Addr: struct {
			City string
		}{
			City: "chengdu",
		},
	}

	rv := reflect.ValueOf(&p)
	rule2(rv)
}

func rule2(rv reflect.Value) {
	// check
	if !rv.CanInterface() {
		fmt.Println("rv is not settable: ", rv.Type())
		return
	}

	// convert
	rv = DerefValue(rv)
	irv := rv.Interface()
	fmt.Println(irv)

	// type assert
	v, ok := irv.(Person)
	fmt.Println(v, ok) // {zhangsan 18 {chengdu}} true

}
