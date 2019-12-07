package windows

import (
	"encoding/binary"
	"errors"
	"fmt"
	"syscall"
	"unicode/utf8"
	"unsafe"
)

var (
	user                              = syscall.NewLazyDLL("user32.dll")
	procRegisterClass                 = user.NewProc("RegisterClassExW")
	procCreateWindow                  = user.NewProc("CreateWindowExW")
	procLoadString                    = user.NewProc("LoadStringW")
	procLoadBitmap                    = user.NewProc("LoadBitmapW")
	procLoadImage                     = user.NewProc("LoadImageW")
	procRegisterWindowMessage         = user.NewProc("RegisterWindowMessageW")
	procAppendMenu                    = user.NewProc("AppendMenuW")
	procCreateMenu                    = user.NewProc("CreateMenu")
	procCreatePopupMenu               = user.NewProc("CreatePopupMenu")
	procDestroyMenu                   = user.NewProc("DestroyMenu")
	procAddClipboardFormatListener    = user.NewProc("AddClipboardFormatListener")
	procAdjustWindowRect              = user.NewProc("AdjustWindowRect")
	procAdjustWindowRectEx            = user.NewProc("AdjustWindowRectEx")
	procBeginPaint                    = user.NewProc("BeginPaint")
	procCallNextHookEx                = user.NewProc("CallNextHookEx")
	procCallWindowProc                = user.NewProc("CallWindowProcW")
	procChangeDisplaySettingsEx       = user.NewProc("ChangeDisplaySettingsExW")
	procClientToScreen                = user.NewProc("ClientToScreen")
	procCloseClipboard                = user.NewProc("CloseClipboard")
	procCopyRect                      = user.NewProc("CopyRect")
	procCreateDialogParam             = user.NewProc("CreateDialogParamW")
	procCreateIcon                    = user.NewProc("CreateIcon")
	procCreateWindowEx                = user.NewProc("CreateWindowExW")
	procDefDlgProc                    = user.NewProc("DefDlgProcW")
	procDefWindowProc                 = user.NewProc("DefWindowProcW")
	procDestroyIcon                   = user.NewProc("DestroyIcon")
	procDestroyWindow                 = user.NewProc("DestroyWindow")
	procDialogBoxParam                = user.NewProc("DialogBoxParamW")
	procDispatchMessage               = user.NewProc("DispatchMessageW")
	procDrawIcon                      = user.NewProc("DrawIcon")
	procDrawText                      = user.NewProc("DrawTextW")
	procEmptyClipboard                = user.NewProc("EmptyClipboard")
	procEnableWindow                  = user.NewProc("EnableWindow")
	procEndDialog                     = user.NewProc("EndDialog")
	procEndPaint                      = user.NewProc("EndPaint")
	procEnumChildWindows              = user.NewProc("EnumChildWindows")
	procEnumClipboardFormats          = user.NewProc("EnumClipboardFormats")
	procEnumDisplayMonitors           = user.NewProc("EnumDisplayMonitors")
	procEnumDisplaySettingsEx         = user.NewProc("EnumDisplaySettingsExW")
	procEqualRect                     = user.NewProc("EqualRect")
	procFillRect                      = user.NewProc("FillRect")
	procFindWindowExW                 = user.NewProc("FindWindowExW")
	procFindWindowW                   = user.NewProc("FindWindowW")
	procGetAsyncKeyState              = user.NewProc("GetAsyncKeyState")
	procGetClassName                  = user.NewProc("GetClassNameW")
	procGetClientRect                 = user.NewProc("GetClientRect")
	procGetClipboardData              = user.NewProc("GetClipboardData")
	procGetClipboardFormatName        = user.NewProc("GetClipboardFormatNameW")
	procGetCursorPos                  = user.NewProc("GetCursorPos")
	procGetDC                         = user.NewProc("GetDC")
	procGetDlgItem                    = user.NewProc("GetDlgItem")
	procGetForegroundWindow           = user.NewProc("GetForegroundWindow")
	procGetKeyState                   = user.NewProc("GetKeyState")
	procGetKeyboardState              = user.NewProc("GetKeyboardState")
	procGetMessage                    = user.NewProc("GetMessageW")
	procGetMonitorInfo                = user.NewProc("GetMonitorInfoW")
	procGetSystemMetrics              = user.NewProc("GetSystemMetrics")
	procGetWindowLong                 = user.NewProc("GetWindowLongW")
	procGetWindowLongPtr              = user.NewProc("GetWindowLongW")
	procGetWindowRect                 = user.NewProc("GetWindowRect")
	procGetWindowText                 = user.NewProc("GetWindowTextW")
	procGetWindowTextLength           = user.NewProc("GetWindowTextLengthW")
	procGetWindowTextW                = user.NewProc("GetWindowTextW")
	procGetWindowThreadProcessId      = user.NewProc("GetWindowThreadProcessId")
	procInflateRect                   = user.NewProc("InflateRect")
	procIntersectRect                 = user.NewProc("IntersectRect")
	procInvalidateRect                = user.NewProc("InvalidateRect")
	procIsClipboardFormatAvailable    = user.NewProc("IsClipboardFormatAvailable")
	procIsDialogMessage               = user.NewProc("IsDialogMessageW")
	procIsRectEmpty                   = user.NewProc("IsRectEmpty")
	procIsWindow                      = user.NewProc("IsWindow")
	procIsWindowEnabled               = user.NewProc("IsWindowEnabled")
	procIsWindowVisible               = user.NewProc("IsWindowVisible")
	procLoadCursor                    = user.NewProc("LoadCursorW")
	procLoadIcon                      = user.NewProc("LoadIconW")
	procMapVirtualKey                 = user.NewProc("MapVirtualKeyExW")
	procMessageBox                    = user.NewProc("MessageBoxW")
	procMonitorFromPoint              = user.NewProc("MonitorFromPoint")
	procMonitorFromRect               = user.NewProc("MonitorFromRect")
	procMonitorFromWindow             = user.NewProc("MonitorFromWindow")
	procMoveWindow                    = user.NewProc("MoveWindow")
	procOffsetRect                    = user.NewProc("OffsetRect")
	procOpenClipboard                 = user.NewProc("OpenClipboard")
	procPeekMessage                   = user.NewProc("PeekMessageW")
	procPostMessage                   = user.NewProc("PostMessageW")
	procPostQuitMessage               = user.NewProc("PostQuitMessage")
	procPtInRect                      = user.NewProc("PtInRect")
	procRegisterClassEx               = user.NewProc("RegisterClassExW")
	procRegisterHotKey                = user.NewProc("RegisterHotKey")
	procReleaseCapture                = user.NewProc("ReleaseCapture")
	procReleaseDC                     = user.NewProc("ReleaseDC")
	procRemoveClipboardFormatListener = user.NewProc("RemoveClipboardFormatListener")
	procScreenToClient                = user.NewProc("ScreenToClient")
	procSendInput                     = user.NewProc("SendInput")
	procSendMessage                   = user.NewProc("SendMessageW")
	procSendMessageTimeout            = user.NewProc("SendMessageTimeoutW")
	procSetCapture                    = user.NewProc("SetCapture")
	procSetClipboardData              = user.NewProc("SetClipboardData")
	procSetCursor                     = user.NewProc("SetCursor")
	procSetCursorPos                  = user.NewProc("SetCursorPos")
	procSetFocus                      = user.NewProc("SetFocus")
	procSetForegroundWindow           = user.NewProc("SetForegroundWindow")
	procSetRect                       = user.NewProc("SetRect")
	procSetRectEmpty                  = user.NewProc("SetRectEmpty")
	procSetWinEventHook               = user.NewProc("SetWinEventHook")
	procSetWindowLong                 = user.NewProc("SetWindowLongW")
	procSetWindowLongPtr              = user.NewProc("SetWindowLongW")
	procSetWindowPos                  = user.NewProc("SetWindowPos")
	procSetWindowText                 = user.NewProc("SetWindowTextW")
	procSetWindowsHookEx              = user.NewProc("SetWindowsHookExW")
	procShowWindow                    = user.NewProc("ShowWindow")
	procSubtractRect                  = user.NewProc("SubtractRect")
	procSwapMouseButton               = user.NewProc("SwapMouseButton")
	procToAscii                       = user.NewProc("ToAscii")
	procTranslateAccelerator          = user.NewProc("TranslateAcceleratorW")
	procTranslateMessage              = user.NewProc("TranslateMessage")
	procUnhookWinEvent                = user.NewProc("UnhookWinEvent")
	procUnhookWindowsHookEx           = user.NewProc("UnhookWindowsHookEx")
	procUnionRect                     = user.NewProc("UnionRect")
	procUnregisterHotKey              = user.NewProc("UnregisterHotKey")
	procUpdateWindow                  = user.NewProc("UpdateWindow")
	procVkKeyScanW                    = user.NewProc("VkKeyScanW")
	procVkKeyScanExW                  = user.NewProc("VkKeyScanExW")
	procWaitMessage                   = user.NewProc("WaitMessage")
	setProcessDPIAware                = user.NewProc("SetProcessDPIAware")
)

