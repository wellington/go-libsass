package libsass

import (
	"io"

	"github.com/tom-un/go-libsass/libs"
)

// ToScss converts Sass to Scss with libsass sass2scss.h
func ToScss(r io.Reader, w io.Writer) error {
	return libs.ToScss(r, w)
}
