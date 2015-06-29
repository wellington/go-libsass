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

// Error takes a Go error and returns a libsass Error
func Error(err error) SassValue {
	return SassValue{value: libs.MakeError(err.Error())}
}

// Warn takes a string and causes a warning in libsass
func Warn(s string) SassValue {
	//return C.sass_make_warning(C.CString(s))
	return SassValue{value: libs.MakeWarning(s)}
}

// WarnHandler captures Sass warnings and redirects to stdout
func WarnHandler(v interface{}, csv SassValue, rsv *SassValue) error {
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
