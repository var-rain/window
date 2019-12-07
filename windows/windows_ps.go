package windows

import (
	"syscall"
	"unsafe"
)

var (
	ps                = syscall.NewLazyDLL("psapi.dll")
	procEnumProcesses = ps.NewProc("EnumProcesses")
)

func EnumProcesses(processIds []uint32, cb uint32, bytesReturned *uint32) bool {
	ret, _, _ := procEnumProcesses.Call(
		uintptr(unsafe.Pointer(&processIds[0])),
		uintptr(cb),
		uintptr(unsafe.Pointer(bytesReturned)))

	return ret != 0
}
