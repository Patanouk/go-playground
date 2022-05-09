package reflection

import (
	"reflect"
)

func walk(x interface{}, f func(string)) {
	value := getValue(x)

	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), f)
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), f)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(key.Interface(), f)
			walk(value.MapIndex(key).Interface(), f)
		}
	case reflect.Chan:
		for v, ok := value.Recv(); ok; v, ok = value.Recv() {
			walk(v.Interface(), f)
		}
	case reflect.Func:
		call := value.Call(nil)
		for _, res := range call {
			walk(res.Interface(), f)
		}
	case reflect.String:
		f(value.String())
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}
