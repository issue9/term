// SPDX-License-Identifier: MIT

// +build !windows

package colors

import (
	"io"
	"os"
)

// 判断 w 是否为 stderr、stdout、stdin 三者之一
func setVirtualTerminalProcessing(w io.Writer, enabled bool) (console bool, err error) {
	return (w == os.Stderr || w == os.Stdout || w == os.Stdin), nil
}
