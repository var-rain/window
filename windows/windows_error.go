package windows

import (
	"fmt"
)

type ErrorCode uint32

func (code ErrorCode) Error() string {
	var flags uint32 = FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_ARGUMENT_ARRAY | FORMAT_MESSAGE_IGNORE_INSERTS
	str, err := FormatMessage(flags, nil, uint32(code), 0, nil)
	n := uint32(code)
	if err == nil {
		return fmt.Sprintf("error: %d(0x%08X) - ", n, n) + str
	} else {
		return fmt.Sprintf("error: %d(0x%08X)", n, n)
	}
}

func (hr HRESULT) Succeeded() bool {
	return hr >= 0
}

func (hr HRESULT) Failed() bool {
	return hr < 0
}

func (hr HRESULT) Error() string {
	var flags uint32 = FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_ARGUMENT_ARRAY | FORMAT_MESSAGE_IGNORE_INSERTS
	str, err := FormatMessage(flags, nil, uint32(int32(hr)), 0, nil)
	if err == nil {
		return fmt.Sprintf("HRESULT = %d(0x%08X) - ", int32(hr), uint32(hr)) + str
	} else {
		return fmt.Sprintf("HRESULT = %d(0x%08X)", int32(hr), uint32(hr))
	}
}