func newWndProc(proc WNDPROC) uintptr {
	return syscall.NewCallback(proc)
}

func RegisterClass(pWndClass *WNDCLASS) (atom uint16, err error) {
	if pWndClass == nil {
		err = errors.New("pWndClass must not be nil")
		return
	}
	_pClassName, err := syscall.UTF16PtrFromString(pWndClass.PszClassName)
	if err != nil {
		return
	}
	if pWndClass.Menu == nil {
		err = errors.New("can't find menu")
		return
	}
	var Menu uintptr = 70000
	var _pMenuName *uint16 = nil
	switch v := pWndClass.Menu.(type) {
	case uint16:
		Menu = uintptr(v)
	case string:
		_pMenuName, err = syscall.UTF16PtrFromString(v)
		if err != nil {
			return
		}
	default:
		return 0, errors.New("menu's type must be uint16 or string")
	}
	var wc _WNDCLASS
	wc.cbSize = uint32(unsafe.Sizeof(wc))
	wc.style = pWndClass.Style
	wc.pfnWndProcPtr = newWndProc(pWndClass.PfnWndProc)
	wc.cbClsExtra = pWndClass.CbClsExtra
	wc.cbWndExtra = pWndClass.CbWndExtra
	wc.hInstance = pWndClass.HInstance
	wc.hIcon = pWndClass.HIcon
	wc.hCursor = pWndClass.HCursor
	wc.hbrBackground = pWndClass.HbrBackground
	if _pClassName != nil {
		wc.pszMenuName = _pMenuName
	} else {
		wc.pszMenuName = (*uint16)(unsafe.Pointer(Menu))
	}
	wc.pszClassName = _pClassName
	wc.hIconSmall = pWndClass.HIconSmall

	r1, _, e1 := syscall.Syscall(procRegisterClass.Addr(), 1, uintptr(unsafe.Pointer(&wc)), 0, 0)
	n := uint16(r1)
	if n == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("register class failed")
		}
	} else {
		atom = n
	}
	return
}

