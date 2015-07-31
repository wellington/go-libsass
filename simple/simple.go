package main

import (
	"fmt"
	"sync"

	"github.com/wellington/go-libsass/libs"
)

func main() {
	separate()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		separate()
		wg.Done()
	}()
	wg.Wait()
}

func separate() {
	godc := libs.SassMakeDataContext("div { p { color: red; } }")
	gocompiler := libs.SassMakeDataCompiler(godc)
	libs.SassCompilerParse(gocompiler)
	fmt.Println("successfully executed")
}
