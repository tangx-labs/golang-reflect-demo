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

	dert := deref(rt)
	logrus.Info(dert, dert.Kind())

}

func deref(rt reflect.Type) reflect.Type {
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return rt
}

func reflectValue(v interface{}) {
	rv := reflect.ValueOf(v)
	logrus.Info(rv.Kind(), rv.Type())

	irv := reflect.Indirect(rv)
	logrus.Info(irv.Kind(), irv.Type())

	if irv.CanAddr() {
		irvPtr := irv.Addr()
		logrus.Info(irvPtr.Kind(), irvPtr.Type())
	} else {
		logrus.Warn("irv can not addr", irv.Type())
	}

	// 4. deref
	iirv := derefValue(rv)
	logrus.Info(iirv.Kind(), iirv.Type())
}

func derefValue(rv reflect.Value) reflect.Value {
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	return rv
}
