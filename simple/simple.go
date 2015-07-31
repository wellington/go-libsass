package main

import (
	"fmt"

	"github.com/wellington/go-libsass/libs"
)

func main() {
	separate()
	go separate()
}

func separate() {
	godc := libs.SassMakeDataContext("div { p { color: red; } }")
	gocompiler := libs.SassMakeDataCompiler(godc)
	libs.SassCompilerParse(gocompiler)
	fmt.Println("successfully executed")
}
