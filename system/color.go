package system

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	dwmapi                      = windows.NewLazySystemDLL("dwmapi.dll")
	procDwmGetColorizationColor = dwmapi.NewProc("DwmGetColorizationColor")
)

type ColorARGB struct {
	A, R, G, B uint32
}

func (s *SystemApiService) GetWindowsThemeColor() (*ColorARGB, error) {
	var crColorization uint32
	var fOpaqueBlend int32

	ret, _, callErr := procDwmGetColorizationColor.Call(
		uintptr(unsafe.Pointer(&crColorization)),
		uintptr(unsafe.Pointer(&fOpaqueBlend)),
	)

	if ret != 0 { // S_OK = 0
		return nil, fmt.Errorf("DwmGetColorizationColor failed: %v", callErr)
	}

	a := (crColorization >> 24) & 0xFF
	r := (crColorization >> 16) & 0xFF
	g := (crColorization >> 8) & 0xFF
	b := crColorization & 0xFF

	return &ColorARGB{
		A: a,
		R: r,
		G: g,
		B: b,
	}, nil
}