func MustMessageBox(hWnd HWND, Text string, Caption string, Type uint32) (ret int32) {
	ret = MessageBox(hWnd, Text, Caption, Type)
	return ret
}

func ErrorAssert(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func CreateWindow(ClassName string, WindowName string, Style uint32, ExStyle uint32, X int32, Y int32, Width int32, Height int32, WndParent HWND, Menu HMENU, inst HINSTANCE, Param uintptr) (hWnd HWND, err error) {
	pClassName, err := syscall.UTF16PtrFromString(ClassName)
	if err != nil {
		return
	}
	pWindowName, err := syscall.UTF16PtrFromString(WindowName)
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall12(procCreateWindow.Addr(), 12, uintptr(ExStyle), uintptr(unsafe.Pointer(pClassName)), uintptr(unsafe.Pointer(pWindowName)), uintptr(Style), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(WndParent), uintptr(Menu), uintptr(inst), uintptr(Param))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = errors.New("create window failed")
		}
	} else {
		hWnd = HWND(r1)
	}
	return
}

func _LoadString(Inst HINSTANCE, id uint16, Buffer *uint16, BufferMax int32) (int32, error) {
	r1, _, e1 := syscall.Syscall6(procLoadString.Addr(), 4, uintptr(Inst), uintptr(id), uintptr(unsafe.Pointer(Buffer)), uintptr(BufferMax), 0, 0)
	r := int32(r1)
	if r > 0 {
		return r, nil
	} else {
		wec := ErrorCode(e1)
		if wec != 0 {
			return 0, wec
		} else {
			return 0, errors.New("load string failed")
		}
	}
}

func LoadString(hInstance HINSTANCE, id uint16) (string, error) {
	var err error
	var Len, Len1 int32
	var p *uint16 = nil
	Len, err = _LoadString(hInstance, id, (*uint16)(unsafe.Pointer(&p)), 0)
	if err == nil && Len > 0 {
		Buffer := make([]uint16, Len+1)
		Len1, err = _LoadString(hInstance, id, &Buffer[0], Len+1)
		if err == nil && Len == Len1 {
			Buffer[Len] = 0
			return syscall.UTF16ToString(Buffer), nil
		} else {
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		return "", errors.New("load string failed")
	}
}

func LoadBitmapById(hInst HINSTANCE, id uint16) (HBITMAP, error) {
	r1, _, e1 := syscall.Syscall(procLoadBitmap.Addr(), 2, uintptr(hInst), uintptr(id), 0)
	if r1 != 0 {
		return HBITMAP(r1), nil
	} else {
		wec := ErrorCode(e1)
		if wec != 0 {
			return 0, wec
		} else {
			return 0, errors.New("load bitmap by id failed")
		}
	}
}

func LoadBitmapByName(hInst HINSTANCE, Name string) (HBITMAP, error) {
	p, err := syscall.UTF16PtrFromString(Name)
	if err != nil {
		return 0, err
	}
	r1, _, e1 := syscall.Syscall(procLoadBitmap.Addr(), 2, uintptr(hInst), uintptr(unsafe.Pointer(p)), 0)
	if r1 != 0 {
		return HBITMAP(r1), nil
	} else {
		wec := ErrorCode(e1)
		if wec != 0 {
			return 0, wec
		} else {
			return 0, errors.New("load bitmap by name failed")
		}
	}
}

func LoadCursorById(hInst HINSTANCE, id uint16) (cursor HCURSOR, err error) {
	r1, _, e1 := syscall.Syscall(procLoadCursor.Addr(), 2, uintptr(hInst), uintptr(id), 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("load cursor by id failed")
		}
	} else {
		cursor = HCURSOR(r1)
	}
	return
}

