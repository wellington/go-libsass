// +build dev

package libs

// #cgo CFLAGS: -DUSE_LIBSASS
// #cgo LDFLAGS: -lsass -ldl -lm
//
import "C"
