// +build !windows
package libsass

import (
	"bytes"
	"testing"
)

func TestLibsassError(t *testing.T) {

	in := bytes.NewBufferString(`div {
  color: red(blue, purple);
}`)

	var out bytes.Buffer
	ctx := NewContext()

	ctx.Funcs.Add(Func{
		Sign: "foo()",
		Fn:   TestCallback,
		Ctx:  &ctx,
	})
	err := ctx.Compile(in, &out)

	if err == nil {
		t.Error("No error thrown for incorrect arity")
	}

	if e := "wrong number of arguments (2 for 1) for `red'"; e != ctx.err.Message {
		t.Errorf("wanted:%s\ngot:%s\n", e, ctx.err.Message)
	}
	e := `Error > stdin:2
wrong number of arguments (2 for 1) for ` + "`" + `red'
div {
  color: red(blue, purple);
}
`
	if e != err.Error() {
		t.Errorf("wanted:\n%s\ngot:\n%s\n", e, err)
	}
}
