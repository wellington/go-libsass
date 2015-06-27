package libs

// #cgo CFLAGS: -O2 -fPIC
// #cgo CPPFLAGS: -w
// #cgo CXXFLAGS: -g -std=c++0x -pedantic -Wno-c++11-extensions -O2 -fPIC
// #cgo darwin linux LDFLAGS: -ldl
// #cgo LDFLAGS: -lstdc++ -lm
//
// #include "sass_context.h"

// extern struct Sass_Import** HeaderBridge(void* cookie);
//
//
// #//for C.free
// #include "stdlib.h"
//
// #cgo pkg-config: --cflags --libs libsass
// #cgo LDFLAGS: -lsass -lstdc++ -ldl -lm
// #include "sass_context.h"
import "C"
import (
	"errors"
	"image/color"
	"reflect"
	"strings"
	"unsafe"
)

type SassImporter *C.struct_Sass_Importer
type SassImporterList C.Sass_Importer_List

// SassMakeImporterList maps to C.sass_make_importer_list
func SassMakeImporterList(gol int) SassImporterList {
	l := C.size_t(gol)
	cimp := C.sass_make_importer_list(l)
	return (SassImporterList)(cimp)
}

type ImportEntry struct {
	Parent string
	Path   string
	Source string
	SrcMap string
}

func GetEntry(es []ImportEntry, parent string, path string) (string, error) {
	for _, e := range es {
		if parent == e.Parent && path == e.Path {
			return e.Source, nil
		}
	}
	return "", errors.New("entry not found")
}

//export HeaderBridge
func HeaderBridge(ptr unsafe.Pointer) C.Sass_Import_List {
	entries := *(*[]ImportEntry)(ptr)

	cents := C.sass_make_import_list(C.size_t(len(entries)))

	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cents)),
		Len:  len(entries), Cap: len(entries),
	}
	goents := *(*[]C.Sass_Import_Entry)(unsafe.Pointer(&hdr))

	for i, ent := range entries {
		// Each entry needs a unique name
		cent := C.sass_make_import_entry(
			C.CString(ent.Path),
			C.CString(ent.Source),
			C.CString(ent.SrcMap))
		// There is a function for modifying an import list, but a proper
		// slice might be more useful.
		// C.sass_import_set_list_entry(cents, C.size_t(i), cent)
		goents[i] = cent
	}

	return cents
}

// ImporterBridge is called by C to pass Importer arguments into Go land. A
// Sass_Import is returned for libsass to resolve.
//
//export ImporterBridge
func ImporterBridge(url *C.char, prev *C.char, ptr unsafe.Pointer) C.Sass_Import_List {
	entries := *(*[]ImportEntry)(ptr)
	parent := C.GoString(prev)
	rel := C.GoString(url)
	list := C.sass_make_import_list(1)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(list)),
		Len:  1, Cap: 1,
	}

	golist := *(*[]C.Sass_Import_Entry)(unsafe.Pointer(&hdr))
	if body, err := GetEntry(entries, parent, rel); err == nil {
		ent := C.sass_make_import_entry(url, C.CString(body), nil)
		cent := (C.Sass_Import_Entry)(ent)
		golist[0] = cent
	} else if strings.HasPrefix(rel, "compass") {
		ent := C.sass_make_import_entry(url, C.CString(""), nil)
		cent := (C.Sass_Import_Entry)(ent)
		golist[0] = cent
	} else {
		ent := C.sass_make_import_entry(url, nil, nil)
		cent := (C.Sass_Import_Entry)(ent)
		golist[0] = cent
	}

	return list
}

type SassImportList C.Sass_Import_List

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

func SassOptionSetSourceComments(goopts SassOptions, b bool) {
	C.sass_option_set_source_comments(goopts, C.bool(b))
}

func SassOptionSetIncludePath(goopts SassOptions, path string) {
	C.sass_option_set_include_path(goopts, C.CString(path))
}

func SassOptionSetSourceMapEmbed() {

}

func SassOptionSetSourceMapContents() {

}

func SassOptionSetOmitSourceMapURL() {

}

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

