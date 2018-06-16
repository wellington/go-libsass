package libsass

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestUnicode(t *testing.T) {
	t.Skip("unicode can not be stored in variables")
	var dst bytes.Buffer

	src := bytes.NewBufferString(`
$test: "\f00d";

p {
	a {
		content: $test;
	}

	b {
		content: "\f00c";
	}
}
`)

	comp, err := New(&dst, src)
	if err != nil {
		t.Error("Expected to create compiler successfully")
	}

	err = comp.Run()
	if err != nil {
		t.Error("Expected to run the compiler successfully")
	}

	css := dst.String()

	if !strings.Contains(css, "\\f00c") { // This one is passing
		t.Errorf("Expected to find `%s` in compiled css", "\\f00c")
	}

	if !strings.Contains(css, "\\f00d") { // This one isn't
		t.Errorf("Expected to find `%s` in compiled css", "\\f00d")
	}

	fmt.Println(dst.String())
}
