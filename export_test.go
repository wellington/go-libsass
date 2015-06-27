package context

import (
	"testing"

	"github.com/wellington/go-libsass/libs"
)

func TestRegisterHandler(t *testing.T) {
	l := len(handlers)
	RegisterHandler("foo",
		func(v interface{}, csv libs.UnionSassValue, rsv *libs.UnionSassValue) error {
			u, _ := Marshal(false)
			*rsv = u.Val()
			return nil
		})
	if e := l + 1; len(handlers) != e {
		t.Errorf("got: %d wanted: %d", len(handlers), e)
	}
}
