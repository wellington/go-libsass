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

type SassImporter *C.struct_Sass_Importer
type SassImporterList C.Sass_Importer_List

// SassMakeImporterList maps to C.sass_make_importer_list
func SassMakeImporterList(gol int) SassImporterList {
	l := C.size_t(gol)
	cimp := C.sass_make_importer_list(l)
	return (SassImporterList)(cimp)
}

type SassFileContext *C.struct_Sass_File_Context

// SassMakeFileContext maps to C.sass_make_file_context
func SassMakeFileContext(gos string) SassFileContext {
	s := C.CString(gos)
	fctx := C.sass_make_file_context(s)
	return (SassFileContext)(fctx)
}

type SassDataContext *C.struct_Sass_Data_Context

func SassMakeDataContext(gos string) SassDataContext {
	s := C.CString(gos)
	dctx := C.sass_make_data_context(s)
	return (SassDataContext)(dctx)
}

func SassDeleteDataContext(dc SassDataContext) {
	C.sass_delete_data_context(dc)
}

type SassOptions *C.struct_Sass_Options

type SassContext *C.struct_Sass_Context

func SassDataContextGetContext(godc SassDataContext) SassContext {
	opts := C.sass_data_context_get_context(godc)
	return (SassContext)(opts)
}

func SassFileContextSetOptions(gofc SassFileContext, goopts SassOptions) {
	C.sass_file_context_set_options(gofc, goopts)
}

func SassFileContextGetContext(gofc SassFileContext) SassContext {
	opts := C.sass_file_context_get_context(gofc)
	return (SassContext)(opts)
}

func SassFileContextGetOptions(gofc SassFileContext) SassOptions {
	fcopts := C.sass_file_context_get_options(gofc)
	return (SassOptions)(fcopts)
}

func SassDataContextGetOptions(godc SassDataContext) SassOptions {
	dcopts := C.sass_data_context_get_options(godc)
	return (SassOptions)(dcopts)
}

func SassDataContextSetOptions(godc SassDataContext, goopts SassOptions) {
	C.sass_data_context_set_options(godc, goopts)
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

func SassOptionSetCHeaders(gofc SassOptions, goimp SassImporterList) {
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

func SassOptionSetPrecision(goopts SassOptions, i int) {
	C.sass_option_set_precision(goopts, C.int(i))
}

func SassOptionSetOutputStyle(goopts SassOptions, i int) {
	C.sass_option_set_output_style(goopts, uint32(i))
}

func SassOptionSetSourceComments() {

}

func SassOptionSetSourceMapEmbed() {

}

func SassOptionSetSourceMapContents() {

}

func SassOptionSetOmitSourceMapURL() {

}

func SassMakeImporter() {}

type SassImportEntry C.Sass_Import_Entry

func SassMakeImport(path string, base string, source string, srcmap string) SassImportEntry {
	impent := C.sass_make_import(C.CString(path), C.CString(base),
		C.CString(source), C.CString(srcmap))
	return (SassImportEntry)(impent)
}

type SassImporterFN C.Sass_Importer_Fn

func SassImporterGetFunction(goimp SassImporter) SassImporterFN {
	impfn := C.sass_importer_get_function(goimp)
	return (SassImporterFN)(impfn)
}

func SassImporterGetListEntry() {}

func SassOptionSetCImporters(goopts SassOptions, golst SassImporterList) {
	C.sass_option_set_c_importers(goopts, golst)
}

func SassOptionSetCFunctions() {

}
