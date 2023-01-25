package libs

import "testing"

func TestBindFuncs(t *testing.T) {
	fc := SassMakeFileContext("./")
	opts := SassFileContextGetOptions(fc)

	cookies := []Cookie{
		Cookie{
			Sign: "hi",
		},
	}
	ids := BindFuncs(opts, cookies)

	for _, id := range ids {
		if globalFuncs.Get(id) == nil {
			t.Errorf("id not found: %d", id)
		}
	}

	if err := RemoveFuncs(ids); err != nil {
		t.Errorf("%s", err)
	}

}
