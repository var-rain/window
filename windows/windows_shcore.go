package windows

import (
	"syscall"
	"unsafe"
)

var (
	core                     = syscall.NewLazyDLL("Shcore.dll")
	getScaleFactorForMonitor = core.NewProc("GetScaleFactorForMonitor")
)

func GetScaleFactorForMonitor(hMonitor HMONITOR, scale *int) HRESULT {
	ret, _, _ := getScaleFactorForMonitor.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(scale)))

	return HRESULT(ret)
}
