// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

//go:build darwin || dragonfly || freebsd || hurd || linux || netbsd || openbsd

package term

import "golang.org/x/sys/unix"

// Size 获取终端的尺寸
func Size(h int) (width, height int, err error) {
	size, err := unix.IoctlGetWinsize(h, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	return int(size.Col), int(size.Row), nil
}
