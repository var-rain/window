package win

import (
	"errors"
	"unsafe"
)

// Windows Data Types
// ATOM                  WORD
// BOOL                  int32
// BOOLEAN               byte
// BYTE                  byte
// CCHAR                 int8
// CHAR                  int8
// COLORREF              DWORD
// DWORD                 uint32
// DWORDLONG             ULONGLONG
// DWORD_PTR             ULONG_PTR
// DWORD32               uint32
// DWORD64               uint64
// FLOAT                 float32
// HACCEL                HANDLE
// HALF_PTR              struct{}
// HANDLE                PVOID
// HBITMAP               HANDLE
// HBRUSH                HANDLE
// HCOLORSPACE           HANDLE
// HCONV                 HANDLE
// HCONVLIST             HANDLE
// HCURSOR               HANDLE
// HDC                   HANDLE
// HDDEDATA              HANDLE
// HDESK                 HANDLE
// HDROP                 HANDLE
// HDWP                  HANDLE
// HENHMETAFILE          HANDLE
// HFILE                 HANDLE
// HFONT                 HANDLE
// HGDIOBJ               HANDLE
// HGLOBAL               HANDLE
// HHOOK                 HANDLE
// HICON                 HANDLE
// HINSTANCE             HANDLE
// HKEY                  HANDLE
// HKL                   HANDLE
// HLOCAL                HANDLE
// HMENU                 HANDLE
// HMETAFILE             HANDLE
// HMODULE               HANDLE
// HPALETTE              HANDLE
// HPEN                  HANDLE
// HRESULT               int32
// HRGN                  HANDLE
// HSZ                   HANDLE
// HWINSTA               HANDLE
// HWND                  HANDLE
// INT                   int32
// INT_PTR               uintptr
// INT8                  int8
// INT16                 int16
// INT32                 int32
// INT64                 int64
// LANGID                WORD
// LCID                  DWORD
// LCTYPE                DWORD
// LGRPID                DWORD
// LONG                  int32
// LONGLONG              int64
// LONG_PTR              uintptr
// LONG32                int32
// LONG64                int64
// LPARAM                LONG_PTR
// LPBOOL                *BOOL
// LPBYTE                *BYTE
// LPCOLORREF            *COLORREF
// LPCSTR                *int8
// LPCTSTR               LPCWSTR
// LPCVOID               unsafe.Pointer
// LPCWSTR               *WCHAR
// LPDWORD               *DWORD
// LPHANDLE              *HANDLE
// LPINT                 *INT
// LPLONG                *LONG
// LPSTR                 *CHAR
// LPTSTR                LPWSTR
// LPVOID                unsafe.Pointer
// LPWORD                *WORD
// LPWSTR                *WCHAR
// LRESULT               LONG_PTR
// PBOOL                 *BOOL
// PBOOLEAN              *BOOLEAN
// PBYTE                 *BYTE
// PCHAR                 *CHAR
// PCSTR                 *CHAR
// PCTSTR                PCWSTR
// PCWSTR                *WCHAR
// PDWORD                *DWORD
// PDWORDLONG            *DWORDLONG
// PDWORD_PTR            *DWORD_PTR
// PDWORD32              *DWORD32
// PDWORD64              *DWORD64
// PFLOAT                *FLOAT
// PHALF_PTR             *HALF_PTR
// PHANDLE               *HANDLE
// PHKEY                 *HKEY
// PINT_PTR              *INT_PTR
// PINT8                 *INT8
// PINT16                *INT16
// PINT32                *INT32
// PINT64                *INT64
// PLCID                 *LCID
// PLONG                 *LONG
// PLONGLONG             *LONGLONG
// PLONG_PTR             *LONG_PTR
// PLONG32               *LONG32
// PLONG64               *LONG64
// POINTER_32            struct{}
// POINTER_64            struct{}
// POINTER_SIGNED        uintptr
// POINTER_UNSIGNED      uintptr
// PSHORT                *SHORT
// PSIZE_T               *SIZE_T
// PSSIZE_T              *SSIZE_T
// PSTR                  *CHAR
// PTBYTE                *TBYTE
// PTCHAR                *TCHAR
// PTSTR                 PWSTR
// PUCHAR                *UCHAR
// PUHALF_PTR            *UHALF_PTR
// PUINT                 *UINT
// PUINT_PTR             *UINT_PTR
// PUINT8                *UINT8
// PUINT16               *UINT16
// PUINT32               *UINT32
// PUINT64               *UINT64
// PULONG                *ULONG
// PULONGLONG            *ULONGLONG
// PULONG_PTR            *ULONG_PTR
// PULONG32              *ULONG32
// PULONG64              *ULONG64
// PUSHORT               *USHORT
// PVOID                 unsafe.Pointer
// PWCHAR                *WCHAR
// PWORD                 *WORD
// PWSTR                 *WCHAR
// QWORD                 uint64
// SC_HANDLE             HANDLE
// SC_LOCK               LPVOID
// SERVICE_STATUS_HANDLE HANDLE
// SHORT                 int16
// SIZE_T                ULONG_PTR
// SSIZE_T               LONG_PTR
// TBYTE                 WCHAR
// TCHAR                 WCHAR
// UCHAR                 uint8
// UHALF_PTR             struct{}
// UINT                  uint32
// UINT_PTR              uintptr
// UINT8                 uint8
// UINT16                uint16
// UINT32                uint32
// UINT64                uint64
// ULONG                 uint32
// ULONGLONG             uint64
// ULONG_PTR             uintptr
// ULONG32               uint32
// ULONG64               uint64
// USHORT                uint16
// USN                   LONGLONG
// WCHAR                 uint16
// WORD                  uint16
// WPARAM                UINT_PTR

