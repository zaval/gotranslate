package utils

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

var (
	user32    = syscall.NewLazyDLL("user32.dll")
	reghotkey = user32.NewProc("RegisterHotKey")
	peekmsg   = user32.NewProc("PeekMessageW")
)

type Hotkey struct {
	Id        int // Unique id
	Modifiers int // Mask of modifiers
	KeyCode   int // Key code, e.g. 'A'
}

type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

func RegisterHotkey() {

	hotkey := &Hotkey{1, ModAlt + ModCtrl, 'X'}
	res, _, err := reghotkey.Call(0, uintptr(hotkey.Id), uintptr(hotkey.Modifiers), uintptr(hotkey.KeyCode))
	if res == 1 {
		fmt.Println("ctrl+alt+X registered")
	} else {
		fmt.Println("Failed to register, error:", err)
		return
	}

	for {
		var msg = &MSG{}
		peekmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)
		fmt.Println(msg)
		if id := msg.WPARAM; id != 0 {
			fmt.Println(msg)
			translated, err := Translate(GetSelectedText())
			if err == nil {
				ShowNotification(translated)
			}
		}
		time.Sleep(time.Millisecond * 100)
	}

}
