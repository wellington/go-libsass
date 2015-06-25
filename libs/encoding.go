package libs

// #include "sass_context.h"
import "C"

type UnionSassValue *C.union_Sass_Value

func NewUnionSassValue() UnionSassValue {
	return &C.union_Sass_Value{}
}
