package windows_api

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output windows_generate.go windows.go
//sys ClipCursor(rect uintptr)(ok int, err error) = user32.ClipCursor
//sys ShowCursor(state bool) (counter int) = user32.ShowCursor
//sys EnumDesktopWindows(hDesktop uintptr, lpEnumFunc uintptr, lParam uintptr) (ok bool) = user32.EnumDesktopWindows
//sys SetLayeredWindowAttributes(hwnd uintptr, color uint32, bAlpha byte, dwFlags uint32) (ok bool) = user32.SetLayeredWindowAttributes
//sys FillRect(hdc uintptr, rect uintptr, hbr uintptr) (ok bool) = user32.FillRect
//sys SetWindowRgn(hwnd uintptr, hRgn uintptr, bRedraw bool) (ok bool) = user32.SetWindowRgn
//sys UpdateLayeredWindow(hwnd uintptr, hdcDst uintptr, pptDst uintptr, psize uintptr, hdcSrc uintptr, pptSrc uintptr, crKey uint32, pblend uintptr, dwFlags uint32) (ok bool) = user32.UpdateLayeredWindow
//sys FindWindowEx(hwndParent uintptr, hwndChildAfter uintptr, lpszClass *uint16, lpszWindow *uint16) (hwnd uintptr) = user32.FindWindowExW
//sys GetWindowText(hwnd uintptr, lpString uintptr, nMax uintptr) (length int) = user32.GetWindowTextW
//sys GetClassName(hwnd uintptr, lpClassName uintptr, nMax uintptr) (length int) = user32.GetClassNameW
//sys SetWindowText(hwnd uintptr, lpString *uint16) (ok bool) = user32.SetWindowTextW
//sys InvalidateRect(hwnd uintptr, rect uintptr, bErase bool) (ok bool) = user32.InvalidateRect

//sys CreateSolidBrush(color uint32) (hbrush uintptr) = Gdi32.CreateSolidBrush
//sys CreatePen(iStyle int, cWidth int, color uint32) (hpen uintptr) = Gdi32.CreatePen
//sys PolyDraw(hdc uintptr, apt uintptr, aj uintptr, cpt int) (ok bool) = Gdi32.PolyDraw
//sys CreateRectRgnIndirect(rect uintptr) (rgn uintptr) = Gdi32.CreateRectRgnIndirect
//sys CreateDIBSection(hdc uintptr, pbmi uintptr, usage uint, ppvBits uintptr, hSection uintptr, offset uint32) (hBitMap uintptr) = Gdi32.CreateDIBSection
