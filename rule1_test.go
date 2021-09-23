package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 第一定律: 对象类型转指针类型
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
	rule1(p)  // ptr *main.Person
	rule1(&p) // ptr **main.Person

}

func rule1(v interface{}) {
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Kind(), rv.Type())
}
