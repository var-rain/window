package windows

import (
	"encoding/binary"
	"syscall"
	"unsafe"
)

var (
	kernel                         = syscall.NewLazyDLL("kernel32.dll")
	procExitProcess                = kernel.NewProc("ExitProcess")
	procFormatMessage              = kernel.NewProc("FormatMessageW")
	procCreateNamedPipe            = kernel.NewProc("CreateNamedPipeW")
	procCloseHandle                = kernel.NewProc("CloseHandle")
	procConnectNamedPipe           = kernel.NewProc("ConnectNamedPipe")
	procCreateFileW                = kernel.NewProc("CreateFileW")
	procCreateProcessA             = kernel.NewProc("CreateProcessA")
	procCreateProcessW             = kernel.NewProc("CreateProcessW")
	procCreateRemoteThread         = kernel.NewProc("CreateRemoteThread")
	procCreateToolhelp32Snapshot   = kernel.NewProc("CreateToolhelp32Snapshot")
	procFindResource               = kernel.NewProc("FindResourceW")
	procGetConsoleScreenBufferInfo = kernel.NewProc("GetConsoleScreenBufferInfo")
	procGetConsoleWindow           = kernel.NewProc("GetConsoleWindow")
	procGetCurrentThread           = kernel.NewProc("GetCurrentThread")
	procGetDiskFreeSpaceEx         = kernel.NewProc("GetDiskFreeSpaceExW")
	procGetExitCodeProcess         = kernel.NewProc("GetExitCodeProcess")
	procGetLastError               = kernel.NewProc("GetLastError")
	procGetLogicalDrives           = kernel.NewProc("GetLogicalDrives")
	procGetModuleHandle            = kernel.NewProc("GetModuleHandleW")
	procGetProcAddress             = kernel.NewProc("GetProcAddress")
	procGetProcessTimes            = kernel.NewProc("GetProcessTimes")
	procGetSystemTime              = kernel.NewProc("GetSystemTime")
	procGetSystemTimes             = kernel.NewProc("GetSystemTimes")
	procGetSystemInfo              = kernel.NewProc("GetSystemInfo")
	procGetUserDefaultLCID         = kernel.NewProc("GetUserDefaultLCID")
	procGlobalAlloc                = kernel.NewProc("GlobalAlloc")
	procGlobalFree                 = kernel.NewProc("GlobalFree")
	procGlobalLock                 = kernel.NewProc("GlobalLock")
	procGlobalUnlock               = kernel.NewProc("GlobalUnlock")
	procLoadResource               = kernel.NewProc("LoadResource")
	procLockResource               = kernel.NewProc("LockResource")
	procLstrcpy                    = kernel.NewProc("lstrcpyW")
	procLstrlen                    = kernel.NewProc("lstrlenW")
	procProcess32First             = kernel.NewProc("Process32FirstW")
	procProcess32Next              = kernel.NewProc("Process32NextW")
	procModule32First              = kernel.NewProc("Module32FirstW")
	procModule32Next               = kernel.NewProc("Module32NextW")
	procMoveMemory                 = kernel.NewProc("RtlMoveMemory")
	procMulDiv                     = kernel.NewProc("MulDiv")
	procOpenProcess                = kernel.NewProc("OpenProcess")
	procQueryPerformanceCounter    = kernel.NewProc("QueryPerformanceCounter")
	procQueryPerformanceFrequency  = kernel.NewProc("QueryPerformanceFrequency")
	procReadProcessMemory          = kernel.NewProc("ReadProcessMemory")
	procSetConsoleCtrlHandler      = kernel.NewProc("SetConsoleCtrlHandler")
	procSetConsoleTextAttribute    = kernel.NewProc("SetConsoleTextAttribute")
	procSetSystemTime              = kernel.NewProc("SetSystemTime")
	procSizeofResource             = kernel.NewProc("SizeofResource")
	procTerminateProcess           = kernel.NewProc("TerminateProcess")
	procVirtualAlloc               = kernel.NewProc("VirtualAlloc")
	procVirtualAllocEx             = kernel.NewProc("VirtualAllocEx")
	procVirtualFreeEx              = kernel.NewProc("VirtualFreeEx")
	procVirtualProtect             = kernel.NewProc("VirtualProtect")
	procVirtualQuery               = kernel.NewProc("VirtualQuery")
	procVirtualQueryEx             = kernel.NewProc("VirtualQueryEx")
	procWaitForSingleObject        = kernel.NewProc("WaitForSingleObject")
	procWriteFile                  = kernel.NewProc("WriteFile")
	procWriteProcessMemory         = kernel.NewProc("WriteProcessMemory")
	procResumeThread               = kernel.NewProc("ResumeThread")
	procSuspendThread              = kernel.NewProc("SuspendThread")
)

