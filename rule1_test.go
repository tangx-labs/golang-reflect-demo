package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 指向指针的指针对象
func Test_Rule1(t *testing.T) {
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
	fmt.Println(rv.Kind(), rv.Type()) // ptr **main.Person
}
