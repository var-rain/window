package windows

import (
	"syscall"
)

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func IsErrSuccess(err error) bool {
	if errno, ok := err.(syscall.Errno); ok {
		if errno == 0 {
			return true
		}
	}
	return false
}
