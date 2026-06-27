package pkg

import (
	"os"
	"syscall"
	"unsafe"
)

type terminalSizeEntity struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func TerminalSize() (*terminalSizeEntity, error) {
	ws := &terminalSizeEntity{Col: 99}

	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if err != 0 {
		return nil, err
	}

	return ws, nil
}
