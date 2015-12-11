// +build dev

package libs

// #cgo CPPFLAGS: -DUSE_LIBSASS
// #cgo CPPFLAGS: -I../libsass-build -I../libsass-build/include
// #cgo LDFLAGS: -lsass
//
import "C"
