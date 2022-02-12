package windows_api

import "github.com/lxn/win"

func EnumChildWindows(hwnd uintptr) []uintptr {
	var res = []uintptr{}

	var chwnd uintptr
	for chwnd = FindWindowEx(hwnd, chwnd, nil, nil); chwnd != uintptr(0); chwnd = FindWindowEx(hwnd, chwnd, nil, nil) {
		res = append(res, chwnd)
	}

	return res
}

func SendMessage(hwnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	return win.SendMessage(hwnd, msg, wParam, lParam)
}
