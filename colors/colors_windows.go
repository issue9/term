// SPDX-License-Identifier: MIT

package colors

import (
	"io"
	"os"

	"golang.org/x/sys/windows"
)

func enableVirtualTerminalProcessing(h windows.Handle, enable bool) error {
	var mode uint32
	if err := windows.GetConsoleMode(h, &mode); err != nil {
		return err
	}

	if enable {
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	} else {
		mode &^= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	}

	return windows.SetConsoleMode(h, mode)
}

// 判断 w 是否为 stderr、stdout、stdin 三者之一
func setVirtualTerminalProcessing(w io.Writer, enabled bool) (console bool, err error) {
	switch w {
	case os.Stderr:
		return true, enableVirtualTerminalProcessing(windows.Stderr, enabled)
	case os.Stdout:
		return true, enableVirtualTerminalProcessing(windows.Stdout, enabled)
	case os.Stdin:
		return true, enableVirtualTerminalProcessing(windows.Stdin, enabled)
	default:
		return false, nil
	}
}