func LoadCursorByName(hInst HINSTANCE, name string) (cursor HCURSOR, err error) {
	pName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall(procLoadCursor.Addr(), 2, uintptr(hInst), uintptr(unsafe.Pointer(pName)), 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("load cursor by name failed")
		}
	} else {
		cursor = HCURSOR(r1)
	}
	return
}

func LoadIconById(hInst HINSTANCE, id uint16) (icon HICON, err error) {
	r1, _, e1 := syscall.Syscall(procLoadIcon.Addr(), 2, uintptr(hInst), uintptr(id), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = errors.New("load icon by id failed")
		}
	} else {
		icon = HICON(r1)
	}
	return
}

func LoadIconByName(hInst HINSTANCE, name string) (icon HICON, err error) {
	pName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall(procLoadIcon.Addr(), 2, uintptr(hInst), uintptr(unsafe.Pointer(pName)), 0)
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("load icon by name failed")
		}
	} else {
		icon = HICON(r1)
	}
	return
}

func LoadImageById(hInst HINSTANCE, id uint16, Type uint32, cxDesired int32, cyDesired int32, fLoad uint32) (hImage HANDLE, err error) {
	r1, _, e1 := syscall.Syscall6(procLoadImage.Addr(), 6, uintptr(hInst), uintptr(id), uintptr(Type), uintptr(cxDesired), uintptr(cyDesired), uintptr(fLoad))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = errors.New("load image by id failed")
		}
	} else {
		hImage = HANDLE(r1)
	}
	return
}

func LoadImageByName(hInst HINSTANCE, name string, Type uint32, cxDesired int32, cyDesired int32, fLoad uint32) (hImage HANDLE, err error) {
	pName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	r1, _, e1 := syscall.Syscall6(procLoadImage.Addr(), 6, uintptr(hInst), uintptr(unsafe.Pointer(pName)), uintptr(Type), uintptr(cxDesired), uintptr(cyDesired), uintptr(fLoad))
	if r1 == 0 {
		wec := ErrorCode(e1)
		if wec != 0 {
			err = wec
		} else {
			err = errors.New("load image by name failed")
		}
	} else {
		hImage = HANDLE(r1)
	}
	return
}

func SendMessageTimeout(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr, fuFlags uint32, uTimeout uint32, lpdwResult uintptr) uintptr {
	ret, _, _ := procSendMessageTimeout.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
		uintptr(fuFlags),
		uintptr(uTimeout),
		lpdwResult)

	return ret
}

func GetClassNameW(hwnd HWND) string {
	buf := make([]uint16, 255)
	_, _, _ = procGetClassName.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(255))

	return syscall.UTF16ToString(buf)
}

func SetForegroundWindow(hwnd HWND) bool {
	ret, _, _ := procSetForegroundWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

func FindWindowExW(hwndParent HWND, hwndChildAfter HWND, className *uint16, windowName *uint16) HWND {
	ret, _, _ := procFindWindowExW.Call(
		uintptr(hwndParent),
		uintptr(hwndChildAfter),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)))

	return HWND(ret)
}

func FindWindowExS(hwndParent HWND, hwndChildAfter HWND, className *string, windowName *string) HWND {
	var class *uint16 = nil
	if className != nil {
		class = syscall.StringToUTF16Ptr(*className)
	}
	var window *uint16 = nil
	if windowName != nil {
		window = syscall.StringToUTF16Ptr(*windowName)
	}
	return FindWindowExW(hwndParent, hwndChildAfter, class, window)
}

func FindWindowW(className *uint16, windowName *uint16) HWND {
	ret, _, _ := procFindWindowW.Call(
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)))

	return HWND(ret)
}

func FindWindowS(className *string, windowName *string) HWND {
	var class *uint16 = nil
	if className != nil {
		class = syscall.StringToUTF16Ptr(*className)
	}
	var window *uint16 = nil
	if windowName != nil {
		window = syscall.StringToUTF16Ptr(*windowName)
	}
	return FindWindowW(class, window)
}

func EnumChildWindows(hWndParent HWND, lpEnumFunc WNDENUMPROC, lParam LPARAM) bool {
	ret, _, _ := procEnumChildWindows.Call(
		uintptr(hWndParent),
		uintptr(syscall.NewCallback(lpEnumFunc)),
		uintptr(lParam),
	)

	return ret != 0
}

func GetWindowTextW(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetForegroundWindow() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
		return
	}
	hwnd = syscall.Handle(r0)
	return
}

