//go:build windows

package common

import (
	"syscall"
	"unsafe"
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	procMessageBoxW = user32.NewProc("MessageBoxW")
)

const (
	MB_OK              = 0x00000000
	MB_ICONINFORMATION = 0x00000040
	MB_ICONERROR       = 0x00000010
)

func ShowMessage(title, message string) {
	showMessageBox(title, message, MB_OK|MB_ICONINFORMATION)
}

func ShowError(title, message string) {
	showMessageBox(title, message, MB_OK|MB_ICONERROR)
}

func showMessageBox(title, message string, style uintptr) {
	t, _ := syscall.UTF16PtrFromString(title)
	m, _ := syscall.UTF16PtrFromString(message)
	procMessageBoxW.Call(0, uintptr(unsafe.Pointer(m)), uintptr(unsafe.Pointer(t)), style)
}
