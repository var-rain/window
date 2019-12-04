package win

import (
	"syscall"
	"unsafe"
)

var (
	dlg                      = syscall.NewLazyDLL("comdlg32.dll")
	procCommDlgExtendedError = dlg.NewProc("CommDlgExtendedError")
	procGetOpenFileName      = dlg.NewProc("GetOpenFileNameW")
	procGetSaveFileName      = dlg.NewProc("GetSaveFileNameW")
)

func GetOpenFileName(ofn *OPENFILENAME) bool {
	ret, _, _ := procGetOpenFileName.Call(
		uintptr(unsafe.Pointer(ofn)))

	return ret != 0
}

func GetSaveFileName(ofn *OPENFILENAME) bool {
	ret, _, _ := procGetSaveFileName.Call(
		uintptr(unsafe.Pointer(ofn)))

	return ret != 0
}

func CommDlgExtendedError() uint {
	ret, _, _ := procCommDlgExtendedError.Call()

	return uint(ret)
}
