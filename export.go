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

	"github.com/wellington/go-libsass/libs"
)

// SampleCB example how a callback is defined
func SampleCB(v interface{}, usv libs.UnionSassValue, rsv *libs.UnionSassValue) error {
	var sv []interface{}
	libs.Unmarshal(usv, &sv)
	*rsv = libs.SassMakeBoolean(false)
	return nil
}

// Error takes a Go error and returns a libsass Error
func Error(err error) libs.UnionSassValue {
	return libs.SassMakeError(err.Error())
}

// Warn takes a string and causes a warning in libsass
func Warn(s string) libs.UnionSassValue {
	//return C.sass_make_warning(C.CString(s))
	return libs.SassMakeWarning(s)
}

// WarnHandler captures Sass warnings and redirects to stdout
func WarnHandler(v interface{}, csv libs.UnionSassValue, rsv *libs.UnionSassValue) error {
	var s string
	libs.Unmarshal(csv, &s)
	log.Println("WARNING: " + s)

	r, _ := libs.Marshal("")
	*rsv = r
	return nil
}

func init() {
	RegisterHandler("@warn", WarnHandler)
}
