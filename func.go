package context

// #include <stdlib.h>
// #include "sass_context.h"
//
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/wellington/go-libsass/libs"
)

// Cookie is used for passing context information to libsass.  Cookie is
// passed to custom handlers when libsass executes them through the go
// bridge.
type Cookie struct {
	Sign string
	Fn   libs.SassCallback
	Ctx  interface{}
}

type handler struct {
	sign     string
	callback libs.SassCallback
}

// handlers is the list of registered sass handlers
var handlers []handler

// RegisterHandler sets the passed signature and callback to the
// handlers array.
func RegisterHandler(sign string, callback libs.SassCallback) {
	handlers = append(handlers, handler{sign, callback})
}

func (ctx *Context) SetFunc(opts *C.struct_Sass_Options) {
	cookies := make([]libs.Cookie, len(handlers)+len(ctx.Cookies))
	// Append registered handlers to cookie array
	for i, h := range handlers {
		cookies[i] = libs.Cookie{
			Sign: h.sign,
			Fn:   h.callback,
			Ctx:  ctx,
		}
	}
	for i, h := range ctx.Cookies {
		cookies[i+len(handlers)] = libs.Cookie{
			Sign: h.Sign,
			Fn:   h.Fn,
			Ctx:  ctx,
		}
	}
	// disable garbage collection of cookies. These need to
	// be manually freed in the wrapper
	runtime.SetFinalizer(&cookies, nil)
	gofns := make([]libs.SassFunc, len(cookies))
	for i, cookie := range cookies {
		fn := libs.SassMakeFunction(cookie.Sign,
			unsafe.Pointer(&cookies[i]))
		gofns[i] = fn
	}
	goopts := (libs.SassOptions)(unsafe.Pointer(opts))
	libs.BindFuncs(goopts, gofns)
}