type (
	ATOM            uint16
	BOOL            int32
	COLORREF        uint32
	DWM_FRAME_COUNT uint64
	DWORD           uint32
	HACCEL          HANDLE
	HANDLE          uintptr
	HBITMAP         HANDLE
	HBRUSH          HANDLE
	HCURSOR         HANDLE
	HDC             HANDLE
	HDROP           HANDLE
	HDWP            HANDLE
	HENHMETAFILE    HANDLE
	HFONT           HANDLE
	HGDIOBJ         HANDLE
	HGLOBAL         HANDLE
	HGLRC           HANDLE
	HHOOK           HANDLE
	HICON           HANDLE
	HIMAGELIST      HANDLE
	HINSTANCE       HANDLE
	HKEY            HANDLE
	HKL             HANDLE
	HMENU           HANDLE
	HMODULE         HANDLE
	HMONITOR        HANDLE
	HPEN            HANDLE
	HRESULT         int32
	HRGN            HANDLE
	HRSRC           HANDLE
	HTHUMBNAIL      HANDLE
	HWND            HANDLE
	LPARAM          uintptr
	LPCVOID         unsafe.Pointer
	LRESULT         uintptr
	PVOID           unsafe.Pointer
	QPC_TIME        uint64
	SIZE_T          uintptr
	TRACEHANDLE     uintptr
	ULONG_PTR       uintptr
	WPARAM          uintptr
	ACCESS_MASK     uint32
)

type POINT struct {
	X int32
	Y int32
}

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HICON
}

type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type LOGFONT struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         int32
	Italic         byte
	Underline      byte
	StrikeOut      byte
	CharSet        byte
	OutPrecision   byte
	ClipPrecision  byte
	Quality        byte
	PitchAndFamily byte
	FaceName       [LF_FACESIZE]uint16
}

type OPENFILENAME struct {
	StructSize      uint32
	Owner           HWND
	Instance        HINSTANCE
	Filter          *uint16
	CustomFilter    *uint16
	MaxCustomFilter uint32
	FilterIndex     uint32
	File            *uint16
	MaxFile         uint32
	FileTitle       *uint16
	MaxFileTitle    uint32
	InitialDir      *uint16
	Title           *uint16
	Flags           uint32
	FileOffset      uint16
	FileExtension   uint16
	DefExt          *uint16
	CustData        uintptr
	FnHook          uintptr
	TemplateName    *uint16
	PvReserved      unsafe.Pointer
	DwReserved      uint32
	FlagsEx         uint32
}

type BROWSEINFO struct {
	Owner        HWND
	Root         *uint16
	DisplayName  *uint16
	Title        *uint16
	Flags        uint32
	CallbackFunc uintptr
	LParam       uintptr
	Image        int32
}

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type VARIANT struct {
	VT         uint16 //  2
	WReserved1 uint16 //  4
	WReserved2 uint16 //  6
	WReserved3 uint16 //  8
	Val        int64  // 16
}

