package windows

import (
	"syscall"
	"unsafe"
)

var (
	ogl                       = syscall.NewLazyDLL("opengl32.dll")
	procWglCreateContext      = ogl.NewProc("wglCreateContext")
	procWglCreateLayerContext = ogl.NewProc("wglCreateLayerContext")
	procWglDeleteContext      = ogl.NewProc("wglDeleteContext")
	procWglGetProcAddress     = ogl.NewProc("wglGetProcAddress")
	procWglMakeCurrent        = ogl.NewProc("wglMakeCurrent")
	procWglShareLists         = ogl.NewProc("wglShareLists")
)

func WglCreateContext(hdc HDC) HGLRC {
	ret, _, _ := procWglCreateContext.Call(
		uintptr(hdc),
	)

	return HGLRC(ret)
}

func WglCreateLayerContext(hdc HDC, iLayerPlane int) HGLRC {
	ret, _, _ := procWglCreateLayerContext.Call(
		uintptr(hdc),
		uintptr(iLayerPlane),
	)

	return HGLRC(ret)
}

func WglDeleteContext(hglrc HGLRC) bool {
	ret, _, _ := procWglDeleteContext.Call(
		uintptr(hglrc),
	)

	return ret == TRUE
}

func WglGetProcAddress(szProc string) uintptr {
	ret, _, _ := procWglGetProcAddress.Call(
		uintptr(unsafe.Pointer(syscall.StringBytePtr(szProc))),
	)

	return ret
}

func WglMakeCurrent(hdc HDC, hglrc HGLRC) bool {
	ret, _, _ := procWglMakeCurrent.Call(
		uintptr(hdc),
		uintptr(hglrc),
	)

	return ret == TRUE
}

func WglShareLists(hglrc1 HGLRC, hglrc2 HGLRC) bool {
	ret, _, _ := procWglShareLists.Call(
		uintptr(hglrc1),
		uintptr(hglrc2),
	)

	return ret == TRUE
}
