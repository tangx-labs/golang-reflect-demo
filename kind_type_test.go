package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 打印 kind 和 type 的值
func kind_type_value(v interface{}) {
	rv := reflect.ValueOf(v)
	fmt.Println(rv.Kind(), rv.Type())
}

// kind 和 type 相同字面值
func Test_Kind_Type_Same(t *testing.T) {
	name := "tangxin"
	age := 18

	kind_type_value(name) // string string
	kind_type_value(age)  // int int

	kind_type_value(&name) // ptr *string
	kind_type_value(&age)  // ptr *int
}

// 根据内置类型 string 的自定义类型
type MyString string

// kind 和 type 不同
func Test_KindType_Different(t *testing.T) {
	p := Person{
		Name: "tagnxin",
		Age:  18,
	}
	kind_type_value(p)  // struct main.Person
	kind_type_value(&p) // ptr *main.Person

	s1 := MyString("tangxin")
	kind_type_value(s1)  // string main.MyString
	kind_type_value(&s1) // ptr *main.MyString
}
