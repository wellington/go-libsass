package context

import (
	"testing"

	"github.com/wellington/go-libsass/libs"
)

func TestSampleCB(t *testing.T) {
	ctx := NewContext()
	ctx.Cookies = make([]Cookie, 1)
	usv, err := Marshal(float64(1))
	if err != nil {
		t.Error(err)
	}
	var rsv libs.UnionSassValue
	err = SampleCB(ctx, usv.Val(), &rsv)
	if err != nil {
		t.Fatal(err)
	}

	var b bool
	err = Unmarshal(SassValue{value: rsv}, &b)
	if err != nil {
		t.Error(err)
	}
	if e := false; b != e {
		t.Errorf("wanted: %t got: %t", e, b)
	}
}

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
