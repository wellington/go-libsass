package context

// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
// #include "sass_context.h"
//
// extern struct Sass_Import** ImporterBridge(const char* url, const char* prev, void* cookie);
//
// Sass_Import_List SassImporterHandler(const char* cur_path, Sass_Importer_Entry cb, struct Sass_Compiler* comp)
// {
//   void* cookie = sass_importer_get_cookie(cb);
//   struct Sass_Import* previous = sass_compiler_get_last_import(comp);
//   const char* prev_path = sass_import_get_path(previous);
//   Sass_Import_List list = ImporterBridge(cur_path, prev_path, cookie);
//   return list;
// }
//
// #ifndef UINTMAX_MAX
// #  ifdef __UINTMAX_MAX__
// #    define UINTMAX_MAX __UINTMAX_MAX__
// #  endif
// #endif
//
// size_t max_size = UINTMAX_MAX;
import "C"
import (
	"errors"
	"io"
	"sync"
	"time"

	"github.com/wellington/go-libsass/libs"
)

// MaxSizeT is safe way of specifying size_t = -1
var MaxSizeT = C.max_size

var (
	ErrImportNotFound = errors.New("Import unreachable or not found")
)

// Import contains Rel and Abs path and a string of the contents
// representing an import.
type Import struct {
	Body  io.ReadCloser
	bytes []byte
	mod   time.Time
}

// ModTime returns modification time
func (i Import) ModTime() time.Time {
	return i.mod
}

// Imports is a map with key of "path/to/file"
type Imports struct {
	sync.RWMutex
	m map[string]Import
}

// Init sets up a new Imports map
func (p *Imports) Init() {
	p.m = make(map[string]Import)
}

// Add registers an import in the context.Imports
func (p *Imports) Add(prev string, cur string, bs []byte) error {
	p.Lock()
	defer p.Unlock()

	im := Import{
		bytes: bs,
		mod:   time.Now(),
	}
	// TODO: align these with libsass name "stdin"
	if len(prev) == 0 || prev == "string" {
		prev = "stdin"
	}
	p.m[prev+":"+cur] = im
	return nil
}

// Del removes the import from the context.Imports
func (p *Imports) Del(path string) {
	p.Lock()
	defer p.Unlock()

	delete(p.m, path)
}

// Get retrieves import bytes by path
func (p *Imports) Get(prev, path string) ([]byte, error) {
	p.RLock()
	defer p.RUnlock()
	imp, ok := p.m[prev+":"+path]
	if !ok {
		return nil, ErrImportNotFound
	}
	return imp.bytes, nil
}

// Update attempts to create a fresh Body from the given path
// Files last modified stamps are compared against import timestamp
func (p *Imports) Update(name string) {
	p.Lock()
	defer p.Unlock()

}

// Len counts the number of entries in context.Imports
func (p *Imports) Len() int {
	return len(p.m)
}

// SetImporter enables custom importer in libsass
func (ctx *Context) SetImporter(opts libs.SassOptions) {

	ggoimps := libs.SassMakeImporterList(1)
	goimp, _ := libs.SassMakeImporter(
		libs.SassImporterFN(C.SassImporterHandler), 0, ctx)
	libs.SassImporterSetListEntry(ggoimps, 0, goimp)
	libs.SassOptionSetCImporters(opts, ggoimps)

}
