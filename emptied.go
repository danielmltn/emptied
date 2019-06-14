package utilities

import (
	"reflect"
)


/*
IsEmpty will check if a value is set to default value
provided by the Go language. If so it will return
true. By default it will return true if value passed
that does not match.
*/
func IsEmpty(obj interface{}) (isDefault bool) {

	switch objType := obj.(type) {
	case string:
		if objType == "" {
			isDefault = true
		}
		return

	case uintptr, uint, uint8, uint16, uint32, uint64:
		if reflect.ValueOf(objType).Uint() == 0 {
			isDefault = true
		}
		return

	case int, int8, int16, int32, int64:
		if reflect.ValueOf(objType).Int() == 0 {
			isDefault = true
		}
		return

	case float32, float64:
		if reflect.ValueOf(objType).Float() == 0 {
			isDefault = true
		}
		return

	case complex64, complex128:
		if reflect.ValueOf(objType).Complex() == 0 {
			isDefault = true
		}
		return

	default:
		return
	}

}
