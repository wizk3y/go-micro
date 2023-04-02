package util

import (
	"reflect"
)

// IsZeroOfUnderlyingType --
func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
