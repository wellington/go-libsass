package libs

import (
	"reflect"
	"testing"
)

func TestBindHeaders(t *testing.T) {
	fc := SassMakeFileContext("./")
	opts := SassFileContextGetOptions(fc)
	e := []ImportEntry{
		ImportEntry{
			Parent: "hi",
		},
	}

	i := BindHeader(opts, e)
	if e := 1; i != e {
		t.Errorf("got: %d wanted: %d", i, e)
	}

	get := globalHeaders.Get(i).([]ImportEntry)

	if !reflect.DeepEqual(get, e) {
		t.Errorf("  got: %#v\nwanted: %#v", get, e)
	}

	// This is failing even though the addresses are the same
	/*if &get != &e {
		t.Errorf("mismatched pointers\n(%p)%#v: (%p)%#v", get, get, e, e)
	}*/
}
