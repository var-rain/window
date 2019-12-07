package main

import (
	"fmt"
	"github.com/var-rain/window/windows"
)

// ---------------------------------------------------------
// Remove console window please use -ldflags="-H windowsgui"
// example:
// go build -ldflags="-H windowsgui" source.go
// ---------------------------------------------------------

const WINDOW_CLASS string = "MAIN_WINDOW_CLASS"

func main() {
	inst := windows.GetModuleHandle("")

	var err error

	// register window class.
	_, err = WindowRegisterClass(inst)
	if err != nil {
		fmt.Println("window register class failed")
		return
	}

	// create registered window.
	wnd, err := windows.CreateWindow(WINDOW_CLASS, "Window By Golang",
		windows.WS_OVERLAPPEDWINDOW, 0,
		windows.CW_USEDEFAULT, windows.CW_USEDEFAULT, windows.CW_USEDEFAULT, windows.CW_USEDEFAULT,
		0, 0, inst, 0)
	if err != nil {
		fmt.Println("window create failed")
		return
	}
	windows.ShowWindow(wnd, windows.SW_SHOW)
	windows.UpdateWindow(wnd)

	// main message loop.
	var msg windows.MSG
	msg.Message = windows.WM_QUIT + 1

	for windows.GetMessage(&msg, 0, 0, 0) > 0 {
		windows.TranslateMessage(&msg)
		windows.DispatchMessage(&msg)
	}

	fmt.Println("application finished with exit code", msg.WParam)
}

func WndProc(hWnd windows.HWND, message uint32, wParam uintptr, lParam uintptr) uintptr {
	switch message {
	case windows.WM_DESTROY:
		windows.PostQuitMessage(0)
	case windows.WM_COMMAND:
		OnCommand(hWnd, wParam, lParam)
	default:
		return windows.DefWindowProc(hWnd, message, wParam, lParam)
	}
	return 0
}

func OnCommand(hWnd windows.HWND, wParam uintptr, lParam uintptr) {
	windows.DefWindowProc(hWnd, windows.WM_COMMAND, wParam, lParam)
}

func WindowRegisterClass(hInstance windows.HINSTANCE) (atom uint16, err error) {
	var wc windows.WNDCLASS
	wc.Style = windows.CS_HREDRAW | windows.CS_VREDRAW
	wc.PfnWndProc = WndProc
	wc.CbClsExtra = 0
	wc.CbWndExtra = 0
	wc.HInstance = hInstance
	wc.HIcon = 0
	wc.HCursor, err = windows.LoadCursorById(0, windows.IDC_ARROW)
	if err != nil {
		return
	}
	wc.Menu = ""
	wc.HbrBackground = windows.COLOR_WINDOW + 1
	wc.PszClassName = WINDOW_CLASS
	wc.HIconSmall = 0

	return windows.RegisterClass(&wc)
}
