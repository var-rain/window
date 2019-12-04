package main

import (
	"fmt"
	"github.com/var-rain/window/win"
)

// ---------------------------------------------------------
// Remove console window please use -ldflags="-H windowsgui"
// example:
// go build -ldflags="-H windowsgui" source.go
// ---------------------------------------------------------

const WINDOW_CLASS string = "MAIN_WINDOW_CLASS"

func main() {
	inst := win.GetModuleHandle("")

	var err error

	// register window class.
	_, err = WindowRegisterClass(inst)
	if err != nil {
		fmt.Println("window register class failed")
		return
	}

	// create registered window.
	wnd, err := win.CreateWindow(WINDOW_CLASS, "Window By Golang",
		win.WS_OVERLAPPEDWINDOW, 0,
		win.CW_USEDEFAULT, win.CW_USEDEFAULT, win.CW_USEDEFAULT, win.CW_USEDEFAULT,
		0, 0, inst, 0)
	if err != nil {
		fmt.Println("window create failed")
		return
	}
	win.ShowWindow(wnd, win.SW_SHOW)
	win.UpdateWindow(wnd)

	// main message loop.
	var msg win.MSG
	msg.Message = win.WM_QUIT + 1

	for win.GetMessage(&msg, 0, 0, 0) > 0 {
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}

	fmt.Println("application finished with exit code", msg.WParam)
}

func WndProc(hWnd win.HWND, message uint32, wParam uintptr, lParam uintptr) uintptr {
	switch message {
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
	case win.WM_COMMAND:
		OnCommand(hWnd, wParam, lParam)
	default:
		return win.DefWindowProc(hWnd, message, wParam, lParam)
	}
	return 0
}

func OnCommand(hWnd win.HWND, wParam uintptr, lParam uintptr) {
	win.DefWindowProc(hWnd, win.WM_COMMAND, wParam, lParam)
}

func WindowRegisterClass(hInstance win.HINSTANCE) (atom uint16, err error) {
	var wc win.WNDCLASS
	wc.Style = win.CS_HREDRAW | win.CS_VREDRAW
	wc.PfnWndProc = WndProc
	wc.CbClsExtra = 0
	wc.CbWndExtra = 0
	wc.HInstance = hInstance
	wc.HIcon = 0
	wc.HCursor, err = win.LoadCursorById(0, win.IDC_ARROW)
	if err != nil {
		return
	}
	wc.Menu = ""
	wc.HbrBackground = win.COLOR_WINDOW + 1
	wc.PszClassName = WINDOW_CLASS
	wc.HIconSmall = 0

	return win.RegisterClass(&wc)
}
