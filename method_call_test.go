package main

import (
	"fmt"
	"reflect"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func (stu *student) SetDefaults() {
	stu.Name = "tangxin"
	if stu.Age == 0 {
		stu.Age = 100
	}
}

func (stu *student) Greeting() {
	fmt.Printf("hello %s, %d years old\n", stu.Name, stu.Age)
}

func (stu *student) Aloha(name string) {
	fmt.Println("aloha,", name)
}

func Test_MethodCall(t *testing.T) {
	stu := student{
		Name: "wangwu",
	}

	// 注意
	// 方法对象的方法接收者， 可以是 **指针对象** 也可以是 **结构体对象**
	// 如果是指针对象的方法， **结构体对象** 是不能调用起方法的
	rv := reflect.ValueOf(stu)
	prv := reflect.ValueOf(&stu)

	stu.Greeting()
	methodCall(prv, "SetDefaults")
	methodCall(rv, "Greeting") // 这里找到到
	methodCall(prv, "Aloha", reflect.ValueOf("boss"))
}

// 对象方法调用
// rv 目标对象, method 方法名称, in 参数
func methodCall(rv reflect.Value, method string, in ...reflect.Value) {

	// 通过方法名称获取 反射的方法对象
	mv := rv.MethodByName(method)
	// check mv 是否存在
	if !mv.IsValid() {
		fmt.Printf("mv is zero value, method %s not found\n", method)
		return
	}

	// 调用
	// nil 这里代表参数
	mv.Call(in)
}
