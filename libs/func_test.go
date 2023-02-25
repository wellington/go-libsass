package libs

import (
	"fmt"
	"testing"
)

func TestBindFuncs(t *testing.T) {
	fc := SassMakeFileContext("./")
	opts := SassFileContextGetOptions(fc)

	cookies := []Cookie{
		Cookie{
			Sign: "hi",
			Fn: func(v interface{}, usv UnionSassValue, rsv *UnionSassValue) error {
				res := MakeString("test-me")
				*rsv = res
				return nil
			},
		},
	}
	fns := []SassFunc{}

	for i, c := range cookies {
		fns = append(fns, SassMakeFunction(c.Sign, i))
	}

	for _, fn := range fns {
		fmt.Printf("%#v\n", SassGetFunc(fn))
	}

	return
	ids := BindFuncs(opts, cookies)

	for _, id := range ids {
		if globalFuncs.Get(id) == nil {
			t.Errorf("id not found: %d", id)
		}
	}

}