type DISPPARAMS struct {
	Rgvarg            uintptr
	RgdispidNamedArgs uintptr
	CArgs             uint32
	CNamedArgs        uint32
}

type EXCEPINFO struct {
	WCode             uint16
	WReserved         uint16
	BstrSource        *uint16
	BstrDescription   *uint16
	BstrHelpFile      *uint16
	DwHelpContext     uint32
	PvReserved        uintptr
	PfnDeferredFillIn uintptr
	Scode             int32
}

type LOGBRUSH struct {
	LbStyle uint32
	LbColor COLORREF
	LbHatch uintptr
}

type DEVMODE struct {
	DmDeviceName       [CCHDEVICENAME]uint16
	DmSpecVersion      uint16
	DmDriverVersion    uint16
	DmSize             uint16
	DmDriverExtra      uint16
	DmFields           uint32
	DmOrientation      int16
	DmPaperSize        int16
	DmPaperLength      int16
	DmPaperWidth       int16
	DmScale            int16
	DmCopies           int16
	DmDefaultSource    int16
	DmPrintQuality     int16
	DmColor            int16
	DmDuplex           int16
	DmYResolution      int16
	DmTTOption         int16
	DmCollate          int16
	DmFormName         [CCHFORMNAME]uint16
	DmLogPixels        uint16
	DmBitsPerPel       uint32
	DmPelsWidth        uint32
	DmPelsHeight       uint32
	DmDisplayFlags     uint32
	DmDisplayFrequency uint32
	DmICMMethod        uint32
	DmICMIntent        uint32
	DmMediaType        uint32
	DmDitherType       uint32
	DmReserved1        uint32
	DmReserved2        uint32
	DmPanningWidth     uint32
	DmPanningHeight    uint32
}

type CIEXYZ struct {
	CiexyzX int32
	CiexyzY int32
	CiexyzZ int32
}

type CIEXYZTRIPLE struct {
	CiexyzRed   CIEXYZ
	CiexyzGreen CIEXYZ
	CiexyzBlue  CIEXYZ
}

type BITMAPV4HEADER struct {
	BITMAPINFOHEADER
	BV4RedMask    uint32
	BV4GreenMask  uint32
	BV4BlueMask   uint32
	BV4AlphaMask  uint32
	BV4CSType     uint32
	BV4Endpoints  CIEXYZTRIPLE
	BV4GammaRed   uint32
	BV4GammaGreen uint32
	BV4GammaBlue  uint32
}

type BITMAPV5HEADER struct {
	BITMAPV4HEADER
	BV5Intent      uint32
	BV5ProfileData uint32
	BV5ProfileSize uint32
	BV5Reserved    uint32
}

type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors *RGBQUAD
}

type BITMAP struct {
	BmType       int32
	BmWidth      int32
	BmHeight     int32
	BmWidthBytes int32
	BmPlanes     uint16
	BmBitsPixel  uint16
	BmBits       unsafe.Pointer
}

type DIBSECTION struct {
	DsBm        BITMAP
	DsBmih      BITMAPINFOHEADER
	DsBitfields [3]uint32
	DshSection  HANDLE
	DsOffset    uint32
}

type ENHMETAHEADER struct {
	IType          uint32
	NSize          uint32
	RclBounds      RECT
	RclFrame       RECT
	DSignature     uint32
	NVersion       uint32
	NBytes         uint32
	NRecords       uint32
	NHandles       uint16
	SReserved      uint16
	NDescription   uint32
	OffDescription uint32
	NPalEntries    uint32
	SzlDevice      SIZE
	SzlMillimeters SIZE
	CbPixelFormat  uint32
	OffPixelFormat uint32
	BOpenGL        uint32
	SzlMicrometers SIZE
}

type SIZE struct {
	CX int32
	CY int32
}

type TEXTMETRIC struct {
	TmHeight           int32
	TmAscent           int32
	TmDescent          int32
	TmInternalLeading  int32
	TmExternalLeading  int32
	TmAveCharWidth     int32
	TmMaxCharWidth     int32
	TmWeight           int32
	TmOverhang         int32
	TmDigitizedAspectX int32
	TmDigitizedAspectY int32
	TmFirstChar        uint16
	TmLastChar         uint16
	TmDefaultChar      uint16
	TmBreakChar        uint16
	TmItalic           byte
	TmUnderlined       byte
	TmStruckOut        byte
	TmPitchAndFamily   byte
	TmCharSet          byte
}

