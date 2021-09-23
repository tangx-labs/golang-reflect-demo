package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Rule3(t *testing.T) {
	p := &Person{
		Name: "zhangsan",
		Age:  18,
		Addr: struct {
			City string
		}{
			City: "chengdu",
		},
	}

	rv := reflect.ValueOf(p)
	rv = DerefValue(rv)

	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)

		rule3(fv)
	}

	fmt.Println(*p)
}

func rule3(rv reflect.Value) {
	if !rv.CanSet() {
		fmt.Println("rv is not settable", rv.Type())
		return
	}

	switch rv.Kind() {
	case reflect.String:
		rv.SetString("tangxin")
		fmt.Println("set string")
	case reflect.Int:
		rv.SetInt(333)
		fmt.Println("set int")
	default:
		fmt.Println("not support kind: ", rv.Kind())
	}
}
