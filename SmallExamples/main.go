package main

/* Go Library, to build as a .so -> go build -buildmode=c-shared -o _go-add.so */

import (
	"C"
	"strconv"
)

//export add_test
func add_test(value int) int {
	var stringValue string = strconv.Itoa(value)
	print("\nAdding the value supplied from Python : ", stringValue)
	print("\nTo 42 from go ... equals -> ")
	return value + 42
}

//export string_test
func string_test(text *C.char) *C.char {
	var return_string string = "Hello from Go !!"
	return_string += "\n" + C.GoString(text)
	return C.CString(return_string)
}

func main() {}
