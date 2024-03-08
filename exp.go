package exp

import (
	"reflect"
	"unsafe"
)

// FastStructField returns reflect struct fields without allocations.
func FastStructField(v *reflect.Value, i int) (reflect.StructField, error) {
	castedValue := (*reflectValue)(noescape(unsafe.Pointer(v)))
	tt := (*structType)(unsafe.Pointer(castedValue.typ_))

	return tt.Field(i)
}
