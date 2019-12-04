package win

import (
	"errors"
	"syscall"
	"unsafe"
)

func AppendMenu(hMenu HMENU, Flags uint32, IdNewItem uintptr, NewItem string) error {
	var err error
	var pNewItem *uint16

	if NewItem != "" {
		pNewItem, err = syscall.UTF16PtrFromString(NewItem)
		if err != nil {
			return err
		}
	}

	return _AppendMenu(hMenu, Flags, IdNewItem, pNewItem)
}

func _AppendMenu(hMenu HMENU, Flags uint32, IdNewItem uintptr, NewItem *uint16) (err error) {
	r1, _, e1 := syscall.Syscall6(procAppendMenu.Addr(), 4, uintptr(hMenu), uintptr(Flags), IdNewItem, uintptr(unsafe.Pointer(NewItem)), 0, 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("append menu failed")
		}
	}
	return
}

func CreateMenu() (hMenu HMENU, err error) {
	r1, _, e1 := syscall.Syscall(procCreateMenu.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("create menu failed")
		}
	} else {
		hMenu = HMENU(r1)
	}
	return
}

func CreatePopupMenu() (hMenu HMENU, err error) {
	r1, _, e1 := syscall.Syscall(procCreatePopupMenu.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("create popup menu failed")
		}
	} else {
		hMenu = HMENU(r1)
	}
	return
}

func DestroyMenu(hMenu HMENU) bool {
	r1, _, _ := procDestroyMenu.Call(uintptr(hMenu))
	return r1 != 0
}
