//go:build windows

package monitor

import (
	"fmt"
	"gomonitor/utils"
	"log"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	mod                     = windows.NewLazyDLL("user32.dll")
	procGetWindowText       = mod.NewProc("GetWindowTextW")
	procGetWindowTextLength = mod.NewProc("GetWindowTextLengthW")
	procGetProcessById      = mod.NewProc("GetWindowThreadProcessId")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func getAppWindowName(hwnd HWND) {
	r1, r2, _ := procGetProcessById.Call(uintptr(hwnd))

	log.Println("Get App Name", r1, r2)

}

func getWindow(funcName string) uintptr {
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}

func RunWindow() {
	for {
		if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
			log.Println("=============================")
			text := GetWindowText(HWND(hwnd))

			fmt.Println("window :", text)
			fmt.Println("# hwnd:", hwnd)

			utils.WriteLog(time.Now().Format(time.RFC3339) + " " + text)

			time.Sleep(3 * time.Second)
		}
	}
}
