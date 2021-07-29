package main

/* Go Library, to build as a .so -> go build -buildmode=c-shared -o _go-add.so */

import (
	"C"
	"strconv"
)

//export add
func add(value int) int {
	var stringValue string = strconv.Itoa(value)
	print("\nAdding the value supplied from Python : ", stringValue)
	print("\nTo 42 from go ... equals -> ")
	return value + 42
}

func main() {
}
