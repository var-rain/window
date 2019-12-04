package win

import (
	"errors"
	"syscall"
	"unsafe"
)

func PostQuitMessage(exitCode int32) {
	_, _, _ = procPostQuitMessage.Call(
		uintptr(exitCode))
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin int32, msgFilterMax uint32) int32 {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))

	return int32(ret)
}

func TranslateMessage(msg *MSG) bool {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret != 0

}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := procDispatchMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret

}

func SendMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := procSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func PostMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) bool {
	ret, _, _ := procPostMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret != 0
}

func WaitMessage() bool {
	ret, _, _ := procWaitMessage.Call()
	return ret != 0
}

func RegisterWindowMessage(str string) (message uint32, err error) {
	p, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall(procRegisterWindowMessage.Addr(), 1, uintptr(unsafe.Pointer(p)), 0, 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("register window message failed")
		}
	} else {
		message = uint32(r1)
	}
	return
}
