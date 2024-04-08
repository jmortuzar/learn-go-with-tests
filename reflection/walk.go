package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v, fn)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res, fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

func walkValue(value reflect.Value, fn func(input string)) {
	walk(value.Interface(), fn)
}
