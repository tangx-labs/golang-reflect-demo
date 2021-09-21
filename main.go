package main

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
)

func main() {
	p := Person{
		Name: "zhangsan",
		Age:  10,
		Addr: struct {
			City string
		}{
			City: "chengdu",
		},
	}

	reflectx(&p)
}

func reflectx(p interface{}) {
	// ====
	logrus.Info("1. 获取反射对象， 可能是指针对象")
	rv := reflect.ValueOf(p)
	fmt.Println(rv.Kind(), rv.Type()) // ptr *main.Person

	// ====
	logrus.Info("2. 获取反射对象的真实对象")
	irv := reflect.Indirect(rv)
	fmt.Println(irv.Kind(), irv.Type()) // struct main.Person

	// ====
	logrus.Info("3. 通过反射对象获取原始对象( interface )")
	if irv.CanInterface() {
		v := irv.Interface()
		fmt.Println(v) // {zhangsan 10 {chengdu}}

		logrus.Info("3.2. 通过断言获取类型")
		vv, ok := v.(Person)
		if ok {
			// v.(Person) =  {zhangsan 10 {chengdu}}
			fmt.Println("v.(Person) = ", vv)
		}
	}

	// ====
	logrus.Info("4. 获取反射对象的指针地址")
	if irv.CanAddr() {
		irvPtr := irv.Addr()
		fmt.Println(irvPtr.Kind(), irvPtr.Type())
	}

	// if !rv.CanAddr() {
	// 	defer func() {
	// 		err := recover()
	// 		logrus.Warnln("4.2. 被传入的原始对象 CanAddr == false")
	// 		logrus.Warn(rv, ": ", err)
	// 	}()

	// 	rvPtr := rv.Addr()
	// 	fmt.Println(rvPtr.Kind(), rvPtr.Type())
	// }

	// ===

	logrus.Info("5. value to type")
	rt := irv.Type()
	fmt.Println(rt.Kind())
}

type Person struct {
	Name string
	Age  int
	Addr struct {
		City string
	}
}