type DOCINFO struct {
	CbSize       int32
	LpszDocName  *uint16
	LpszOutput   *uint16
	LpszDatatype *uint16
	FwType       uint32
}

type NMHDR struct {
	HwndFrom HWND
	IdFrom   uintptr
	Code     uint32
}

type LVCOLUMN struct {
	Mask       uint32
	Fmt        int32
	Cx         int32
	PszText    *uint16
	CchTextMax int32
	ISubItem   int32
	IImage     int32
	IOrder     int32
}

type LVITEM struct {
	Mask       uint32
	IItem      int32
	ISubItem   int32
	State      uint32
	StateMask  uint32
	PszText    *uint16
	CchTextMax int32
	IImage     int32
	LParam     uintptr
	IIndent    int32
	IGroupId   int32
	CColumns   uint32
	PuColumns  uint32
}

type LVHITTESTINFO struct {
	Pt       POINT
	Flags    uint32
	IItem    int32
	ISubItem int32
	IGroup   int32
}

type NMITEMACTIVATE struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    uintptr
	UKeyFlags uint32
}

type NMLISTVIEW struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    uintptr
}

type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

type INITCOMMONCONTROLSEX struct {
	DwSize uint32
	DwICC  uint32
}

type TOOLINFO struct {
	CbSize     uint32
	UFlags     uint32
	Hwnd       HWND
	UId        uintptr
	Rect       RECT
	Hinst      HINSTANCE
	LpszText   *uint16
	LParam     uintptr
	LpReserved unsafe.Pointer
}

type TRACKMOUSEEVENT struct {
	CbSize      uint32
	DwFlags     uint32
	HwndTrack   HWND
	DwHoverTime uint32
}

type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type PAINTSTRUCT struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

type EVENTLOGRECORD struct {
	Length              uint32
	Reserved            uint32
	RecordNumber        uint32
	TimeGenerated       uint32
	TimeWritten         uint32
	EventID             uint32
	EventType           uint16
	NumStrings          uint16
	EventCategory       uint16
	ReservedFlags       uint16
	ClosingRecordNumber uint32
	StringOffset        uint32
	UserSidLength       uint32
	UserSidOffset       uint32
	DataLength          uint32
	DataOffset          uint32
}

type SERVICE_STATUS struct {
	DwServiceType             uint32
	DwCurrentState            uint32
	DwControlsAccepted        uint32
	DwWin32ExitCode           uint32
	DwServiceSpecificExitCode uint32
	DwCheckPoint              uint32
	DwWaitHint                uint32
}

type MODULEENTRY32 struct {
	Size         uint32
	ModuleID     uint32
	ProcessID    uint32
	GlblcntUsage uint32
	ProccntUsage uint32
	ModBaseAddr  *uint8
	ModBaseSize  uint32
	HModule      HMODULE
	SzModule     [MAX_MODULE_NAME32 + 1]uint16
	SzExePath    [MAX_PATH]uint16
}

type PROCESSENTRY32 struct {
	Size            uint32
	CntUsage        uint32
	ProcessID       uint32
	DefaultHeapID   ULONG_PTR
	ModuleID        uint32
	Threads         uint32
	ParentProcessID uint32
	PriClassBase    int32
	Flags           uint32
	ExeFile         [MAX_PATH]uint16
}

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type COORD struct {
	X int16
	Y int16
}

type SMALL_RECT struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         uint16
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

type MARGINS struct {
	CxLeftWidth    int32
	CxRightWidth   int32
	CyTopHeight    int32
	CyBottomHeight int32
}

type DWM_BLURBEHIND struct {
	DwFlags                uint32
	fEnable                BOOL
	hRgnBlur               HRGN
	fTransitionOnMaximized BOOL
}

type DWM_PRESENT_PARAMETERS struct {
	cbSize             uint32
	fQueue             BOOL
	cRefreshStart      DWM_FRAME_COUNT
	cBuffer            uint32
	fUseSourceRate     BOOL
	rateSource         UNSIGNED_RATIO
	cRefreshesPerFrame uint32
	eSampling          DWM_SOURCE_FRAME_SAMPLING
}

type DWM_THUMBNAIL_PROPERTIES struct {
	dwFlags               uint32
	rcDestination         RECT
	rcSource              RECT
	opacity               byte
	fVisible              BOOL
	fSourceClientAreaOnly BOOL
}