func ExitProcess(ExitCode uint32) {
	_, _, _ = syscall.Syscall(procExitProcess.Addr(), 1, uintptr(ExitCode), 0, 0)
}

func SuspendThread(hThread HANDLE) (count int, e error) {
	ret, _, lastErr := procSuspendThread.Call(uintptr(hThread))
	if ret == 0xffffffff {
		e = lastErr
		return
	}

	return int(ret), nil
}

func ResumeThread(hThread HANDLE) (count int, e error) {
	ret, _, lastErr := procResumeThread.Call(uintptr(hThread))
	if ret == 0xffffffff {
		e = lastErr
		return
	}

	return int(ret), nil
}

func GetExitCodeProcess(hProcess HANDLE) (code uintptr, e error) {
	ret, _, lastErr := procGetExitCodeProcess.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&code)),
	)

	if ret == 0 {
		e = lastErr
	}

	return
}

func WaitForSingleObject(hHandle HANDLE, msecs uint32) (ok bool, e error) {

	ret, _, lastErr := procWaitForSingleObject.Call(
		uintptr(hHandle),
		uintptr(msecs),
	)

	if ret == WAIT_OBJECT_0 {
		ok = true
		return
	}

	// don't set e for timeouts, or it will be ERROR_SUCCESS which is
	// confusing
	if ret != WAIT_TIMEOUT {
		e = lastErr
	}
	return

}

func CreateProcessW(
	lpApplicationName string,
	lpCommandLine string,
	lpProcessAttributes *SECURITY_ATTRIBUTES,
	lpThreadAttributes *SECURITY_ATTRIBUTES,
	bInheritHandles BOOL,
	dwCreationFlags uint32,
	lpEnvironment unsafe.Pointer,
	lpCurrentDirectory string,
	lpStartupInfo *STARTUPINFOW,
	lpProcessInformation *PROCESS_INFORMATION,
) (e error) {

	var lpAN, lpCL, lpCD *uint16
	if len(lpApplicationName) > 0 {
		lpAN, e = syscall.UTF16PtrFromString(lpApplicationName)
		if e != nil {
			return
		}
	}
	if len(lpCommandLine) > 0 {
		lpCL, e = syscall.UTF16PtrFromString(lpCommandLine)
		if e != nil {
			return
		}
	}
	if len(lpCurrentDirectory) > 0 {
		lpCD, e = syscall.UTF16PtrFromString(lpCurrentDirectory)
		if e != nil {
			return
		}
	}

	ret, _, lastErr := procCreateProcessW.Call(
		uintptr(unsafe.Pointer(lpAN)),
		uintptr(unsafe.Pointer(lpCL)),
		uintptr(unsafe.Pointer(lpProcessAttributes)),
		uintptr(unsafe.Pointer(lpProcessInformation)),
		uintptr(bInheritHandles),
		uintptr(dwCreationFlags),
		uintptr(lpEnvironment),
		uintptr(unsafe.Pointer(lpCD)),
		uintptr(unsafe.Pointer(lpStartupInfo)),
		uintptr(unsafe.Pointer(lpProcessInformation)),
	)

	if ret == 0 {
		e = lastErr
	}

	return
}

func VirtualQuery(lpAddress uintptr, lpBuffer *MEMORY_BASIC_INFORMATION, dwLength int) int {
	ret, _, _ := procVirtualQuery.Call(
		lpAddress,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(dwLength))
	return int(ret)
}

func VirtualQueryEx(hProcess HANDLE, lpAddress uintptr, lpBuffer *MEMORY_BASIC_INFORMATION, dwLength int) int {
	ret, _, _ := procVirtualQueryEx.Call(
		uintptr(hProcess), // The handle to a process.
		lpAddress,
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(dwLength))
	return int(ret)
}

func VirtualProtect(lpAddress uintptr, dwSize int, flNewProtect int, lpflOldProtect *DWORD) bool {
	ret, _, _ := procVirtualProtect.Call(
		lpAddress,
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(unsafe.Pointer(lpflOldProtect)))
	return ret != 0
}

func CreateFile(
	name string,
	desiredAccess DWORD,
	shareMode DWORD,
	securityAttributes *SECURITY_ATTRIBUTES,
	creationDisposition DWORD,
	flagsAndAttributes DWORD,
	templateFile HANDLE,
) (HANDLE, error) {
	handle, _, err := procCreateFileW.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(desiredAccess),
		uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(flagsAndAttributes),
		uintptr(templateFile),
	)
	if !IsErrSuccess(err) {
		return HANDLE(handle), err
	}
	return HANDLE(handle), err
}

