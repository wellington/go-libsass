package context

import "testing"

func TestSampleCB(t *testing.T) {
	ctx := NewContext()
	ctx.Cookies = make([]Cookie, 1)
	usv, err := Marshal(float64(1))
	if err != nil {
		t.Error(err)
	}
	var rsv UnionSassValue
	err = SampleCB(ctx, usv, &rsv)
	if err != nil {
		t.Fatal(err)
	}

	var b bool
	err = Unmarshal(rsv, &b)
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
		func(c *Context, csv UnionSassValue, rsv *UnionSassValue) error {
			u, _ := Marshal(false)
			*rsv = u
			return nil
		})
	if e := l + 1; len(handlers) != e {
		t.Errorf("got: %d wanted: %d", len(handlers), e)
	}
}
