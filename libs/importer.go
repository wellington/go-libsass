package libs

// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
// #include "sass/context.h"
//
// extern struct Sass_Import** ImporterBridge(const char* url, const char* prev, uintptr_t idx);
//
// Sass_Import_List SassImporterPathsHandler(const char* cur_path, Sass_Importer_Entry cb, struct Sass_Compiler* comp)
// {
//   void* cookie = sass_importer_get_cookie(cb);
//   struct Sass_Import* previous = sass_compiler_get_last_import(comp);
//   const char* prev_path = sass_import_get_imp_path(previous);
//   uintptr_t idx = (uintptr_t)cookie;
//   Sass_Import_List list = ImporterBridge(cur_path, prev_path, idx);
//   return list;
// }
//
// Sass_Import_List SassImporterAbsHandler(const char* cur_path, Sass_Importer_Entry cb, struct Sass_Compiler* comp)
// {
//   void* cookie = sass_importer_get_cookie(cb);
//   struct Sass_Import* previous = sass_compiler_get_last_import(comp);
//   const char* prev_abs_path = sass_import_get_abs_path(previous);
//   uintptr_t idx = (uintptr_t)cookie;
//   Sass_Import_List list = ImporterBridge(cur_path, prev_abs_path, idx);
//   return list;
// }
//
//
// #ifndef UINTMAX_MAX
// #  ifdef __UINTMAX_MAX__
// #    define UINTMAX_MAX __UINTMAX_MAX__
// #  endif
// #endif
//
// //size_t max_size = UINTMAX_MAX;
import "C"
import "unsafe"

var globalImports SafeMap

// ImportResolver can be used as a custom import resolver. Return an empty body to
// signal loading the import body from the URL.
type ImportResolver func(url string, prev string) (newURL string, body string, resolved bool)

type ResolverMode int

const (
	ResolverModeImporterUrl ResolverMode = iota
	ResolverModeImporterAbsPath
)

func init() {
	globalImports.init()
}

// BindImporter attaches a custom importer Go function to an import in Sass
func BindImporter(opts SassOptions, resolverMode ResolverMode, resolver ImportResolver) int {

	idx := globalImports.Set(resolver)
	ptr := unsafe.Pointer(&idx)

	var handler unsafe.Pointer
	if resolverMode == ResolverModeImporterAbsPath {
		handler = C.SassImporterAbsHandler
	} else {
		handler = C.SassImporterPathsHandler
	}

	imper := C.sass_make_importer(
		C.Sass_Importer_Fn(handler),
		C.double(0),
		ptr,
	)
	impers := C.sass_make_importer_list(1)
	C.sass_importer_set_list_entry(impers, 0, imper)

	C.sass_option_set_c_importers(
		(*C.struct_Sass_Options)(unsafe.Pointer(opts)),
		impers,
	)
	return idx
}

func RemoveImporter(idx int) error {
	globalImports.Del(idx)
	return nil
}