func CreateProcessA(lpApplicationName *string,
	lpCommandLine string,
	lpProcessAttributes *syscall.SecurityAttributes,
	lpThreadAttributes *syscall.SecurityAttributes,
	bInheritHandles bool,
	dwCreationFlags uint32,
	lpEnvironment *string,
	lpCurrentDirectory *uint16,
	lpStartupInfo *syscall.StartupInfo,
	lpProcessInformation *syscall.ProcessInformation) {

	inherit := 0
	if bInheritHandles {
		inherit = 1
	}

	_, _, _ = procCreateProcessA.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(*lpApplicationName))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpCommandLine))),
		uintptr(unsafe.Pointer(lpProcessAttributes)),
		uintptr(unsafe.Pointer(lpThreadAttributes)),
		uintptr(inherit),
		uintptr(dwCreationFlags),
		uintptr(unsafe.Pointer(lpEnvironment)),
		uintptr(unsafe.Pointer(lpCurrentDirectory)),
		uintptr(unsafe.Pointer(lpStartupInfo)),
		uintptr(unsafe.Pointer(lpProcessInformation)))
}

func VirtualAllocEx(hProcess HANDLE, lpAddress int, dwSize int, flAllocationType int, flProtect int) (addr uintptr, err error) {
	ret, _, err := procVirtualAllocEx.Call(
		uintptr(hProcess),  // The handle to a process.
		uintptr(lpAddress), // The pointer that specifies a desired starting address for the region of pages that you want to allocate.
		uintptr(dwSize),    // The size of the region of memory to allocate, in bytes.
		uintptr(flAllocationType),
		uintptr(flProtect))
	if int(ret) == 0 {
		return ret, err
	}
	return ret, nil
}

func VirtualAlloc(lpAddress int, dwSize int, flAllocationType int, flProtect int) (addr uintptr, err error) {
	ret, _, err := procVirtualAlloc.Call(
		uintptr(lpAddress), // The starting address of the region to allocate
		uintptr(dwSize),    // The size of the region of memory to allocate, in bytes.
		uintptr(flAllocationType),
		uintptr(flProtect))
	if int(ret) == 0 {
		return ret, err
	}
	return ret, nil
}

func VirtualFreeEx(hProcess HANDLE, lpAddress, dwSize uintptr, dwFreeType uint32) bool {
	ret, _, _ := procVirtualFreeEx.Call(
		uintptr(hProcess),
		lpAddress,
		dwSize,
		uintptr(dwFreeType),
	)

	return ret != 0
}

func GetProcAddress(hProcess HANDLE, procname string) (addr uintptr, err error) {
	var pn uintptr

	if procname == "" {
		pn = 0
	} else {
		pn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(procname)))
	}

	ret, _, err := procGetProcAddress.Call(uintptr(hProcess), pn)
	if int(ret) == 0 {
		return ret, err
	}
	return ret, nil
}

func CreateRemoteThread(hprocess HANDLE, sa *syscall.SecurityAttributes,
	stackSize uint32, startAddress uint32, parameter uintptr, creationFlags uint32) (HANDLE, uint32, error) {
	var threadId uint32
	r1, _, e1 := procCreateRemoteThread.Call(
		uintptr(hprocess),
		uintptr(unsafe.Pointer(sa)),
		uintptr(stackSize),
		uintptr(startAddress),
		parameter,
		uintptr(creationFlags),
		uintptr(unsafe.Pointer(&threadId)))

	if int(r1) == 0 {
		return INVALID_HANDLE, 0, e1
	}
	return HANDLE(r1), threadId, nil
}

func GetModuleHandle(modulename string) HINSTANCE {
	var mn uintptr
	if modulename == "" {
		mn = 0
	} else {
		mn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(modulename)))
	}
	ret, _, _ := procGetModuleHandle.Call(mn)
	return HINSTANCE(ret)
}

func MulDiv(number, numerator, denominator int) int {
	ret, _, _ := procMulDiv.Call(
		uintptr(number),
		uintptr(numerator),
		uintptr(denominator))

	return int(ret)
}

func GetConsoleWindow() HWND {
	ret, _, _ := procGetConsoleWindow.Call()

	return HWND(ret)
}

func GetCurrentThread() HANDLE {
	ret, _, _ := procGetCurrentThread.Call()

	return HANDLE(ret)
}

func GetLogicalDrives() uint32 {
	ret, _, _ := procGetLogicalDrives.Call()

	return uint32(ret)
}