type DWM_TIMING_INFO struct {
	cbSize                 uint32
	rateRefresh            UNSIGNED_RATIO
	qpcRefreshPeriod       QPC_TIME
	rateCompose            UNSIGNED_RATIO
	qpcVBlank              QPC_TIME
	cRefresh               DWM_FRAME_COUNT
	cDXRefresh             uint32
	qpcCompose             QPC_TIME
	cFrame                 DWM_FRAME_COUNT
	cDXPresent             uint32
	cRefreshFrame          DWM_FRAME_COUNT
	cFrameSubmitted        DWM_FRAME_COUNT
	cDXPresentSubmitted    uint32
	cFrameConfirmed        DWM_FRAME_COUNT
	cDXPresentConfirmed    uint32
	cRefreshConfirmed      DWM_FRAME_COUNT
	cDXRefreshConfirmed    uint32
	cFramesLate            DWM_FRAME_COUNT
	cFramesOutstanding     uint32
	cFrameDisplayed        DWM_FRAME_COUNT
	qpcFrameDisplayed      QPC_TIME
	cRefreshFrameDisplayed DWM_FRAME_COUNT
	cFrameComplete         DWM_FRAME_COUNT
	qpcFrameComplete       QPC_TIME
	cFramePending          DWM_FRAME_COUNT
	qpcFramePending        QPC_TIME
	cFramesDisplayed       DWM_FRAME_COUNT
	cFramesComplete        DWM_FRAME_COUNT
	cFramesPending         DWM_FRAME_COUNT
	cFramesAvailable       DWM_FRAME_COUNT
	cFramesDropped         DWM_FRAME_COUNT
	cFramesMissed          DWM_FRAME_COUNT
	cRefreshNextDisplayed  DWM_FRAME_COUNT
	cRefreshNextPresented  DWM_FRAME_COUNT
	cRefreshesDisplayed    DWM_FRAME_COUNT
	cRefreshesPresented    DWM_FRAME_COUNT
	cRefreshStarted        DWM_FRAME_COUNT
	cPixelsReceived        uint64
	cPixelsDrawn           uint64
	cBuffersEmpty          DWM_FRAME_COUNT
}

type MilMatrix3x2D struct {
	S_11 float64
	S_12 float64
	S_21 float64
	S_22 float64
	DX   float64
	DY   float64
}

type UNSIGNED_RATIO struct {
	uiNumerator   uint32
	uiDenominator uint32
}

type CREATESTRUCT struct {
	CreateParams uintptr
	Instance     HINSTANCE
	Menu         HMENU
	Parent       HWND
	Cy           int32
	Cx           int32
	Y            int32
	X            int32
	Style        int32
	Name         *uint16
	Class        *uint16
	dwExStyle    uint32
}

type MONITORINFO struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
}

type MONITORINFOEX struct {
	MONITORINFO
	SzDevice [CCHDEVICENAME]uint16
}

type PIXELFORMATDESCRIPTOR struct {
	Size           uint16
	Version        uint16
	DwFlags        uint32
	IPixelType     byte
	ColorBits      byte
	RedBits        byte
	RedShift       byte
	GreenBits      byte
	GreenShift     byte
	BlueBits       byte
	BlueShift      byte
	AlphaBits      byte
	AlphaShift     byte
	AccumBits      byte
	AccumRedBits   byte
	AccumGreenBits byte
	AccumBlueBits  byte
	AccumAlphaBits byte
	DepthBits      byte
	StencilBits    byte
	AuxBuffers     byte
	ILayerType     byte
	Reserved       byte
	DwLayerMask    uint32
	DwVisibleMask  uint32
	DwDamageMask   uint32
}

type INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
	Ki   KEYBDINPUT
	Hi   HARDWAREINPUT
}

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
}

type KbdInput struct {
	typ uint32
	ki  KEYBDINPUT
}

type MouseInput struct {
	typ uint32
	mi  MOUSEINPUT
}

type HardwareInput struct {
	typ uint32
	hi  HARDWAREINPUT
}

type SYSTEMTIME struct {
	Year         uint16
	Month        uint16
	DayOfWeek    uint16
	Day          uint16
	Hour         uint16
	Minute       uint16
	Second       uint16
	Milliseconds uint16
}

