package context

// The use of //export prevents being able to define any C code in the
// preamble of that file.  Export
// defines additional C code necessary for the context<->sass_context bridge.
// See: http://golang.org/cmd/cgo/#hdr-C_references_to_Go

// #include "sass_context.h"
//
import "C"
import (
	"log"
	"unsafe"
)

// GoBridge is exported to C for linking libsass to Go.  This function
// adheres to the interface provided by libsass.
//
//export GoBridge
func GoBridge(cargs UnionSassValue, ptr unsafe.Pointer) UnionSassValue {
	// Recover the Cookie struct passed in
	ck := *(*Cookie)(ptr)
	var usv UnionSassValue
	err := ck.Fn(ck.Ctx, cargs, &usv)
	_ = err
	return usv
}

// SampleCB example how a callback is defined
func SampleCB(ctx *Context, usv UnionSassValue, rsv *UnionSassValue) error {
	var sv []interface{}
	Unmarshal(usv, &sv)
	*rsv = C.sass_make_boolean(false)
	return nil
}

// Error takes a Go error and returns a libsass Error
func Error(err error) UnionSassValue {
	return C.sass_make_error(C.CString(err.Error()))
}

// Warn takes a string and causes a warning in libsass
func Warn(s string) UnionSassValue {
	return C.sass_make_warning(C.CString(s))
}

// WarnHandler captures Sass warnings and redirects to stdout
func WarnHandler(ctx *Context, csv UnionSassValue, rsv *UnionSassValue) error {
	var s string
	Unmarshal(csv, &s)
	log.Println("WARNING: " + s)

	r, _ := Marshal("")
	*rsv = r
	return nil
}

func init() {
	RegisterHandler("@warn", WarnHandler)
}
