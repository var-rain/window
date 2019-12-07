package windows

import "syscall"

var (
	ole                = syscall.NewLazyDLL("ole32.dll")
	procCoInitialize   = ole.NewProc("CoInitialize")
	procCoInitializeEx = ole.NewProc("CoInitializeEx")
	procCoUninitialize = ole.NewProc("CoUninitialize")
)

func CoInitializeEx(coInit uintptr) HRESULT {
	ret, _, _ := procCoInitializeEx.Call(
		0,
		coInit)

	switch uint32(ret) {
	case E_INVALIDARG:
		panic("CoInitializeEx failed with E_INVALIDARG")
	case E_OUTOFMEMORY:
		panic("CoInitializeEx failed with E_OUTOFMEMORY")
	case E_UNEXPECTED:
		panic("CoInitializeEx failed with E_UNEXPECTED")
	}

	return HRESULT(ret)
}

func CoInitialize() {
	_, _, _ = procCoInitialize.Call(0)
}

func CoUninitialize() {
	_, _, _ = procCoUninitialize.Call()
}
