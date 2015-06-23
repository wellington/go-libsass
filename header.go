package context

// #include <stdio.h>
// #include "sass_context.h"
//
// extern struct Sass_Import** HeaderBridge(void* cookie);
//
// Sass_Import_List SassHeaders(const char* cur_path, Sass_Importer_Entry cb, struct Sass_Compiler* comp)
// {
//   void* cookie = sass_importer_get_cookie(cb);
//   Sass_Import_List list = HeaderBridge(cookie);
//   return list;
// }
//

import "C"

import (
	"strconv"
	"sync"

	"github.com/wellington/go-libsass/libs"
)

func (ctx *Context) SetHeaders(opts libs.SassOptions) {
	// Push the headers into the local array
	for _, gh := range globalHeaders {
		ctx.Headers.Add(gh)
	}

	// Loop through headers creating ImportEntry
	entries := make([]libs.ImportEntry, ctx.Headers.Len())
	for i, ent := range ctx.Headers.h {
		uniquename := "hdr" + strconv.FormatInt(int64(i), 10)
		entries[i] = libs.ImportEntry{
			// Each entry requires a unique identifier
			// https://github.com/sass/libsass/issues/1292
			uniquename,
			ent.Content,
			"",
		}
	}
	libs.BindHeader(opts, entries)
}

type Header struct {
	Content string
}

type Headers struct {
	sync.RWMutex
	h []Header
}

func (h *Headers) Add(s string) {
	h.Lock()
	defer h.Unlock()

	h.h = append(h.h, Header{
		Content: s,
	})
}

func (h *Headers) Len() int {
	return len(h.h)
}

var globalHeaders []string

// RegisterHeader fifo
func RegisterHeader(body string) {
	globalHeaders = append(globalHeaders, body)
}