func RegisterClassEx(wndClassEx *WNDCLASSEX) ATOM {
	ret, _, _ := procRegisterClassEx.Call(uintptr(unsafe.Pointer(wndClassEx)))
	return ATOM(ret)
}

func LoadIcon(instance HINSTANCE, iconName *uint16) HICON {
	ret, _, _ := procLoadIcon.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(iconName)))

	return HICON(ret)

}

func LoadIconS(instance HINSTANCE, iconName *string) HICON {
	var icon *uint16 = nil
	if iconName != nil {
		icon = syscall.StringToUTF16Ptr(*iconName)
	}
	return LoadIcon(instance, icon)
}

func LoadCursor(instance HINSTANCE, cursorName *uint16) HCURSOR {
	ret, _, _ := procLoadCursor.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(cursorName)))

	return HCURSOR(ret)

}

func LoadCursorS(instance HINSTANCE, cursorName *string) HCURSOR {
	var cursor *uint16 = nil
	if cursorName != nil {
		cursor = syscall.StringToUTF16Ptr(*cursorName)
	}
	return LoadCursor(instance, cursor)
}

func ShowWindow(hwnd HWND, cmdshow int) bool {
	ret, _, _ := procShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdshow))

	return ret != 0

}

func UpdateWindow(hwnd HWND) bool {
	ret, _, _ := procUpdateWindow.Call(
		uintptr(hwnd))
	return ret != 0
}

func CreateWindowEx(exStyle uint, className *uint16, windowName *uint16,
	style uint, x int32, y int32, width int32, height int32, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {
	ret, _, _ := procCreateWindowEx.Call(
		uintptr(exStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(param))

	return HWND(ret)
}

func CreateWindowExS(exStyle uint, className *string, windowName *string,
	style uint, x int32, y int32, width int32, height int32, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {
	var class *uint16 = nil
	if className != nil {
		class = syscall.StringToUTF16Ptr(*className)
	}
	var window *uint16 = nil
	if windowName != nil {
		window = syscall.StringToUTF16Ptr(*windowName)
	}
	return CreateWindowEx(
		exStyle,
		class,
		window,
		style,
		x,
		y,
		width,
		height,
		parent,
		menu,
		instance,
		param,
	)
}

func AdjustWindowRectEx(rect *RECT, style uint, menu bool, exStyle uint) bool {
	ret, _, _ := procAdjustWindowRectEx.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)),
		uintptr(exStyle))

	return ret != 0
}

func AdjustWindowRect(rect *RECT, style uint, menu bool) bool {
	ret, _, _ := procAdjustWindowRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)))

	return ret != 0
}

func DestroyWindow(hwnd HWND) bool {
	ret, _, _ := procDestroyWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

func DefWindowProc(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := procDefWindowProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)
	return ret
}

func DefDlgProc(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := procDefDlgProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func SetWindowText(hwnd HWND, text string) {
	_, _, _ = procSetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	_, _, _ = procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func GetWindowRect(hwnd HWND) *RECT {
	var rect RECT
	_, _, _ = procGetWindowRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	return &rect
}

func MoveWindow(hwnd HWND, x int32, y int32, width int32, height int32, repaint bool) bool {
	ret, _, _ := procMoveWindow.Call(
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(BoolToBOOL(repaint)))

	return ret != 0

}

func ScreenToClient(hwnd HWND, x int32, y int32) (X, Y int, ok bool) {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procScreenToClient.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y), ret != 0
}

func CallWindowProc(preWndProc uintptr, hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := procCallWindowProc.Call(
		preWndProc,
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func SetWindowLong(hwnd HWND, index int, value uint32) uint32 {
	ret, _, _ := procSetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index),
		uintptr(value))

	return uint32(ret)
}

func SetWindowLongPtr(hwnd HWND, index int, value uintptr) uintptr {
	ret, _, _ := procSetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index),
		value)

	return ret
}

func GetWindowLong(hwnd HWND, index int) int32 {
	ret, _, _ := procGetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index))

	return int32(ret)
}

func GetWindowLongPtr(hwnd HWND, index int) uintptr {
	ret, _, _ := procGetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index))

	return ret
}

func EnableWindow(hwnd HWND, b bool) bool {
	ret, _, _ := procEnableWindow.Call(
		uintptr(hwnd),
		uintptr(BoolToBOOL(b)))
	return ret != 0
}

func IsWindowEnabled(hwnd HWND) bool {
	ret, _, _ := procIsWindowEnabled.Call(
		uintptr(hwnd))

	return ret != 0
}

func IsWindowVisible(hwnd HWND) bool {
	ret, _, _ := procIsWindowVisible.Call(
		uintptr(hwnd))

	return ret != 0
}

func SetFocus(hwnd HWND) HWND {
	ret, _, _ := procSetFocus.Call(
		uintptr(hwnd))

	return HWND(ret)
}