func GetUserDefaultLCID() uint32 {
	ret, _, _ := procGetUserDefaultLCID.Call()

	return uint32(ret)
}

func Lstrlen(lpString *uint16) int {
	ret, _, _ := procLstrlen.Call(uintptr(unsafe.Pointer(lpString)))

	return int(ret)
}

func Lstrcpy(buf []uint16, lpString *uint16) {
	_, _, _ = procLstrcpy.Call(
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(lpString)))
}

func GlobalAlloc(uFlags uint, dwBytes uint32) HGLOBAL {
	ret, _, _ := procGlobalAlloc.Call(
		uintptr(uFlags),
		uintptr(dwBytes))

	if ret == 0 {
		panic("GlobalAlloc failed")
	}

	return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) {
	ret, _, _ := procGlobalFree.Call(uintptr(hMem))

	if ret != 0 {
		panic("GlobalFree failed")
	}
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
	ret, _, _ := procGlobalLock.Call(uintptr(hMem))

	if ret == 0 {
		panic("GlobalLock failed")
	}

	return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) bool {
	ret, _, _ := procGlobalUnlock.Call(uintptr(hMem))

	return ret != 0
}

func MoveMemory(destination, source unsafe.Pointer, length uint32) {
	_, _, _ = procMoveMemory.Call(
		uintptr(destination),
		uintptr(source),
		uintptr(length))
}

func FindResource(hModule HMODULE, lpName *uint16, lpType *uint16) (HRSRC, error) {
	ret, _, _ := procFindResource.Call(
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(lpType)))

	if ret == 0 {
		return 0, syscall.GetLastError()
	}

	return HRSRC(ret), nil
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC) uint32 {
	ret, _, _ := procSizeofResource.Call(
		uintptr(hModule),
		uintptr(hResInfo))

	if ret == 0 {
		panic("SizeofResource failed")
	}

	return uint32(ret)
}

func LockResource(hResData HGLOBAL) unsafe.Pointer {
	ret, _, _ := procLockResource.Call(uintptr(hResData))

	if ret == 0 {
		panic("LockResource failed")
	}

	return unsafe.Pointer(ret)
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) HGLOBAL {
	ret, _, _ := procLoadResource.Call(
		uintptr(hModule),
		uintptr(hResInfo))

	if ret == 0 {
		panic("LoadResource failed")
	}

	return HGLOBAL(ret)
}

func GetLastError() uint32 {
	ret, _, _ := procGetLastError.Call()
	return uint32(ret)
}

func OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) (handle HANDLE, err error) {
	inherit := 0
	if inheritHandle {
		inherit = 1
	}

	ret, _, err := procOpenProcess.Call(
		uintptr(desiredAccess),
		uintptr(inherit),
		uintptr(processId))
	if err != nil && IsErrSuccess(err) {
		err = nil
	}
	handle = HANDLE(ret)
	return
}

func TerminateProcess(hProcess HANDLE, uExitCode uint) bool {
	ret, _, _ := procTerminateProcess.Call(
		uintptr(hProcess),
		uintptr(uExitCode))
	return ret != 0
}

func CloseHandle(object HANDLE) bool {
	ret, _, _ := procCloseHandle.Call(
		uintptr(object))
	return ret != 0
}

func CreateToolhelp32Snapshot(flags uint32, processId uint32) HANDLE {
	ret, _, _ := procCreateToolhelp32Snapshot.Call(
		uintptr(flags),
		uintptr(processId))

	if ret <= 0 {
		return HANDLE(0)
	}

	return HANDLE(ret)
}

func Process32First(snapshot HANDLE, pe *PROCESSENTRY32) bool {
	if pe.Size == 0 {
		pe.Size = uint32(unsafe.Sizeof(*pe))
	}
	ret, _, _ := procProcess32First.Call(
		uintptr(snapshot),
		uintptr(unsafe.Pointer(pe)))

	return ret != 0
}

func Process32Next(snapshot HANDLE, pe *PROCESSENTRY32) bool {
	if pe.Size == 0 {
		pe.Size = uint32(unsafe.Sizeof(*pe))
	}
	ret, _, _ := procProcess32Next.Call(
		uintptr(snapshot),
		uintptr(unsafe.Pointer(pe)))

	return ret != 0
}

func Module32First(snapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32First.Call(
		uintptr(snapshot),
		uintptr(unsafe.Pointer(me)))

	return ret != 0
}

func Module32Next(snapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32Next.Call(
		uintptr(snapshot),
		uintptr(unsafe.Pointer(me)))

	return ret != 0
}

