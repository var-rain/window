package win

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	adv                              = syscall.NewLazyDLL("advapi32.dll")
	procCloseEventLog                = adv.NewProc("CloseEventLog")
	procCloseServiceHandle           = adv.NewProc("CloseServiceHandle")
	procControlService               = adv.NewProc("ControlService")
	procControlTrace                 = adv.NewProc("ControlTraceW")
	procInitializeSecurityDescriptor = adv.NewProc("InitializeSecurityDescriptor")
	procOpenEventLog                 = adv.NewProc("OpenEventLogW")
	procOpenSCManager                = adv.NewProc("OpenSCManagerW")
	procOpenService                  = adv.NewProc("OpenServiceW")
	procReadEventLog                 = adv.NewProc("ReadEventLogW")
	procRegCloseKey                  = adv.NewProc("RegCloseKey")
	procRegCreateKeyEx               = adv.NewProc("RegCreateKeyExW")
	procRegDeleteKeyValue            = adv.NewProc("RegDeleteKeyValueW")
	procRegDeleteTree                = adv.NewProc("RegDeleteTreeW")
	procRegDeleteValue               = adv.NewProc("RegDeleteValueW")
	procRegEnumKeyEx                 = adv.NewProc("RegEnumKeyExW")
	procRegGetValue                  = adv.NewProc("RegGetValueW")
	procRegOpenKeyEx                 = adv.NewProc("RegOpenKeyExW")
	procRegSetValueEx                = adv.NewProc("RegSetValueExW")
	procSetSecurityDescriptorDacl    = adv.NewProc("SetSecurityDescriptorDacl")
	procStartService                 = adv.NewProc("StartServiceW")
	procStartTrace                   = adv.NewProc("StartTraceW")
)

func SetSecurityDescriptorDacl(pSecurityDescriptor *SECURITY_DESCRIPTOR, pDacl *ACL) (e error) {

	if pSecurityDescriptor == nil {
		return errors.New("null descriptor")
	}

	var ret uintptr
	if pDacl == nil {
		ret, _, _ = procSetSecurityDescriptorDacl.Call(
			uintptr(unsafe.Pointer(pSecurityDescriptor)),
			uintptr(1), // DaclPresent
			uintptr(0), // pDacl
			uintptr(0), // DaclDefaulted
		)
	} else {
		ret, _, _ = procSetSecurityDescriptorDacl.Call(
			uintptr(unsafe.Pointer(pSecurityDescriptor)),
			uintptr(1), // DaclPresent
			uintptr(unsafe.Pointer(pDacl)),
			uintptr(0), //DaclDefaulted
		)
	}

	if ret != 0 {
		return
	}
	e = syscall.GetLastError()
	return
}

func ControlTrace(hTrace TRACEHANDLE, lpSessionName string, props *EVENT_TRACE_PROPERTIES, dwControl uint32) (success bool, e error) {

	ret, _, _ := procControlTrace.Call(
		uintptr(unsafe.Pointer(hTrace)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpSessionName))),
		uintptr(unsafe.Pointer(props)),
		uintptr(dwControl))

	if ret == ERROR_SUCCESS {
		return true, nil
	}
	e = errors.New(fmt.Sprintf("error: 0x%x", ret))
	return
}

func StartTrace(lpSessionName string, props *EVENT_TRACE_PROPERTIES) (hTrace TRACEHANDLE, e error) {

	ret, _, _ := procStartTrace.Call(
		uintptr(unsafe.Pointer(&hTrace)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpSessionName))),
		uintptr(unsafe.Pointer(props)))

	if ret == ERROR_SUCCESS {
		return
	}
	e = errors.New(fmt.Sprintf("error: 0x%x", ret))
	return
}

func InitializeSecurityDescriptor(rev uint16) (pSecurityDescriptor *SECURITY_DESCRIPTOR, e error) {

	pSecurityDescriptor = &SECURITY_DESCRIPTOR{}

	ret, _, _ := procInitializeSecurityDescriptor.Call(
		uintptr(unsafe.Pointer(pSecurityDescriptor)),
		uintptr(rev),
	)

	if ret != 0 {
		return
	}
	e = syscall.GetLastError()
	return
}