func InvalidateRect(hwnd HWND, rect *RECT, erase bool) bool {
	ret, _, _ := procInvalidateRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rect)),
		uintptr(BoolToBOOL(erase)))

	return ret != 0
}

func GetClientRect(hwnd HWND) *RECT {
	var rect RECT
	ret, _, _ := procGetClientRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	if ret == 0 {
		panic(fmt.Sprintf("GetClientRect(%d) failed", hwnd))
	}

	return &rect
}

func GetDC(hwnd HWND) HDC {
	ret, _, _ := procGetDC.Call(
		uintptr(hwnd))

	return HDC(ret)
}

func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}

func SetCapture(hwnd HWND) HWND {
	ret, _, _ := procSetCapture.Call(
		uintptr(hwnd))

	return HWND(ret)
}

func ReleaseCapture() bool {
	ret, _, _ := procReleaseCapture.Call()

	return ret != 0
}

func GetWindowThreadProcessId(hwnd HWND) (HANDLE, int) {
	var processId int
	ret, _, _ := procGetWindowThreadProcessId.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&processId)))

	return HANDLE(ret), processId
}

func MessageBox(hwnd HWND, text string, caption string, flags uint32) int32 {
	ret, _, _ := procMessageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(flags))

	return int32(ret)
}

func GetSystemMetrics(index int) int {
	ret, _, _ := procGetSystemMetrics.Call(
		uintptr(index))

	return int(ret)
}

func CopyRect(dst *RECT, src *RECT) bool {
	ret, _, _ := procCopyRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)))

	return ret != 0
}

func EqualRect(rect1 *RECT, rect2 *RECT) bool {
	ret, _, _ := procEqualRect.Call(
		uintptr(unsafe.Pointer(rect1)),
		uintptr(unsafe.Pointer(rect2)))

	return ret != 0
}

func InflateRect(rect *RECT, dx int32, dy int32) bool {
	ret, _, _ := procInflateRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}

func IntersectRect(dst *RECT, src1 *RECT, src2 *RECT) bool {
	ret, _, _ := procIntersectRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func IsRectEmpty(rect *RECT) bool {
	ret, _, _ := procIsRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}

func OffsetRect(rect *RECT, dx int32, dy int32) bool {
	ret, _, _ := procOffsetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}

func PtInRect(rect *RECT, x int32, y int32) bool {
	pt := POINT{X: x, Y: y}
	ret, _, _ := procPtInRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(&pt)))

	return ret != 0
}

func SetRect(rect *RECT, left int32, top int32, right int32, bottom int32) bool {
	ret, _, _ := procSetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))

	return ret != 0
}

func SetRectEmpty(rect *RECT) bool {
	ret, _, _ := procSetRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}

func SubtractRect(dst *RECT, src1 *RECT, src2 *RECT) bool {
	ret, _, _ := procSubtractRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func UnionRect(dst *RECT, src1 *RECT, src2 *RECT) bool {
	ret, _, _ := procUnionRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func CreateDialog(hInstance HINSTANCE, lpTemplate *uint16, hWndParent HWND, lpDialogProc uintptr) HWND {
	ret, _, _ := procCreateDialogParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplate)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return HWND(ret)
}

func DialogBox(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogProc uintptr) int {
	ret, _, _ := procDialogBoxParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplateName)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return int(ret)
}

func GetDlgItem(hDlg HWND, nIDDlgItem int) HWND {
	ret, _, _ := procGetDlgItem.Call(
		uintptr(unsafe.Pointer(hDlg)),
		uintptr(nIDDlgItem))

	return HWND(ret)
}

func DrawIcon(hDC HDC, x int32, y int32, hIcon HICON) bool {
	ret, _, _ := procDrawIcon.Call(
		uintptr(unsafe.Pointer(hDC)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(hIcon)))

	return ret != 0
}

func ClientToScreen(hwnd HWND, x int32, y int32) (int, int) {
	pt := POINT{X: x, Y: y}

	_, _, _ = procClientToScreen.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y)
}

func IsDialogMessage(hwnd HWND, msg *MSG) bool {
	ret, _, _ := procIsDialogMessage.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(msg)))

	return ret != 0
}

func IsWindow(hwnd HWND) bool {
	ret, _, _ := procIsWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

func EndDialog(hwnd HWND, nResult uintptr) bool {
	ret, _, _ := procEndDialog.Call(
		uintptr(hwnd),
		nResult)

	return ret != 0
}

func PeekMessage(hwnd HWND, wMsgFilterMin uint32, wMsgFilterMax uint32, wRemoveMsg uint32) (msg MSG, err error) {
	_, _, err = procPeekMessage.Call(
		uintptr(unsafe.Pointer(&msg)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg))

	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

func TranslateAccelerator(hwnd HWND, hAccTable HACCEL, lpMsg *MSG) bool {
	ret, _, _ := procTranslateAccelerator.Call(
		uintptr(hwnd),
		uintptr(hAccTable),
		uintptr(unsafe.Pointer(lpMsg)))

	return ret != 0
}

func SetWindowPos(hwnd, hWndInsertAfter HWND, x int32, y int32, cx int32, cy int32, uFlags uint) bool {
	ret, _, _ := procSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags))

	return ret != 0
}