func GetSystemTimes(lpIdleTime *FILETIME, lpKernelTime *FILETIME, lpUserTime *FILETIME) bool {
	ret, _, _ := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(lpIdleTime)),
		uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)))

	return ret != 0
}

// GetSystemInfo retrieves information about the current system.
func GetSystemInfo(sysinfo *SYSTEM_INFO) {
	_, _, _ = procGetSystemInfo.Call(uintptr(unsafe.Pointer(sysinfo)))
}

func GetProcessTimes(hProcess HANDLE, lpCreationTime *FILETIME, lpExitTime *FILETIME, lpKernelTime *FILETIME, lpUserTime *FILETIME) bool {
	ret, _, _ := procGetProcessTimes.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpCreationTime)),
		uintptr(unsafe.Pointer(lpExitTime)),
		uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)))

	return ret != 0
}

func GetConsoleScreenBufferInfo(hConsoleOutput HANDLE) *CONSOLE_SCREEN_BUFFER_INFO {
	var csbi CONSOLE_SCREEN_BUFFER_INFO
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}

func SetConsoleTextAttribute(hConsoleOutput HANDLE, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttributes))
	return ret != 0
}

func GetDiskFreeSpaceEx(dirName string) (r bool,
	freeBytesAvailable uint64, totalNumberOfBytes uint64, totalNumberOfFreeBytes uint64) {
	ret, _, _ := procGetDiskFreeSpaceEx.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(dirName))),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)))
	return ret != 0,
		freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes
}

func GetSystemTime() (time SYSTEMTIME, err error) {
	_, _, err = procGetSystemTime.Call(
		uintptr(unsafe.Pointer(&time)))
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

func SetSystemTime(time *SYSTEMTIME) (err error) {
	_, _, err = procSetSystemTime.Call(
		uintptr(unsafe.Pointer(time)))
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

func WriteFile(handle HANDLE, buf []byte, written *uint32, overlapped *OVERLAPPED) (bool, error) {
	ok, _, err := procWriteFile.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(&written)),
		uintptr(unsafe.Pointer(overlapped)),
	)
	if !IsErrSuccess(err) {
		return ok != 0, err
	}
	return ok != 0, err
}

//Writes data to an area of memory in a specified process. The entire area to be written to must be accessible or the operation fails.
func WriteProcessMemory(hProcess HANDLE, lpBaseAddress uint32, data []byte, size uint) (err error) {
	var numBytesRead uintptr

	_, _, err = procWriteProcessMemory.Call(uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(size),
		uintptr(unsafe.Pointer(&numBytesRead)))
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

//Write process memory with a source of uint32
func WriteProcessMemoryAsUint32(hProcess HANDLE, lpBaseAddress uint32, data uint32) (err error) {

	bData := make([]byte, 4)
	binary.LittleEndian.PutUint32(bData, data)
	err = WriteProcessMemory(hProcess, lpBaseAddress, bData, 4)
	if err != nil {
		return
	}
	return
}

//Reads data from an area of memory in a specified process. The entire area to be read must be accessible or the operation fails.
func ReadProcessMemory(hProcess HANDLE, lpBaseAddress uint32, size uint) (data []byte, err error) {
	var numBytesRead uintptr
	data = make([]byte, size)

	_, _, err = procReadProcessMemory.Call(uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(size),
		uintptr(unsafe.Pointer(&numBytesRead)))
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

// Read process memory and convert the returned data to uint32
func ReadProcessMemoryAsUint32(hProcess HANDLE, lpBaseAddress uint32) (buffer uint32, err error) {
	data, err := ReadProcessMemory(hProcess, lpBaseAddress, 4)
	if err != nil {
		return
	}
	buffer = binary.LittleEndian.Uint32(data)
	return
}

//Adds or removes an application-defined HandlerRoutine function from the list of handler functions for the calling process.
func SetConsoleCtrlHandler(handlerRoutine func(DWORD) int32, add uint) (err error) {
	_, _, err = procSetConsoleCtrlHandler.Call(uintptr(unsafe.Pointer(&handlerRoutine)),
		uintptr(add))
	if !IsErrSuccess(err) {
		return
	}
	err = nil
	return
}

func QueryPerformanceCounter() uint64 {
	result := uint64(0)
	_, _, _ = procQueryPerformanceCounter.Call(
		uintptr(unsafe.Pointer(&result)),
	)

	return result
}

func QueryPerformanceFrequency() uint64 {
	result := uint64(0)
	_, _, _ = procQueryPerformanceFrequency.Call(
		uintptr(unsafe.Pointer(&result)),
	)

	return result
}
