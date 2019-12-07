package windows

import (
	"syscall"
	"unsafe"
)

var (
	gdi                           = syscall.NewLazyDLL("gdi32.dll")
	procAbortDoc                  = gdi.NewProc("AbortDoc")
	procBitBlt                    = gdi.NewProc("BitBlt")
	procChoosePixelFormat         = gdi.NewProc("ChoosePixelFormat")
	procCloseEnhMetaFile          = gdi.NewProc("CloseEnhMetaFile")
	procCopyEnhMetaFile           = gdi.NewProc("CopyEnhMetaFileW")
	procCreateBrushIndirect       = gdi.NewProc("CreateBrushIndirect")
	procCreateCompatibleBitmap    = gdi.NewProc("CreateCompatibleBitmap")
	procCreateCompatibleDC        = gdi.NewProc("CreateCompatibleDC")
	procCreateDC                  = gdi.NewProc("CreateDCW")
	procCreateDIBSection          = gdi.NewProc("CreateDIBSection")
	procCreateEnhMetaFile         = gdi.NewProc("CreateEnhMetaFileW")
	procCreateFontIndirect        = gdi.NewProc("CreateFontIndirectW")
	procCreateIC                  = gdi.NewProc("CreateICW")
	procDeleteDC                  = gdi.NewProc("DeleteDC")
	procDeleteEnhMetaFile         = gdi.NewProc("DeleteEnhMetaFile")
	procDeleteObject              = gdi.NewProc("DeleteObject")
	procDescribePixelFormat       = gdi.NewProc("DescribePixelFormat")
	procEllipse                   = gdi.NewProc("Ellipse")
	procEndDoc                    = gdi.NewProc("EndDoc")
	procEndPage                   = gdi.NewProc("EndPage")
	procExtCreatePen              = gdi.NewProc("ExtCreatePen")
	procGetCurrentObject          = gdi.NewProc("GetCurrentObject")
	procGetDeviceCaps             = gdi.NewProc("GetDeviceCaps")
	procGetEnhMetaFile            = gdi.NewProc("GetEnhMetaFileW")
	procGetEnhMetaFileHeader      = gdi.NewProc("GetEnhMetaFileHeader")
	procGetEnhMetaFilePixelFormat = gdi.NewProc("GetEnhMetaFilePixelFormat")
	procGetObject                 = gdi.NewProc("GetObjectW")
	procGetPixelFormat            = gdi.NewProc("GetPixelFormat")
	procGetStockObject            = gdi.NewProc("GetStockObject")
	procGetTextExtentExPoint      = gdi.NewProc("GetTextExtentExPointW")
	procGetTextExtentPoint32      = gdi.NewProc("GetTextExtentPoint32W")
	procGetTextMetrics            = gdi.NewProc("GetTextMetricsW")
	procLineTo                    = gdi.NewProc("LineTo")
	procMoveToEx                  = gdi.NewProc("MoveToEx")
	procPatBlt                    = gdi.NewProc("PatBlt")
	procPlayEnhMetaFile           = gdi.NewProc("PlayEnhMetaFile")
	procRectangle                 = gdi.NewProc("Rectangle")
	procResetDC                   = gdi.NewProc("ResetDCW")
	procSelectObject              = gdi.NewProc("SelectObject")
	procSetBkColor                = gdi.NewProc("SetBkColor")
	procSetBkMode                 = gdi.NewProc("SetBkMode")
	procSetBrushOrgEx             = gdi.NewProc("SetBrushOrgEx")
	procSetDIBitsToDevice         = gdi.NewProc("SetDIBitsToDevice")
	procStretchDIBits             = gdi.NewProc("StretchDIBits")
	procSetPixelFormat            = gdi.NewProc("SetPixelFormat")
	procSetStretchBltMode         = gdi.NewProc("SetStretchBltMode")
	procSetTextColor              = gdi.NewProc("SetTextColor")
	procStartDoc                  = gdi.NewProc("StartDocW")
	procStartPage                 = gdi.NewProc("StartPage")
	procStretchBlt                = gdi.NewProc("StretchBlt")
	procSwapBuffers               = gdi.NewProc("SwapBuffers")
)

