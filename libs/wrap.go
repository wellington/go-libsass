package libs

// #cgo CFLAGS: -O2 -fPIC
// #cgo CPPFLAGS: -w
// #cgo CXXFLAGS: -g -std=c++0x -pedantic -Wno-c++11-extensions -O2 -fPIC
// #cgo darwin linux LDFLAGS: -ldl
// #cgo LDFLAGS: -lstdc++ -lm
//
// #include "sass_context.h"

// #//for C.free
// #include "stdlib.h"
//
// #cgo pkg-config: --cflags --libs libsass
// #cgo LDFLAGS: -lsass -lstdc++ -ldl -lm
// #include "sass_context.h"
import "C"
import "unsafe"

type SassImporter **C.struct_Sass_Importer

// SassMakeImporterList maps to C.sass_make_importer_list
func SassMakeImporterList(gol int) SassImporter {
	l := C.size_t(gol)
	cimp := C.sass_make_importer_list(l)
	return (SassImporter)(cimp)
}

type SassFileContext *C.struct_Sass_File_Context

// SassMakeFileContext maps to C.sass_make_file_context
func SassMakeFileContext(gos string) SassFileContext {
	s := C.CString(gos)
	fctx := C.sass_make_file_context(s)
	return (SassFileContext)(fctx)
}

type SassOptions *C.struct_Sass_Options

type SassContext *C.struct_Sass_Context

func SassFileContextGetContext(gofc SassFileContext) SassContext {
	opts := C.sass_file_context_get_context(gofc)
	return (SassContext)(opts)
}

func SassFileContextGetOptions(gofc SassFileContext) SassOptions {
	fcopts := C.sass_file_context_get_options(gofc)
	return (SassOptions)(fcopts)
}

func SassMakeFileCompiler(gofc SassFileContext) SassCompiler {
	sc := C.sass_make_file_compiler(gofc)
	return (SassCompiler)(sc)
}

type SassCompiler *C.struct_Sass_Compiler

func SassCompilerExecute(c SassCompiler) {
	C.sass_compiler_execute(c)
}

func SassCompilerParse(c SassCompiler) {
	C.sass_compiler_parse(c)
}

func SassDeleteCompiler(c SassCompiler) {
	C.sass_delete_compiler(c)
}

func SassOptionSetCHeaders(gofc SassOptions, goimp SassImporter) {
	C.sass_option_set_c_headers(gofc, goimp)
}

func SassContextGetOutputString(goctx SassContext) string {
	cstr := C.sass_context_get_output_string(goctx)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func SassContextGetErrorJSON(goctx SassContext) string {
	cstr := C.sass_context_get_error_json(goctx)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr)
}

func SassContextGetErrorStatus(goctx SassContext) int {
	return int(C.sass_context_get_error_status(goctx))
}
