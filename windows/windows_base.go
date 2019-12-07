package windows

import (
	"fmt"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func FormatMessage(flags uint32, msgSrc interface{}, msgId uint32, langId uint32, args *byte) (string, error) {
	var b [300]uint16
	n, err := _FormatMessage(flags, msgSrc, msgId, langId, &b[0], 300, args)
	if err != nil {
		return "", err
	}
	for ; n > 0 && (b[n-1] == '\n' || b[n-1] == '\r'); n-- {
	}
	return string(utf16.Decode(b[:n])), nil
}

func _FormatMessage(flags uint32, msgSrc interface{}, msgId uint32, langId uint32, buf *uint16, nSize uint32, args *byte) (n uint32, err error) {
	r0, _, e1 := syscall.Syscall9(procFormatMessage.Addr(), 7, uintptr(flags), uintptr(0), uintptr(msgId), uintptr(langId), uintptr(unsafe.Pointer(buf)), uintptr(nSize), uintptr(unsafe.Pointer(args)), 0, 0)
	n = uint32(r0)
	if n == 0 {
		err = fmt.Errorf("format message error: %d", uint32(e1))
	}
	return
}
