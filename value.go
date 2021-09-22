package main

import "reflect"

func DerefValue(rv reflect.Value) reflect.Value {
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	return rv
}
