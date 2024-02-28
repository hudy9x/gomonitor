package utils

import (
	"fmt"
	"runtime"
)

func getOS() string {
	os := runtime.GOOS
	fmt.Println(os)

	return os
}

func IsLinux() bool {
	return getOS() == "linux"
}

func IsWindows() bool {
	return getOS() == "windows"
}

func IsMacOs() bool {
	return getOS() == "darwin"
}
