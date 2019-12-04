package win

import (
	"syscall"
	"unsafe"
)

var (
	aut                = syscall.NewLazyDLL("oleaut32.dll")
	procSysAllocString = aut.NewProc("SysAllocString")
	procSysFreeString  = aut.NewProc("SysFreeString")
	procSysStringLen   = aut.NewProc("SysStringLen")
	procVariantInit    = aut.NewProc("VariantInit")
)

func VariantInit(v *VARIANT) {
	hr, _, _ := procVariantInit.Call(uintptr(unsafe.Pointer(v)))
	if hr != 0 {
		panic("Invoke VariantInit error.")
	}
	return
}

func SysAllocString(v string) (ss *int16) {
	pss, _, _ := procSysAllocString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v))))
	ss = (*int16)(unsafe.Pointer(pss))
	return
}

func SysFreeString(v *int16) {
	hr, _, _ := procSysFreeString.Call(uintptr(unsafe.Pointer(v)))
	if hr != 0 {
		panic("Invoke SysFreeString error.")
	}
	return
}

func SysStringLen(v *int16) uint {
	l, _, _ := procSysStringLen.Call(uintptr(unsafe.Pointer(v)))
	return uint(l)
}
