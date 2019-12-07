package windows

import (
	"syscall"
	"unsafe"
)

var (
	ctl                        = syscall.NewLazyDLL("comctl32.dll")
	procImageListAdd           = ctl.NewProc("ImageList_Add")
	procImageListCreate        = ctl.NewProc("ImageList_Create")
	procImageListDestroy       = ctl.NewProc("ImageList_Destroy")
	procImageListGetImageCount = ctl.NewProc("ImageList_GetImageCount")
	procImageListRemove        = ctl.NewProc("ImageList_Remove")
	procImageListReplaceIcon   = ctl.NewProc("ImageList_ReplaceIcon")
	procImageListSetImageCount = ctl.NewProc("ImageList_SetImageCount")
	procInitCommonControlsEx   = ctl.NewProc("InitCommonControlsEx")
	procTrackMouseEvent        = ctl.NewProc("_TrackMouseEvent")
)

func InitCommonControlsEx(lpInitCtrls *INITCOMMONCONTROLSEX) bool {
	ret, _, _ := procInitCommonControlsEx.Call(
		uintptr(unsafe.Pointer(lpInitCtrls)))

	return ret != 0
}

func ImageListCreate(cx int, cy int, flags uint, cInitial int, cGrow int) HIMAGELIST {
	ret, _, _ := procImageListCreate.Call(
		uintptr(cx),
		uintptr(cy),
		uintptr(flags),
		uintptr(cInitial),
		uintptr(cGrow))

	if ret == 0 {
		panic("Create image list failed")
	}

	return HIMAGELIST(ret)
}

func ImageListDestroy(himl HIMAGELIST) bool {
	ret, _, _ := procImageListDestroy.Call(
		uintptr(himl))

	return ret != 0
}

func ImageListGetImageCount(himl HIMAGELIST) int {
	ret, _, _ := procImageListGetImageCount.Call(
		uintptr(himl))

	return int(ret)
}

func ImageListSetImageCount(himl HIMAGELIST, uNewCount uint) bool {
	ret, _, _ := procImageListSetImageCount.Call(
		uintptr(himl),
		uintptr(uNewCount))

	return ret != 0
}

func ImageListAdd(himl HIMAGELIST, hbmImage HBITMAP, hbmMask HBITMAP) int {
	ret, _, _ := procImageListAdd.Call(
		uintptr(himl),
		uintptr(hbmImage),
		uintptr(hbmMask))

	return int(ret)
}

func ImageListReplaceIcon(himl HIMAGELIST, i int, hicon HICON) int {
	ret, _, _ := procImageListReplaceIcon.Call(
		uintptr(himl),
		uintptr(i),
		uintptr(hicon))

	return int(ret)
}

func ImageListAddIcon(himl HIMAGELIST, hicon HICON) int {
	return ImageListReplaceIcon(himl, -1, hicon)
}

func ImageListRemove(himl HIMAGELIST, i int) bool {
	ret, _, _ := procImageListRemove.Call(
		uintptr(himl),
		uintptr(i))

	return ret != 0
}

func ImageListRemoveAll(himl HIMAGELIST) bool {
	return ImageListRemove(himl, -1)
}

func TrackMouseEvent(tme *TRACKMOUSEEVENT) bool {
	ret, _, _ := procTrackMouseEvent.Call(
		uintptr(unsafe.Pointer(tme)))

	return ret != 0
}
