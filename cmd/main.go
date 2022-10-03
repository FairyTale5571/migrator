package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/fairytale5571/migrator/internal/app"
)

func main() {}

var funks = map[string]func() string{
	"version": app.Version,
	"migrate": app.Migrate,
}

//export goRVExtensionVersion
func goRVExtensionVersion(output *C.char, outputsize C.size_t) {
	result := C.CString("GRC 1.0")
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func printInArma(output *C.char, outputsize C.size_t, input string) {
	result := C.CString(input)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export goRVExtension
func goRVExtension(output *C.char, outputsize C.size_t, input *C.char) {
	str := C.GoString(input)
	f, ok := funks[str]
	if !ok {
		printInArma(output, outputsize, fmt.Sprintf("Unkown '%v' method", str))
		return
	}
	id := f()
	printInArma(output, outputsize, id)
}
