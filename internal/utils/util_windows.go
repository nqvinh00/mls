//go:build windows

package utils

import (
	"syscall"

	"github.com/nqvinh00/mls/models"
)

func EnableColor() error {
	var mode uint32
	if err := syscall.GetConsoleMode(syscall.Stdout, &mode); err != nil {
		return err
	}

	if mode&models.ENABLE_VIRTUAL_TERMINAL_PROCESSING != 0 {
		return nil
	}

	mode |= models.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	ret, _, err := syscall.NewLazyDLL("kernel32.dll").NewProc("SetConsoleMode").Call(uintptr(syscall.Stdout), uintptr(mode))
	if ret == 0 || err != nil {
		return err
	}

	return nil
}
