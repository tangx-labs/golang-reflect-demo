package main

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

var p = &Person{
	Name: "zhangsan",
	Age:  18,
	Addr: struct {
		City string
	}{
		City: "chengdu",
	},
}

func Test_ValueAndType(t *testing.T) {
	// reflectType(&p)
	reflectValue(&p)
}

func reflectType(v interface{}) {
	rt := reflect.TypeOf(v)
	logrus.Info(rt, rt.Kind())

	irt := rt.Elem()
	logrus.Info(irt, irt.Kind())

	iirt := irt.Elem()
	logrus.Info(iirt, iirt.Kind())

	dert := DerefType(rt)
	logrus.Info(dert, dert.Kind())

}

func reflectValue(v interface{}) {
	rv := reflect.ValueOf(v)
	logrus.Info(rv.Kind(), rv.Type())

	// Indirect 只包含了一层 direct
	// 如果 **main.Person 这种数据结构， 得到的还是 reflect.Ptr 指针类型。
	irv := reflect.Indirect(rv)
	logrus.Info(irv.Kind(), irv.Type())

	if irv.CanAddr() {
		irvPtr := irv.Addr()
		logrus.Info(irvPtr.Kind(), irvPtr.Type())
	} else {
		logrus.Warn("irv can not addr", irv.Type())
	}

	// 4. deref
	iirv := DerefValue(rv)
	logrus.Info(iirv.Kind(), iirv.Type())
}
