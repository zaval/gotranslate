package utils

import (
	toast "github.com/go-toast/toast"
	"log"
	"syscall"
	"unsafe"
)

const (
	CF_UNICODETEXT = 13
)

var (
	user32DLL        = syscall.NewLazyDLL("user32.dll")
	GetClipboardData = user32DLL.NewProc("GetClipboardData")
	OpenClipboard    = user32DLL.NewProc("OpenClipboard")
	closeClipboard   = user32DLL.NewProc("CloseClipboard")
)

func GetSelectedText() string {

	r, _, err := OpenClipboard.Call(0)
	if r == 0 {
		log.Fatalf("OpenClipboard failed: %v", err)
	}
	defer closeClipboard.Call()
	r, _, err = GetClipboardData.Call(CF_UNICODETEXT)
	if r == 0 {
		log.Fatalf("GetClipboardData failed: %v", err)
	}
	text := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(r))[:])

	return text
}

func ShowNotification(text string) {
	notification := toast.Notification{
		AppID:   "Translate",
		Title:   "Translate",
		Message: text,
	}
	_ = notification.Push()
}