type KBDLLHOOKSTRUCT struct {
	VkCode      DWORD
	ScanCode    DWORD
	Flags       DWORD
	Time        DWORD
	DwExtraInfo ULONG_PTR
}

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

type WNDENUMPROC func(HWND, LPARAM) LRESULT

type MEMORY_BASIC_INFORMATION struct {
	BaseAddress       PVOID
	AllocationBase    PVOID
	AllocationProtect DWORD
	RegionSize        SIZE_T
	State             DWORD
	Protect           DWORD
	Type              DWORD
}

type ACL struct {
	AclRevision byte
	Sbz1        byte
	AclSize     uint16
	AceCount    uint16
	Sbz2        uint16
}

type SECURITY_DESCRIPTOR_CONTROL uint16

type SECURITY_DESCRIPTOR struct {
	Revision byte
	Sbz1     byte
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    uintptr
	Group    uintptr
	Sacl     *ACL
	Dacl     *ACL
}

type OVERLAPPED struct {
	Internal     uintptr
	InternalHigh uintptr
	Offset       uint32
	OffsetHigh   uint32
	HEvent       HANDLE
}

type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]byte
}

type SID struct {
	Revision            byte
	SubAuthorityCount   byte
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        uint32
}

type WNODE_HEADER struct {
	BufferSize        uint32
	ProviderId        uint32
	HistoricalContext uint64
	KernelHandle      HANDLE
	Guid              GUID
	ClientContext     uint32
	Flags             uint32
}

func (w WNODE_HEADER) TimeStamp() uint64 {
	return uint64(w.KernelHandle)
}

func (w WNODE_HEADER) Version() uint32 {
	return uint32(w.HistoricalContext >> 32)
}

func (w WNODE_HEADER) Linkage() uint32 {
	return uint32(w.HistoricalContext)
}

type EVENT_TRACE_PROPERTIES struct {
	Wnode               WNODE_HEADER
	BufferSize          uint32
	MinimumBuffers      uint32
	MaximumBuffers      uint32
	MaximumFileSize     uint32
	LogFileMode         uint32
	FlushTimer          uint32
	EnableFlags         uint32
	AgeLimit            int32
	NumberOfBuffers     uint32
	FreeBuffers         uint32
	EventsLost          uint32
	BuffersWritten      uint32
	LogBuffersLost      uint32
	RealTimeBuffersLost uint32
	LoggerThreadId      HANDLE
	LogFileNameOffset   uint32
	LoggerNameOffset    uint32
}

type ALPC_MESSAGE_ATTRIBUTES struct {
	AllocatedAttributes uint32
	ValidAttributes     uint32
}

type ALPC_CONTEXT_ATTR struct {
	PortContext    *AlpcPortContext
	MessageContext uintptr
	Sequence       uint32
	MessageId      uint32
	CallbackId     uint32
}

type ALPC_HANDLE_ATTR struct {
	Flags         uint32
	Handle        HANDLE
	ObjectType    uint32
	DesiredAccess uint32
}

type CLIENT_ID struct {
	UniqueProcess uintptr
	UniqueThread  uintptr
}

type UNICODE_STRING struct {
	Length        uint16
	MaximumLength uint16
	_             [4]byte
	Buffer        *uint16
}

type OBJECT_ATTRIBUTES struct {
	Length                   uint32
	_                        [4]byte
	RootDirectory            HANDLE
	ObjectName               *UNICODE_STRING
	Attributes               uint32
	_                        [4]byte
	SecurityDescriptor       *SECURITY_DESCRIPTOR
	SecurityQualityOfService *SECURITY_QUALITY_OF_SERVICE
}

type PORT_MESSAGE struct {
	DataLength     uint16
	TotalLength    uint16
	Type           uint16
	DataInfoOffset uint16
	ClientId       CLIENT_ID
	MessageId      uint32
	_              [4]byte
	ClientViewSize uint64
}

func (pm PORT_MESSAGE) CallbackId() uint32 {
	return uint32(pm.ClientViewSize >> 32)
}

func (pm PORT_MESSAGE) DoNotUseThisField() float64 {
	panic("we told you not to use this field")
}

const PORT_MESSAGE_SIZE = 0x28

type SECURITY_QUALITY_OF_SERVICE struct {
	Length              uint32
	ImpersonationLevel  uint32
	ContextTrackingMode byte
	EffectiveOnly       byte
	_                   [2]byte // align to 12 bytes
}

const SECURITY_QOS_SIZE = 12