func BeginPaint(hwnd HWND, paint *PAINTSTRUCT) HDC {
	ret, _, _ := procBeginPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
	return HDC(ret)
}

func EndPaint(hwnd HWND, paint *PAINTSTRUCT) {
	_, _, _ = procEndPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
}

func CreateCompatibleBitmap(hdc HDC, width uint, height uint) HBITMAP {
	ret, _, _ := procCreateCompatibleBitmap.Call(
		uintptr(hdc),
		uintptr(width),
		uintptr(height))

	return HBITMAP(ret)
}

func GetCurrentObject(hdc HDC, uObjectType uint32) HGDIOBJ {
	ret, _, _ := procGetCurrentObject.Call(
		uintptr(hdc),
		uintptr(uObjectType))

	return HGDIOBJ(ret)
}

func GetDeviceCaps(hdc HDC, index int) int {
	ret, _, _ := procGetDeviceCaps.Call(
		uintptr(hdc),
		uintptr(index))

	return int(ret)
}

func DeleteObject(hObject HGDIOBJ) bool {
	ret, _, _ := procDeleteObject.Call(
		uintptr(hObject))
	return ret != 0
}

func CreateFontIndirect(logFont *LOGFONT) HFONT {
	ret, _, _ := procCreateFontIndirect.Call(
		uintptr(unsafe.Pointer(logFont)))

	return HFONT(ret)
}

func AbortDoc(hdc HDC) int {
	ret, _, _ := procAbortDoc.Call(
		uintptr(hdc))

	return int(ret)
}

func BitBlt(hdcDest HDC, nXDest int32, nYDest int32, nWidth int32, nHeight int32, hdcSrc HDC, nXSrc int32, nYSrc int32, dwRop uint32) {
	ret, _, _ := procBitBlt.Call(
		uintptr(hdcDest),
		uintptr(nXDest),
		uintptr(nYDest),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hdcSrc),
		uintptr(nXSrc),
		uintptr(nYSrc),
		uintptr(dwRop))

	if ret == 0 {
		panic("BitBlt failed")
	}
}

func PatBlt(hdc HDC, nXLeft int32, nYLeft int32, nWidth int32, nHeight int32, dwRop uint32) {
	ret, _, _ := procPatBlt.Call(
		uintptr(hdc),
		uintptr(nXLeft),
		uintptr(nYLeft),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(dwRop))

	if ret == 0 {
		panic("PatBlt failed")
	}
}

func CloseEnhMetaFile(hdc HDC) HENHMETAFILE {
	ret, _, _ := procCloseEnhMetaFile.Call(
		uintptr(hdc))

	return HENHMETAFILE(ret)
}

func CopyEnhMetaFile(hemfSrc HENHMETAFILE, lpszFile *uint16) HENHMETAFILE {
	ret, _, _ := procCopyEnhMetaFile.Call(
		uintptr(hemfSrc),
		uintptr(unsafe.Pointer(lpszFile)))

	return HENHMETAFILE(ret)
}

func CreateBrushIndirect(lplb *LOGBRUSH) HBRUSH {
	ret, _, _ := procCreateBrushIndirect.Call(
		uintptr(unsafe.Pointer(lplb)))

	return HBRUSH(ret)
}

func CreateCompatibleDC(hdc HDC) HDC {
	ret, _, _ := procCreateCompatibleDC.Call(
		uintptr(hdc))

	if ret == 0 {
		panic("Create compatible DC failed")
	}

	return HDC(ret)
}

func CreateDC(lpszDriver *uint16, lpszDevice *uint16, lpszOutput *uint16, lpInitData *DEVMODE) HDC {
	ret, _, _ := procCreateDC.Call(
		uintptr(unsafe.Pointer(lpszDriver)),
		uintptr(unsafe.Pointer(lpszDevice)),
		uintptr(unsafe.Pointer(lpszOutput)),
		uintptr(unsafe.Pointer(lpInitData)))

	return HDC(ret)
}

