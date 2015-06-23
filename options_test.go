package context

import (
	"bytes"
	"testing"
)

func TestPrecision(t *testing.T) {
	t.Skip("precision does not work")
	in := bytes.NewBufferString(`a { height: 1.1111px; }`)

	var out bytes.Buffer
	ctx := Context{
		Precision: 3,
	}
	err := ctx.Compile(in, &out)
	if err != nil {
		t.Fatal(err)
	}

	e := `a {
  height: 1.111px; }
`
	if e != out.String() {
		t.Errorf("got:\n%s\nwanted:\n%s\n", out.String(), e)
	}

}
