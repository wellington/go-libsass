package libsass

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func ExampleCompiler_stdin() {

	src := bytes.NewBufferString(`div { p { color: red; } }`)

	comp, err := New(os.Stdout, src)
	if err != nil {
		log.Fatal(err)
	}
	err = comp.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// div p {
	//   color: red; }
	//

}

func ExampleComipler_sass() {
	src := bytes.NewBufferString(`
html
  font-family: 'MonoSocial'
`)
	comp, err := New(os.Stdout, src, WithSyntax(SassSyntax))
	if err != nil {
		log.Fatal(err)
	}
	err = comp.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// html {
	//   font-family: 'MonoSocial'; }

}

func TestCompiler_path(t *testing.T) {
	var dst bytes.Buffer

	comp, err := New(&dst, nil, Path("test/scss/basic.scss"))
	if err != nil {
		t.Fatal(err)
	}
	err = comp.Run()
	if err != nil {
		t.Fatal(err)
	}

	e := `div p {
  color: red; }
`
	if e != dst.String() {
		t.Errorf("got: %s wanted: %s", dst.String(), e)
	}

	if e := 1; len(comp.Imports()) != e {
		t.Errorf("got: %d wanted: %d", len(comp.Imports()), e)
	}

}

func TestCompiler_path_withpathsresolver(t *testing.T) {
	var dst bytes.Buffer

	absArgs := [][]string{}

	comp, err := New(
		&dst,
		nil,
		Path("test/scss/file.scss"),
		ImportsOption(NewImportsWithResolver(func(importedUrl string, importerAbsPath string) (string, string, bool) {
			absArgs = append(absArgs, []string{importedUrl, importerAbsPath})
			if importedUrl == "a" {
				return "test/scss/a.scss", ".a { color: #aaaaaa }\n@import 'b';", true
			} else if importedUrl == "b" {
				return "test/scss/b.scss", ".b { color: #bbbbbb }", true
			}
			return "", "", false
		})),
	)

	if err != nil {
		t.Fatal(err)
	}
	err = comp.Run()
	if err != nil {
		t.Fatal(err)
	}

	e := `.a {
  color: #aaaaaa; }

.b {
  color: #bbbbbb; }
`
	if e != dst.String() {
		t.Errorf("got: %s wanted: %s", dst.String(), e)
	}

	if e := 3; len(comp.Imports()) != e {
		t.Errorf("got: %d wanted: %d", len(comp.Imports()), e)
	}

	expectedAbsArgs := `[[a test/scss/file.scss] [b a]]`
	if absArgsStr := fmt.Sprintf("%+v", absArgs); expectedAbsArgs != absArgsStr {
		t.Errorf("abs args got: %s wanted: %s", absArgsStr, expectedAbsArgs)
	}
}

func TestCompiler_path_withabsresolver(t *testing.T) {
	var dst bytes.Buffer

	absArgs := [][]string{}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	comp, err := New(
		&dst,
		nil,
		Path("test/scss/file.scss"),
		ImportsOption(NewImportsWithAbsResolver(func(importedUrl string, importerAbsPath string) (string, string, bool) {
			// convert the path to a stable representation
			// (convert abs path to a relative path, with the prefix $CWD/)
			var pathToSave string
			if filepath.IsAbs(importerAbsPath) {
				pathToSave, err = filepath.Rel(wd, importerAbsPath)
				if err != nil {
					t.Fatal(err)
				}
				pathToSave = "$CWD/" + pathToSave
			} else {
				pathToSave = importerAbsPath
			}

			// save the args
			absArgs = append(absArgs, []string{
				importedUrl,
				filepath.ToSlash(pathToSave),
			})

			// handle magic modules "a" and "b"
			if importedUrl == "a" {
				return wd + filepath.FromSlash("/test/scss/a.scss"), ".a { color: #aaaaaa }\n@import 'b';", true
			} else if importedUrl == "b" {
				return wd + filepath.FromSlash("/test/scss/b.scss"), ".b { color: #bbbbbb }", true
			}
			return "", "", false
		})),
	)

	if err != nil {
		t.Fatal(err)
	}
	err = comp.Run()
	if err != nil {
		t.Fatal(err)
	}

	e := `.a {
  color: #aaaaaa; }

.b {
  color: #bbbbbb; }
`
	if e != dst.String() {
		t.Errorf("got: %s wanted: %s", dst.String(), e)
	}

	if e := 3; len(comp.Imports()) != e {
		t.Errorf("got: %d wanted: %d", len(comp.Imports()), e)
	}

	expectedAbsArgs := `[[a $CWD/test/scss/file.scss] [b $CWD/test/scss/a.scss]]`
	if absArgsStr := fmt.Sprintf("%+v", absArgs); expectedAbsArgs != absArgsStr {
		t.Errorf("abs args got: %s wanted: %s", absArgsStr, expectedAbsArgs)
	}
}
