package main

//#cgo LDFLAGS: -L. -ldebuglib
//#include "debuglib.h"
import "C"

func main() {
	// on line 8 place a breakpoint to debug
	value := C.myFunction(1, 2)
	println(value)
}