func CreateDIBSection(hdc HDC, pbmi *BITMAPINFO, iUsage uint, ppvBits *unsafe.Pointer, hSection HANDLE, dwOffset uint) HBITMAP {
	ret, _, _ := procCreateDIBSection.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(pbmi)),
		uintptr(iUsage),
		uintptr(unsafe.Pointer(ppvBits)),
		uintptr(hSection),
		uintptr(dwOffset))

	return HBITMAP(ret)
}

func CreateEnhMetaFile(hdcRef HDC, lpFilename *uint16, lpRect *RECT, lpDescription *uint16) HDC {
	ret, _, _ := procCreateEnhMetaFile.Call(
		uintptr(hdcRef),
		uintptr(unsafe.Pointer(lpFilename)),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(unsafe.Pointer(lpDescription)))

	return HDC(ret)
}

func CreateIC(lpszDriver *uint16, lpszDevice *uint16, lpszOutput *uint16, lpdvmInit *DEVMODE) HDC {
	ret, _, _ := procCreateIC.Call(
		uintptr(unsafe.Pointer(lpszDriver)),
		uintptr(unsafe.Pointer(lpszDevice)),
		uintptr(unsafe.Pointer(lpszOutput)),
		uintptr(unsafe.Pointer(lpdvmInit)))

	return HDC(ret)
}

func DeleteDC(hdc HDC) bool {
	ret, _, _ := procDeleteDC.Call(
		uintptr(hdc))

	return ret != 0
}

func DeleteEnhMetaFile(hemf HENHMETAFILE) bool {
	ret, _, _ := procDeleteEnhMetaFile.Call(
		uintptr(hemf))

	return ret != 0
}

func Ellipse(hdc HDC, nLeftRect int32, nTopRect int32, nRightRect int32, nBottomRect int32) bool {
	ret, _, _ := procEllipse.Call(
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect))

	return ret != 0
}

func EndDoc(hdc HDC) int {
	ret, _, _ := procEndDoc.Call(
		uintptr(hdc))

	return int(ret)
}

func EndPage(hdc HDC) int {
	ret, _, _ := procEndPage.Call(
		uintptr(hdc))

	return int(ret)
}

func ExtCreatePen(dwPenStyle uint32, dwWidth uint32, lplb *LOGBRUSH, dwStyleCount uint32, lpStyle *uint) HPEN {
	ret, _, _ := procExtCreatePen.Call(
		uintptr(dwPenStyle),
		uintptr(dwWidth),
		uintptr(unsafe.Pointer(lplb)),
		uintptr(dwStyleCount),
		uintptr(unsafe.Pointer(lpStyle)))

	return HPEN(ret)
}

func GetEnhMetaFile(lpszMetaFile *uint16) HENHMETAFILE {
	ret, _, _ := procGetEnhMetaFile.Call(
		uintptr(unsafe.Pointer(lpszMetaFile)))

	return HENHMETAFILE(ret)
}

func GetEnhMetaFileHeader(hemf HENHMETAFILE, cbBuffer uint, lpemh *ENHMETAHEADER) uint {
	ret, _, _ := procGetEnhMetaFileHeader.Call(
		uintptr(hemf),
		uintptr(cbBuffer),
		uintptr(unsafe.Pointer(lpemh)))

	return uint(ret)
}

func GetObject(hGdiObj HGDIOBJ, cbBuffer uintptr, lpvObject unsafe.Pointer) int32 {
	ret, _, _ := procGetObject.Call(
		uintptr(hGdiObj),
		uintptr(cbBuffer),
		uintptr(lpvObject))

	return int32(ret)
}

func GetStockObject(fnObject int) HGDIOBJ {
	ret, _, _ := procGetStockObject.Call(
		uintptr(fnObject))

	return HGDIOBJ(ret)
}