func FillRect(hDC HDC, lprc *RECT, hbr HBRUSH) bool {
	ret, _, _ := procFillRect.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr))

	return ret != 0
}

func DrawText(hDC HDC, text string, uCount int, lpRect *RECT, uFormat uint) int {
	ret, _, _ := procDrawText.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(uCount),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(uFormat))

	return int(ret)
}

func AddClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procAddClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}

func RemoveClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procRemoveClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}

func OpenClipboard(hWndNewOwner HWND) bool {
	ret, _, _ := procOpenClipboard.Call(
		uintptr(hWndNewOwner))
	return ret != 0
}

func CloseClipboard() bool {
	ret, _, _ := procCloseClipboard.Call()
	return ret != 0
}

func EnumClipboardFormats(format uint) uint {
	ret, _, _ := procEnumClipboardFormats.Call(
		uintptr(format))
	return uint(ret)
}

func GetClipboardData(uFormat uint) HANDLE {
	ret, _, _ := procGetClipboardData.Call(
		uintptr(uFormat))
	return HANDLE(ret)
}

func SetClipboardData(uFormat uint, hMem HANDLE) HANDLE {
	ret, _, _ := procSetClipboardData.Call(
		uintptr(uFormat),
		uintptr(hMem))
	return HANDLE(ret)
}

func EmptyClipboard() bool {
	ret, _, _ := procEmptyClipboard.Call()
	return ret != 0
}

func GetClipboardFormatName(format uint) (string, bool) {
	cchMaxCount := 255
	buf := make([]uint16, cchMaxCount)
	ret, _, _ := procGetClipboardFormatName.Call(
		uintptr(format),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(cchMaxCount))

	if ret > 0 {
		return syscall.UTF16ToString(buf), true
	}

	return "Requested format does not exist or is predefined", false
}

func IsClipboardFormatAvailable(format uint) bool {
	ret, _, _ := procIsClipboardFormatAvailable.Call(uintptr(format))
	return ret != 0
}

func GetKeyboardState(lpKeyState *[]byte) bool {
	ret, _, _ := procGetKeyboardState.Call(
		uintptr(unsafe.Pointer(&(*lpKeyState)[0])))
	return ret != 0
}

func MapVirtualKeyEx(uCode uint, uMapType uint, dwhkl HKL) uint {
	ret, _, _ := procMapVirtualKey.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl))
	return uint(ret)
}

func GetAsyncKeyState(vKey int) (keyState uint16) {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return uint16(ret)
}

func GetKeyState(vKey int) (keyState uint16) {
	ret, _, _ := procGetKeyState.Call(uintptr(vKey))
	return uint16(ret)
}

func ToAscii(uVirtKey uint, uScanCode uint, lpKeyState *byte, lpChar *uint16, uFlags uint) int {
	ret, _, _ := procToAscii.Call(
		uintptr(uVirtKey),
		uintptr(uScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(lpChar)),
		uintptr(uFlags))
	return int(ret)
}

func SwapMouseButton(fSwap bool) bool {
	ret, _, _ := procSwapMouseButton.Call(
		uintptr(BoolToBOOL(fSwap)))
	return ret != 0
}

func GetCursorPos() (x int, y int, ok bool) {
	pt := POINT{}
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y), ret != 0
}

func SetCursorPos(x int, y int) bool {
	ret, _, _ := procSetCursorPos.Call(
		uintptr(x),
		uintptr(y),
	)
	return ret != 0
}

func SetCursor(cursor HCURSOR) HCURSOR {
	ret, _, _ := procSetCursor.Call(
		uintptr(cursor),
	)
	return HCURSOR(ret)
}

func CreateIcon(instance HINSTANCE, nWidth int, nHeight int, cPlanes byte, cBitsPerPixel byte, ANDbits *byte, XORbits *byte) HICON {
	ret, _, _ := procCreateIcon.Call(
		uintptr(instance),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(cPlanes),
		uintptr(cBitsPerPixel),
		uintptr(unsafe.Pointer(ANDbits)),
		uintptr(unsafe.Pointer(XORbits)),
	)
	return HICON(ret)
}

func DestroyIcon(icon HICON) bool {
	ret, _, _ := procDestroyIcon.Call(
		uintptr(icon),
	)
	return ret != 0
}

