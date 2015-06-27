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

type HandlerFunc func(v interface{}, req SassValue, res *SassValue) error

func Handler(h HandlerFunc) libs.SassCallback {
	return func(v interface{}, usv libs.UnionSassValue, rsv *libs.UnionSassValue) error {
		if *rsv == nil {
			*rsv = libs.SassMakeNil()
		}
		req := SassValue{value: usv}
		res := SassValue{value: *rsv}
		err := h(v, req, &res)
		*rsv = res.Val()

		return err
	}
}

var _ libs.SassCallback = TestCallback

// TestCallback implements libs.SassCallback. TestCallback is a useful
// place to start when developing new handlers.
var TestCallback = testCallback(func(_ interface{}, _ SassValue, _ *SassValue) error {
	return nil
})

func testCallback(h HandlerFunc) libs.SassCallback {
	return func(v interface{}, _ libs.UnionSassValue, _ *libs.UnionSassValue) error {
		return nil
	}
}

// Error takes a Go error and returns a libsass Error
func Error(err error) SassValue {
	return SassValue{value: libs.SassMakeError(err.Error())}
}

// Warn takes a string and causes a warning in libsass
func Warn(s string) SassValue {
	//return C.sass_make_warning(C.CString(s))
	return SassValue{value: libs.SassMakeWarning(s)}
}

// WarnHandler captures Sass warnings and redirects to stdout
func WarnHandler(v interface{}, csv libs.UnionSassValue, rsv *libs.UnionSassValue) error {
	var s string
	Unmarshal(SassValue{value: csv}, &s)
	log.Println("WARNING: " + s)

	r, _ := Marshal("")
	*rsv = r.Val()
	return nil
}

func init() {
	RegisterHandler("@warn", WarnHandler)
}