func GetTextExtentExPoint(hdc HDC, lpszStr *uint16, cchString int32, nMaxExtent int32, lpnFit *int, alpDx *int, lpSize *SIZE) bool {
	ret, _, _ := procGetTextExtentExPoint.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(cchString),
		uintptr(nMaxExtent),
		uintptr(unsafe.Pointer(lpnFit)),
		uintptr(unsafe.Pointer(alpDx)),
		uintptr(unsafe.Pointer(lpSize)))

	return ret != 0
}

func GetTextExtentPoint32(hdc HDC, lpString *uint16, c int32, lpSize *SIZE) bool {
	ret, _, _ := procGetTextExtentPoint32.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(c),
		uintptr(unsafe.Pointer(lpSize)))

	return ret != 0
}

func GetTextMetrics(hdc HDC, lptm *TEXTMETRIC) bool {
	ret, _, _ := procGetTextMetrics.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(lptm)))

	return ret != 0
}

func LineTo(hdc HDC, nXEnd int32, nYEnd int32) bool {
	ret, _, _ := procLineTo.Call(
		uintptr(hdc),
		uintptr(nXEnd),
		uintptr(nYEnd))

	return ret != 0
}

func MoveToEx(hdc HDC, x int32, y int32, lpPoint *POINT) bool {
	ret, _, _ := procMoveToEx.Call(
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lpPoint)))

	return ret != 0
}

func PlayEnhMetaFile(hdc HDC, hemf HENHMETAFILE, lpRect *RECT) bool {
	ret, _, _ := procPlayEnhMetaFile.Call(
		uintptr(hdc),
		uintptr(hemf),
		uintptr(unsafe.Pointer(lpRect)))

	return ret != 0
}

func Rectangle(hdc HDC, nLeftRect int32, nTopRect int32, nRightRect int32, nBottomRect int32) bool {
	ret, _, _ := procRectangle.Call(
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect))

	return ret != 0
}

func ResetDC(hdc HDC, lpInitData *DEVMODE) HDC {
	ret, _, _ := procResetDC.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpInitData)))

	return HDC(ret)
}

func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
	ret, _, _ := procSelectObject.Call(
		uintptr(hdc),
		uintptr(hgdiobj))

	if ret == 0 {
		panic("SelectObject failed")
	}

	return HGDIOBJ(ret)
}

func SetBkMode(hdc HDC, iBkMode int32) int {
	ret, _, _ := procSetBkMode.Call(
		uintptr(hdc),
		uintptr(iBkMode))

	if ret == 0 {
		panic("SetBkMode failed")
	}

	return int(ret)
}

func SetBrushOrgEx(hdc HDC, nXOrg int32, nYOrg int32, lppt *POINT) bool {
	ret, _, _ := procSetBrushOrgEx.Call(
		uintptr(hdc),
		uintptr(nXOrg),
		uintptr(nYOrg),
		uintptr(unsafe.Pointer(lppt)))

	return ret != 0
}

func SetStretchBltMode(hdc HDC, iStretchMode int32) int32 {
	ret, _, _ := procSetStretchBltMode.Call(
		uintptr(hdc),
		uintptr(iStretchMode))

	return int32(ret)
}

func SetTextColor(hdc HDC, crColor COLORREF) COLORREF {
	ret, _, _ := procSetTextColor.Call(
		uintptr(hdc),
		uintptr(crColor))

	if ret == CLR_INVALID {
		panic("SetTextColor failed")
	}

	return COLORREF(ret)
}

func SetBkColor(hdc HDC, crColor COLORREF) COLORREF {
	ret, _, _ := procSetBkColor.Call(
		uintptr(hdc),
		uintptr(crColor))

	if ret == CLR_INVALID {
		panic("SetBkColor failed")
	}

	return COLORREF(ret)
}

func StartDoc(hdc HDC, lpdi *DOCINFO) int32 {
	ret, _, _ := procStartDoc.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpdi)))

	return int32(ret)
}

func StartPage(hdc HDC) int32 {
	ret, _, _ := procStartPage.Call(
		uintptr(hdc))

	return int32(ret)
}

