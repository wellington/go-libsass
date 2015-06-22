package main

// #cgo pkg-config: --cflags --libs libsass
// #cgo LDFLAGS: -lsass -lstdc++ -ldl -lm
// #include "sass_context.h"
import "C"
import (
	"fmt"

	"github.com/wellington/go-libsass/libs"
)

func main() {
	src := "blah.scss"
	cheads := libs.SassMakeImporterList(1)

	gofc := libs.SassMakeFileContext(src)
	gofcopts := libs.SassFileContextGetOptions(gofc)
	libs.SassOptionSetCHeaders(gofcopts, cheads)

	gocc := libs.SassFileContextGetContext(gofc)
	gocomp := libs.SassMakeFileCompiler(gofc)

	libs.SassCompilerParse(gocomp)
	libs.SassCompilerExecute(gocomp)
	gostr := libs.SassContextGetOutputString(gocc)
	fmt.Println(gostr)
}
