package main

/*
#include <stdio.h>
#include <stdlib.h>
void print(const char *str) {
   printf("%s", str);
}
*/
import "C"
import (
	"unsafe"
)

func main () {
	s := C.CString("vankichi")
	defer C.free(unsafe.Pointer(s))
	C.print(s)
}
