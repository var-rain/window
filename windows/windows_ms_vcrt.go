package windows

import "syscall"

var (
	crt        = syscall.NewLazyDLL("msvcrt.dll")
	procMemCpy = crt.NewProc("memcpy")
	procStrLen = crt.NewProc("strlen")
)

func CopyMemory(dst uintptr, src uintptr, length int) {
	_, _, _ = procMemCpy.Call(dst, src, uintptr(uint32(length)))
}

func StrLen(p uintptr) uintptr {
	rc, _, _ := procStrLen.Call(p)
	return rc
}