func RegCreateKey(hKey HKEY, subKey string) HKEY {
	var result HKEY
	ret, _, _ := procRegCreateKeyEx.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(KEY_ALL_ACCESS),
		uintptr(0),
		uintptr(unsafe.Pointer(&result)),
		uintptr(0))
	_ = ret
	return result
}

func RegOpenKeyEx(hKey HKEY, subKey string, samDesired uint32) HKEY {
	var result HKEY
	ret, _, _ := procRegOpenKeyEx.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(0),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(&result)))

	if ret != ERROR_SUCCESS {
		panic(fmt.Sprintf("RegOpenKeyEx(%d, %s, %d) failed", hKey, subKey, samDesired))
	}
	return result
}

func RegCloseKey(hKey HKEY) error {
	var err error
	ret, _, _ := procRegCloseKey.Call(
		uintptr(hKey))

	if ret != ERROR_SUCCESS {
		err = errors.New("RegCloseKey failed")
	}
	return err
}

func RegGetRaw(hKey HKEY, subKey string, value string) []byte {
	var bufLen uint32
	var valPtr unsafe.Pointer
	if len(value) > 0 {
		valPtr = unsafe.Pointer(syscall.StringToUTF16Ptr(value))
	}
	_, _, _ = procRegGetValue.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(valPtr),
		uintptr(RRF_RT_ANY),
		0,
		0,
		uintptr(unsafe.Pointer(&bufLen)))

	if bufLen == 0 {
		return nil
	}

	buf := make([]byte, bufLen)
	ret, _, _ := procRegGetValue.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(valPtr),
		uintptr(RRF_RT_ANY),
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&bufLen)))

	if ret != ERROR_SUCCESS {
		return nil
	}

	return buf
}

func RegSetBinary(hKey HKEY, subKey string, value []byte) (errno int) {
	var lPtr, vPtr unsafe.Pointer
	if len(subKey) > 0 {
		lPtr = unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))
	}
	if len(value) > 0 {
		vPtr = unsafe.Pointer(&value[0])
	}
	ret, _, _ := procRegSetValueEx.Call(
		uintptr(hKey),
		uintptr(lPtr),
		uintptr(0),
		uintptr(REG_BINARY),
		uintptr(vPtr),
		uintptr(len(value)))

	return int(ret)
}

func RegSetString(hKey HKEY, subKey string, value string) (errno int) {
	var lPtr, vPtr unsafe.Pointer
	if len(subKey) > 0 {
		lPtr = unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))
	}
	var buf []uint16
	if len(value) > 0 {
		buf, err := syscall.UTF16FromString(value)
		if err != nil {
			return ERROR_BAD_FORMAT
		}
		vPtr = unsafe.Pointer(&buf[0])
	}
	ret, _, _ := procRegSetValueEx.Call(
		uintptr(hKey),
		uintptr(lPtr),
		uintptr(0),
		uintptr(REG_SZ),
		uintptr(vPtr),
		unsafe.Sizeof(buf)+2) // 2 is the size of the terminating null character

	return int(ret)
}

func RegSetUint32(hKey HKEY, subKey string, value uint32) (errno int) {
	var lPtr unsafe.Pointer
	if len(subKey) > 0 {
		lPtr = unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))
	}
	vPtr := unsafe.Pointer(&value)
	ret, _, _ := procRegSetValueEx.Call(
		uintptr(hKey),
		uintptr(lPtr),
		uintptr(0),
		uintptr(REG_DWORD),
		uintptr(vPtr),
		unsafe.Sizeof(value))

	return int(ret)
}

func RegGetString(hKey HKEY, subKey string, value string) string {
	var bufLen uint32
	_, _, _ = procRegGetValue.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),
		uintptr(RRF_RT_REG_SZ),
		0,
		0,
		uintptr(unsafe.Pointer(&bufLen)))

	if bufLen == 0 {
		return ""
	}

	buf := make([]uint16, bufLen)
	ret, _, _ := procRegGetValue.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),
		uintptr(RRF_RT_REG_SZ),
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&bufLen)))

	if ret != ERROR_SUCCESS {
		return ""
	}

	return syscall.UTF16ToString(buf)
}

