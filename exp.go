package exp

import (
	"reflect"
	"unsafe"
)

// FastStructField returns reflect struct fields without allocations.
// Attention! Not all fields are field in reflect.StructField!
func FastStructField(v *reflect.Value, i int) (reflect.StructField, bool) {
	castedValue := (*reflectValue)(noescape(unsafe.Pointer(v)))
	tt := (*structType)(unsafe.Pointer(castedValue.typ_))

	return tt.Field(i)
}

// FastStructFieldByName returns reflect struct fields without allocations.
// Attention! Not all fields are field in reflect.StructField!
func FastStructFieldByName(v *reflect.Value, name string) (reflect.StructField, bool) {
	castedValue := (*reflectValue)(noescape(unsafe.Pointer(v)))
	tt := (*structType)(unsafe.Pointer(castedValue.typ_))

	return tt.FieldByName(name)
}
