package windows

import (
	"errors"
	"syscall"
	"time"
	"unsafe"
)

func ConnectNamedPipe(hNamedPipe HANDLE, po *OVERLAPPED) (err error) {
	r1, _, e1 := syscall.Syscall(procConnectNamedPipe.Addr(), 2, uintptr(hNamedPipe), uintptr(unsafe.Pointer(po)), 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("connect named pipe failed")
		}
	}
	return
}

func CreateNamedPipe(name string, openMode uint32, pipeMode uint32, maxInstances uint32, outBufferSize uint32, inBufferSize uint32, defaultTimeOut time.Duration, sa *SECURITY_ATTRIBUTES) (h HANDLE, err error) {
	pName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	dto := uint32(uint64(defaultTimeOut) / 1e6)
	h, err = _CreateNamedPipe(pName, openMode, pipeMode, maxInstances, outBufferSize, inBufferSize, dto, sa)
	return
}

func _CreateNamedPipe(pName *uint16, dwOpenMode uint32, dwPipeMode uint32, nMaxInstances uint32, nOutBufferSize uint32, nInBufferSize uint32, nDefaultTimeOut uint32, pSecurityAttributes *SECURITY_ATTRIBUTES) (h HANDLE, err error) {
	r1, _, e1 := syscall.Syscall9(procCreateNamedPipe.Addr(), 8, uintptr(unsafe.Pointer(pName)), uintptr(dwOpenMode), uintptr(dwPipeMode), uintptr(nMaxInstances), uintptr(nOutBufferSize), uintptr(nInBufferSize), uintptr(nDefaultTimeOut), uintptr(unsafe.Pointer(pSecurityAttributes)), 0)
	if h == INVALID_HANDLE_VALUE {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("create named pipe failed")
		}
	} else {
		h = HANDLE(r1)
	}
	return
}
