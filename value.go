package experiment

import (
	"reflect"
	"unsafe"

	"github.com/slipros/exp/internal/abi"
)

type flag uintptr

type reflectValue struct {
	// typ_ holds the type of the value represented by a Value.
	// Access using the typ method to avoid escape of v.
	typ_ *abi.Type

	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
	ptr unsafe.Pointer

	// flag holds metadata about the value.
	//
	// The lowest five bits give the Kind of the value, mirroring typ.Kind().
	//
	// The next set of bits are flag bits:
	//	- flagStickyRO: obtained via unexported not embedded field, so read-only
	//	- flagEmbedRO: obtained via unexported embedded field, so read-only
	//	- flagIndir: val holds a pointer to the data
	//	- flagAddr: v.CanAddr is true (implies flagIndir and ptr is non-nil)
	//	- flagMethod: v is a method value.
	// If ifaceIndir(typ), code can assume that flagIndir is set.
	//
	// The remaining 22+ bits give a method number for method values.
	// If flag.kind() != Func, code can assume that flagMethod is unset.
	flag

	// A method value represents a curried method invocation
	// like r.Read for some receiver r. The typ+val+flag bits describe
	// the receiver r, but the flag's Kind bits say Func (methods are
	// functions), and the top bits of the flag give the method number
	// in r's type's method table.
}

// structType represents a struct type.
type structType struct {
	abi.StructType
}

// Field returns the i'th struct field.
func (t *structType) Field(i int) (f reflect.StructField, err error) {
	if i < 0 || i >= len(t.Fields) {
		return reflect.StructField{}, ErrFieldIndexOutOfBounds
	}

	p := &t.Fields[i]

	f.Name = p.Name.Name()
	if !p.Name.IsExported() {
		f.PkgPath = t.PkgPath.Name()
	}

	if tag := p.Name.Tag(); tag != "" {
		f.Tag = reflect.StructTag(tag)
	}

	return
}

//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
