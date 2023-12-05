package libsass

import "github.com/tom-un/go-libsass/libs"

// Version reports libsass version information
func Version() string {
	return libs.Version()
}