type ALPC_PORT_ATTRIBUTES struct {
	Flags               uint32
	SecurityQos         SECURITY_QUALITY_OF_SERVICE
	MaxMessageLength    uint64 // must be filled out
	MemoryBandwidth     uint64
	MaxPoolUsage        uint64
	MaxSectionSize      uint64
	MaxViewSize         uint64
	MaxTotalSectionSize uint64
	DupObjectTypes      uint32
	Reserved            uint32
}

const SHORT_MESSAGE_MAX_SIZE uint16 = 65535 // MAX_USHORT
const SHORT_MESSAGE_MAX_PAYLOAD = SHORT_MESSAGE_MAX_SIZE - PORT_MESSAGE_SIZE

type AlpcShortMessage struct {
	PORT_MESSAGE
	Data [SHORT_MESSAGE_MAX_PAYLOAD]byte
}

func NewAlpcShortMessage() AlpcShortMessage {
	sm := AlpcShortMessage{}
	sm.TotalLength = SHORT_MESSAGE_MAX_SIZE
	return sm
}

func (sm *AlpcShortMessage) SetData(d []byte) (e error) {

	copy(sm.Data[:], d)
	if len(d) > int(SHORT_MESSAGE_MAX_PAYLOAD) {
		e = errors.New("data too big - truncated")
		sm.DataLength = SHORT_MESSAGE_MAX_PAYLOAD
		sm.TotalLength = SHORT_MESSAGE_MAX_SIZE
		return
	}
	sm.TotalLength = uint16(PORT_MESSAGE_SIZE + len(d))
	sm.DataLength = uint16(len(d))
	return

}

func (sm *AlpcShortMessage) GetData() []byte {
	if int(sm.DataLength) > int(SHORT_MESSAGE_MAX_PAYLOAD) {
		return sm.Data[:] // truncate
	}
	return sm.Data[:sm.DataLength]
}

func (sm *AlpcShortMessage) Reset() {
	sm.PORT_MESSAGE = PORT_MESSAGE{}
	sm.TotalLength = SHORT_MESSAGE_MAX_SIZE
	sm.DataLength = 0
}

type AlpcPortContext struct {
	Handle HANDLE
}

type PROCESS_INFORMATION struct {
	Process   HANDLE
	Thread    HANDLE
	ProcessId uint32
	ThreadId  uint32
}

type STARTUPINFOW struct {
	cb            uint32
	_             *uint16
	Desktop       *uint16
	Title         *uint16
	X             uint32
	Y             uint32
	XSize         uint32
	YSize         uint32
	XCountChars   uint32
	YCountChars   uint32
	FillAttribute uint32
	Flags         uint32
	ShowWindow    uint16
	_             uint16
	_             *uint8
	StdInput      HANDLE
	StdOutput     HANDLE
	StdError      HANDLE
}

type SECURITY_ATTRIBUTES struct {
	Length             uint32
	SecurityDescriptor uintptr
	InheritHandle      BOOL
}

type SYSTEM_INFO struct {
	ProcessorArchitecture     uint16
	Reserved                  uint16
	PageSize                  DWORD
	MinimumApplicationAddress unsafe.Pointer
	MaximumApplicationAddress unsafe.Pointer
	ActiveProcessorMask       ULONG_PTR
	NumberOfProcessors        DWORD
	ProcessorType             DWORD
	AllocationGranularity     DWORD
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

type NOTIFYICONDATAA struct {
	cbSize           DWORD
	hWnd             HWND
	uID              uint32
	uFlags           uint32
	uCallbackMessage uint32
	hIcon            HICON
	szTip            [64]int8
	dwState          DWORD
	dwStateMask      DWORD
	szInfo           [256]int8
	uTimeout         uint32
	uVersion         uint32
	szInfoTitle      [256]int8
	dwInfoFlags      DWORD
	guidItem         GUID
	hBalloonIcon     HICON
}

type WNDPROC func(HWND, uint32, uintptr, uintptr) uintptr

type WNDCLASS struct {
	Style         uint32
	PfnWndProc    WNDPROC
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	Menu          interface{}
	PszClassName  string
	HIconSmall    HICON
}

type _WNDCLASS struct {
	cbSize        uint32
	style         uint32
	pfnWndProcPtr uintptr
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	pszMenuName   *uint16
	pszClassName  *uint16
	hIconSmall    HICON
}