func RegGetUint32(hKey HKEY, subKey string, value string) (data uint32, errno int) {
	var dataLen = uint32(unsafe.Sizeof(data))
	ret, _, _ := procRegGetValue.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),
		uintptr(RRF_RT_REG_DWORD),
		0,
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(&dataLen)))
	errno = int(ret)
	return
}

func RegDeleteKeyValue(hKey HKEY, subKey string, valueName string) (errno int) {
	ret, _, _ := procRegDeleteKeyValue.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(valueName))))

	return int(ret)
}

func RegDeleteValue(hKey HKEY, valueName string) (errno int) {
	ret, _, _ := procRegDeleteValue.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(valueName))))

	return int(ret)
}

func RegDeleteTree(hKey HKEY, subKey string) (errno int) {
	ret, _, _ := procRegDeleteTree.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))))

	return int(ret)
}

func RegEnumKeyEx(hKey HKEY, index uint32) string {
	var bufLen uint32 = 255
	buf := make([]uint16, bufLen)
	_, _, _ = procRegEnumKeyEx.Call(
		uintptr(hKey),
		uintptr(index),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&bufLen)),
		0,
		0,
		0,
		0)
	return syscall.UTF16ToString(buf)
}

func OpenEventLog(servername string, sourcename string) HANDLE {
	ret, _, _ := procOpenEventLog.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(servername))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(sourcename))))

	return HANDLE(ret)
}

func ReadEventLog(eventLog HANDLE, readFlags uint32, recordOffset uint32, buffer []byte, numberOfBytesToRead uint32, bytesRead *uint32, minNumberOfBytesNeeded *uint32) bool {
	ret, _, _ := procReadEventLog.Call(
		uintptr(eventLog),
		uintptr(readFlags),
		uintptr(recordOffset),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(numberOfBytesToRead),
		uintptr(unsafe.Pointer(bytesRead)),
		uintptr(unsafe.Pointer(minNumberOfBytesNeeded)))

	return ret != 0
}

func CloseEventLog(eventLog HANDLE) bool {
	ret, _, _ := procCloseEventLog.Call(
		uintptr(eventLog))

	return ret != 0
}

func OpenSCManager(lpMachineName string, lpDatabaseName string, dwDesiredAccess uint32) (HANDLE, error) {
	var p1, p2 uintptr
	if len(lpMachineName) > 0 {
		p1 = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpMachineName)))
	}
	if len(lpDatabaseName) > 0 {
		p2 = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpDatabaseName)))
	}
	ret, _, _ := procOpenSCManager.Call(
		p1,
		p2,
		uintptr(dwDesiredAccess))

	if ret == 0 {
		return 0, syscall.GetLastError()
	}

	return HANDLE(ret), nil
}

func CloseServiceHandle(hSCObject HANDLE) error {
	ret, _, _ := procCloseServiceHandle.Call(uintptr(hSCObject))
	if ret == 0 {
		return syscall.GetLastError()
	}
	return nil
}

func OpenService(hSCManager HANDLE, lpServiceName string, dwDesiredAccess uint32) (HANDLE, error) {
	ret, _, _ := procOpenService.Call(
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpServiceName))),
		uintptr(dwDesiredAccess))

	if ret == 0 {
		return 0, syscall.GetLastError()
	}

	return HANDLE(ret), nil
}

func StartService(hService HANDLE, lpServiceArgVectors []string) error {
	l := len(lpServiceArgVectors)
	var ret uintptr
	if l == 0 {
		ret, _, _ = procStartService.Call(
			uintptr(hService),
			0,
			0)
	} else {
		lpArgs := make([]uintptr, l)
		for i := 0; i < l; i++ {
			lpArgs[i] = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpServiceArgVectors[i])))
		}

		ret, _, _ = procStartService.Call(
			uintptr(hService),
			uintptr(l),
			uintptr(unsafe.Pointer(&lpArgs[0])))
	}

	if ret == 0 {
		return syscall.GetLastError()
	}

	return nil
}

func ControlService(hService HANDLE, dwControl uint32, lpServiceStatus *SERVICE_STATUS) bool {
	if lpServiceStatus == nil {
		panic("ControlService:lpServiceStatus cannot be nil")
	}

	ret, _, _ := procControlService.Call(
		uintptr(hService),
		uintptr(dwControl),
		uintptr(unsafe.Pointer(lpServiceStatus)))

	return ret != 0
}
