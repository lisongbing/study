package main

import (
	"syscall"
	"unsafe"
)

func main() {
	r0, _, e1 := syscall.Syscall(syscall.OPEN_ALWAYS, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)

}
