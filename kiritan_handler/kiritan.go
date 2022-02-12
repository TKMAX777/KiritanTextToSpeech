package kiritan_handler

import (
	"syscall"
	"unsafe"

	"github.com/pkg/errors"

	winapi "github.com/TKMAX777/KiritanWrapper/windows_api"
	"github.com/lxn/win"
)

const windowName = "VOICEROID＋ 東北きりたん EX"

type Handler struct {
	NULL uintptr
	hwnd win.HWND

	controls struct {
		text win.HWND

		play win.HWND
		stop win.HWND
		save win.HWND
		time win.HWND
	}
}

func New() (*Handler, error) {
	var err error
	var handler = Handler{}
	handler.hwnd, err = handler.getHandler(windowName)
	if err != nil {
		handler.hwnd, err = handler.getHandler(windowName + "*")
		if err != nil {
			return nil, errors.Wrap(err, "GetHandler")
		}
	}

	func() {
		defer func() {
			var derr = recover()
			if derr != nil {
				err = errors.Errorf("GetControls: %v", derr)
			}
		}()

		var hwnds = winapi.EnumChildWindows(handler.hwnd)
		hwnds = winapi.EnumChildWindows(hwnds[0])
		hwnds = winapi.EnumChildWindows(hwnds[1])
		hwnds = winapi.EnumChildWindows(hwnds[0])
		hwnds = winapi.EnumChildWindows(hwnds[0])
		hwnds = winapi.EnumChildWindows(hwnds[0])
		hwnds = winapi.EnumChildWindows(hwnds[0])

		handler.controls.text = win.HWND(winapi.EnumChildWindows(hwnds[0])[0])
		handler.controls.play = win.HWND(winapi.EnumChildWindows(hwnds[1])[0])
		handler.controls.stop = win.HWND(winapi.EnumChildWindows(hwnds[1])[1])
		handler.controls.save = win.HWND(winapi.EnumChildWindows(hwnds[1])[2])
		handler.controls.time = win.HWND(winapi.EnumChildWindows(hwnds[1])[3])
	}()

	return &handler, err
}

func (h Handler) getHandler(windowName string) (win.HWND, error) {
	var getHWND win.HWND
	var found bool

	var enumWndProc = func(hwnd win.HWND, lParam *uintptr) uintptr {
		var name = make([]uint16, 1000)
		ok := winapi.GetWindowText(hwnd, uintptr(unsafe.Pointer(&name[0])), 1000)
		if ok == 0 {
			return uintptr(1)
		}

		if syscall.UTF16ToString(name) == windowName {
			getHWND = hwnd
			found = true
			return uintptr(0)
		}

		return uintptr(1)
	}

	winapi.EnumDesktopWindows(win.HANDLE(h.NULL), syscall.NewCallback(enumWndProc), h.NULL)
	if !found {
		return win.HWND(h.NULL), errors.New("NotFound")
	}

	if getHWND == win.HWND(h.NULL) {
		return win.HWND(h.NULL), errors.New("GetWindows")
	}

	return getHWND, nil
}

func (h Handler) SetText(text string) error {
	value, err := syscall.UTF16FromString(text)
	if err != nil {
		return errors.Wrap(err, "NewUTF16Pointer")
	}

	var resRow = winapi.SendMessage(
		h.controls.text,
		win.WM_SETTEXT,
		h.NULL,
		uintptr(unsafe.Pointer(&value[0])),
	)

	var res = int64(resRow)

	switch res {
	default:
		return errors.New("BoxEditingError")
	case win.LB_ERRSPACE:
		return errors.New("BoxInsufficientSpace")
	}
}

func (h Handler) Play() {
	winapi.SendMessage(
		h.controls.play,
		win.BM_CLICK,
		h.NULL,
		h.NULL,
	)
}

func (h Handler) Stop() {
	winapi.SendMessage(
		h.controls.stop,
		win.BM_CLICK,
		h.NULL,
		h.NULL,
	)
}

func (h Handler) Save() {
	winapi.SendMessage(
		h.controls.save,
		win.BM_CLICK,
		h.NULL,
		h.NULL,
	)
}

func (h Handler) TimeMessage() {
	winapi.SendMessage(
		h.controls.time,
		win.BM_CLICK,
		h.NULL,
		h.NULL,
	)
}