func SassMakeOptions() SassOptions {
	opts := C.sass_make_options()
	return (SassOptions)(opts)
}

func SassImporterGetListEntry() {}

func SassMakeImporter(fn SassImporterFN, priority int, v interface{}) (SassImporter, error) {
	vv := reflect.ValueOf(v).Elem()
	if !vv.CanAddr() {
		return nil, errors.New("can not take address of")
	}
	// TODO: this will leak memory, the interface must be freed manually
	lst := C.sass_make_importer(fn, C.double(priority), unsafe.Pointer(vv.Addr().Pointer()))
	return (SassImporter)(lst), nil
}

func SassImporterSetListEntry(golst SassImporterList, idx int, ent SassImporter) {
	C.sass_importer_set_list_entry(golst, C.size_t(idx), ent)
}

func SassOptionSetCImporters(goopts SassOptions, golst SassImporterList) {
	C.sass_option_set_c_importers(goopts, golst)
}

func SassOptionSetCFunctions() {

}

// types
func MakeNil() UnionSassValue {
	return C.sass_make_null()
}

func MakeBoolean(b bool) UnionSassValue {
	cb := C.bool(b)
	return C.sass_make_boolean(cb)
}

func MakeError(s string) UnionSassValue {
	return C.sass_make_error(C.CString(s))
}

func MakeWarning(s string) UnionSassValue {
	return C.sass_make_warning(C.CString(s))
}

func MakeBool(b bool) UnionSassValue {
	return C.sass_make_boolean(C.bool(b))
}

func MakeString(s string) UnionSassValue {
	return C.sass_make_string(C.CString(s))
}

// TODO: validate unit
func MakeNumber(f float64, unit string) UnionSassValue {
	return C.sass_make_number(C.double(f), C.CString(unit))
}

// TODO: accept actual color object?
func MakeColor(c color.RGBA) UnionSassValue {
	return C.sass_make_color(C.double(c.R), C.double(c.G),
		C.double(c.B), C.double(c.A))
}

func MakeList(len int) UnionSassValue {
	return C.sass_make_list(C.size_t(len), C.SASS_COMMA)
}

func IsNil(usv UnionSassValue) bool {
	return bool(C.sass_value_is_null(usv))
}

func IsBool(usv UnionSassValue) bool {
	return bool(C.sass_value_is_boolean(usv))
}

func IsString(usv UnionSassValue) bool {
	return bool(C.sass_value_is_string(usv))
}

func IsColor(usv UnionSassValue) bool {
	return bool(C.sass_value_is_color(usv))
}

func IsNumber(usv UnionSassValue) bool {
	return bool(C.sass_value_is_number(usv))
}

func IsList(usv UnionSassValue) bool {
	return bool(C.sass_value_is_list(usv))
}

func IsError(usv UnionSassValue) bool {
	return bool(C.sass_value_is_error(usv))
}

func Len(usv UnionSassValue) int {
	switch {
	case IsList(usv):
		return int(C.sass_list_get_length(usv))
	}
	panic("call of len on unknown type")
}

func SassListGetVal(usv UnionSassValue, idx int) UnionSassValue {
	rsv := C.sass_list_get_value(usv, C.size_t(idx))
	return rsv
}

func SassListSetVal(usv UnionSassValue, idx int, item UnionSassValue) {
	C.sass_list_set_value(usv, C.size_t(idx), item)
}

func String(usv UnionSassValue) string {
	c := C.sass_string_get_value(usv)
	gc := C.GoString(c)
	return gc
}

func Float(usv UnionSassValue) float64 {
	f := C.sass_number_get_value(usv)
	return float64(f)
}

func Unit(usv UnionSassValue) string {
	return C.GoString(C.sass_number_get_unit(usv))
}

func Bool(usv UnionSassValue) bool {
	b := C.sass_boolean_get_value(usv)
	return bool(b)
}

func Color(usv UnionSassValue) color.Color {
	return color.RGBA{
		R: uint8(C.sass_color_get_r(usv)),
		G: uint8(C.sass_color_get_g(usv)),
		B: uint8(C.sass_color_get_b(usv)),
		A: uint8(C.sass_color_get_a(usv)),
	}
}