func MonitorFromPoint(x int, y int, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromPoint.Call(
		uintptr(x),
		uintptr(y),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func MonitorFromRect(rc *RECT, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromRect.Call(
		uintptr(unsafe.Pointer(rc)),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromWindow.Call(
		uintptr(hwnd),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func GetMonitorInfo(hMonitor HMONITOR, lmpi *MONITORINFO) bool {
	ret, _, _ := procGetMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lmpi)),
	)
	return ret != 0
}

func EnumDisplayMonitors(hdc HDC, clip *RECT, fnEnum uintptr, dwData uintptr) bool {
	ret, _, _ := procEnumDisplayMonitors.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(clip)),
		fnEnum,
		dwData,
	)
	return ret != 0
}

func EnumDisplaySettingsEx(szDeviceName *uint16, iModeNum uint32, devMode *DEVMODE, dwFlags uint32) bool {
	ret, _, _ := procEnumDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(dwFlags),
	)
	return ret != 0
}

func ChangeDisplaySettingsEx(szDeviceName *uint16, devMode *DEVMODE, hwnd HWND, dwFlags uint32, lParam uintptr) int32 {
	ret, _, _ := procChangeDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(hwnd),
		uintptr(dwFlags),
		lParam,
	)
	return int32(ret)
}

//Synthesizes keystrokes, mouse motions, and button clicks.
func SendInput(inputs []INPUT) (err error) {
	var validInputs []INPUT

	for _, oneInput := range inputs {
		input := INPUT{Type: uint32(DWORD(oneInput.Type))}

		switch oneInput.Type {
		case INPUT_MOUSE:
			(*MouseInput)(unsafe.Pointer(&input)).mi = oneInput.Mi
		case INPUT_KEYBOARD:
			(*KbdInput)(unsafe.Pointer(&input)).ki = oneInput.Ki
		case INPUT_HARDWARE:
			(*HardwareInput)(unsafe.Pointer(&input)).hi = oneInput.Hi
		default:
			err = errors.New("Unknown input type passed: " + fmt.Sprintf("%d", oneInput.Type))
			return
		}

		validInputs = append(validInputs, input)
	}
	if validInputs != nil {
		_, _, err = procSendInput.Call(
			uintptr(len(validInputs)),
			uintptr(unsafe.Pointer(&validInputs[0])),
			unsafe.Sizeof(INPUT{}),
		)
	}
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

//Simplifies SendInput for Keyboard related keys. Supports alphanumeric
func SendInputString(input string) (err error) {
	var inputs []INPUT
	b := make([]byte, 3)

	for _, it := range input {

		utf8.EncodeRune(b, it)
		vk := binary.LittleEndian.Uint16(b)
		fmt.Println(vk)

		input := INPUT{
			Type: INPUT_KEYBOARD,
			Ki: KEYBDINPUT{
				WVk:     vk,
				DwFlags: 0x0002 | 0x0008,
				Time:    200,
			},
		}
		inputs = append(inputs, input)
	}
	err = SendInput(inputs)
	return
}

func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	ret, _, _ := procSetWindowsHookEx.Call(
		uintptr(idHook),
		syscall.NewCallback(lpfn),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return HHOOK(ret)
}

func SetWinEventHook(eventMin DWORD, eventMax DWORD, hmodWinEventProc HMODULE, pfnWinEventProc HOOKPROC, idProcess DWORD, idThread DWORD, dwFlags DWORD) HHOOK {

	ret, _, _ := procSetWinEventHook.Call(
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		uintptr(syscall.NewCallback(pfnWinEventProc)),
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
	)
	return HHOOK(ret)
}

func UnhookWinEvent(hWinEventHook HHOOK) bool {
	ret, _, _ := procUnhookWinEvent.Call(
		uintptr(hWinEventHook))
	return ret != 0
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret, _, _ := procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)
	return ret != 0
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
}

//Defines a system-wide hotkey.
func RegisterHotKey(hwnd HWND, id int, fsModifiers uint, vkey uint) (err error) {
	_, _, err = procRegisterHotKey.Call(
		uintptr(hwnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vkey),
	)
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

//Defines a system-wide hotkey.
func UnregisterHotKey(hwnd HWND, id int) (err error) {
	_, _, err = procUnregisterHotKey.Call(
		uintptr(hwnd),
		uintptr(id),
	)
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

// Translates a character to the corresponding virtual-key code and shift state for the current keyboard.
func VkKeyScanW(char uint16) int16 {
	ret, _, _ := procVkKeyScanW.Call(
		uintptr(char),
	)
	return int16(ret)
}

// Translates a character to the corresponding virtual-key code and shift state.
func VkKeyScanExW(char uint16, hkl HKL) int16 {
	ret, _, _ := procVkKeyScanExW.Call(
		uintptr(char),
		uintptr(hkl),
	)
	return int16(ret)
}

// Sets the process-default DPI awareness to system-DPI awareness.
func SetProcessDPIAware() bool {
	ret, _, _ := setProcessDPIAware.Call()
	return ret != 0
}
