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
	run("blah.scss")
	run("error.scss")
}

func run(path string) {

	cheads := libs.SassMakeImporterList(1)

	gofc := libs.SassMakeFileContext(path)
	gofcopts := libs.SassFileContextGetOptions(gofc)
	libs.SassOptionSetCHeaders(gofcopts, cheads)

	goctx := libs.SassFileContextGetContext(gofc)
	gocomp := libs.SassMakeFileCompiler(gofc)
	defer libs.SassDeleteCompiler(gocomp)
	libs.SassCompilerParse(gocomp)
	libs.SassCompilerExecute(gocomp)
	gostr := libs.SassContextGetOutputString(goctx)
	fmt.Println(gostr)
	errStatus := libs.SassContextGetErrorStatus(goctx)
	if errStatus > 0 {
		fmt.Println("error:", libs.SassContextGetErrorJSON(goctx))
	}
}