func StretchBlt(hdcDest HDC, nXOriginDest int32, nYOriginDest int32, nWidthDest int32, nHeightDest int32, hdcSrc HDC, nXOriginSrc int32, nYOriginSrc int32, nWidthSrc int32, nHeightSrc int32, dwRop uint32) {
	ret, _, _ := procStretchBlt.Call(
		uintptr(hdcDest),
		uintptr(nXOriginDest),
		uintptr(nYOriginDest),
		uintptr(nWidthDest),
		uintptr(nHeightDest),
		uintptr(hdcSrc),
		uintptr(nXOriginSrc),
		uintptr(nYOriginSrc),
		uintptr(nWidthSrc),
		uintptr(nHeightSrc),
		uintptr(dwRop))

	if ret == 0 {
		panic("StretchBlt failed")
	}
}

func StretchDIBits(hdc HDC, xDest int32, yDest int32, destWidth int32, destHeight int32, xSrc int32, ySrc int32, nSrcWidth int32, nSrcHeight int32, lpBits unsafe.Pointer, lpBitsInfo *BITMAPINFO, iUsage uint32, rop uint32) int32 {
	r1, _, _ := syscall.Syscall15(procStretchDIBits.Addr(), 13,
		uintptr(hdc),
		uintptr(xDest),
		uintptr(yDest),
		uintptr(destWidth),
		uintptr(destHeight),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(nSrcWidth),
		uintptr(nSrcHeight),
		uintptr(lpBits),
		uintptr(unsafe.Pointer(lpBitsInfo)),
		uintptr(iUsage),
		uintptr(rop),
		0,
		0)
	return int32(r1)
}

func SetDIBitsToDevice(hdc HDC, xDest int32, yDest int32, dwWidth int32, dwHeight int32, xSrc int32, ySrc int32, uStartScan uint32, cScanLines uint32, lpvBits []byte, lpbmi *BITMAPINFO, fuColorUse uint32) int {
	ret, _, _ := procSetDIBitsToDevice.Call(
		uintptr(hdc),
		uintptr(xDest),
		uintptr(yDest),
		uintptr(dwWidth),
		uintptr(dwHeight),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(uStartScan),
		uintptr(cScanLines),
		uintptr(unsafe.Pointer(&lpvBits[0])),
		uintptr(unsafe.Pointer(lpbmi)),
		uintptr(fuColorUse))

	return int(ret)
}

func ChoosePixelFormat(hdc HDC, pfd *PIXELFORMATDESCRIPTOR) int {
	ret, _, _ := procChoosePixelFormat.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(pfd)),
	)
	return int(ret)
}

func DescribePixelFormat(hdc HDC, iPixelFormat int32, nBytes uint, pfd *PIXELFORMATDESCRIPTOR) int {
	ret, _, _ := procDescribePixelFormat.Call(
		uintptr(hdc),
		uintptr(iPixelFormat),
		uintptr(nBytes),
		uintptr(unsafe.Pointer(pfd)),
	)
	return int(ret)
}

func GetEnhMetaFilePixelFormat(hemf HENHMETAFILE, cbBuffer uint32, pfd *PIXELFORMATDESCRIPTOR) uint {
	ret, _, _ := procGetEnhMetaFilePixelFormat.Call(
		uintptr(hemf),
		uintptr(cbBuffer),
		uintptr(unsafe.Pointer(pfd)),
	)
	return uint(ret)
}

func GetPixelFormat(hdc HDC) int {
	ret, _, _ := procGetPixelFormat.Call(
		uintptr(hdc),
	)
	return int(ret)
}

func SetPixelFormat(hdc HDC, iPixelFormat int, pfd *PIXELFORMATDESCRIPTOR) bool {
	ret, _, _ := procSetPixelFormat.Call(
		uintptr(hdc),
		uintptr(iPixelFormat),
		uintptr(unsafe.Pointer(pfd)),
	)
	return ret == TRUE
}

func SwapBuffers(hdc HDC) bool {
	ret, _, _ := procSwapBuffers.Call(uintptr(hdc))
	return ret == TRUE
}
