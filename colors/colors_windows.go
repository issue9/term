// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

package colors

import "golang.org/x/sys/windows"

// EnableVirtualTerminalProcessing 是否启用 ENABLE_VIRTUAL_TERMINAL_PROCESSING 模式
//
// enable 表示设置之前值，之后可调用 [RestoreVirtualTerminalProcessing] 恢复：
//
//	enable, err := EnableVirtualTerminalProcessing(windows.Stdout)
//	RestoreVirtualTerminalProcessing(enable) // 恢复
//
// windows 下，windows terminal 默认支持 ansi，不需要调用该函数，
// 而 cmd 和 powershell 终端则需要调用该函数启用特性才支持。
func EnableVirtualTerminalProcessing(h windows.Handle) (enable bool, err error) {
	var mode uint32
	if err = windows.GetConsoleMode(h, &mode); err != nil {
		return false, err
	}

	// 已经为 true
	enable = (mode & windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING) != 0
	if enable {
		return true, nil
	}

	return enable, windows.SetConsoleMode(h, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

// RestoreVirtualTerminalProcessing 恢复终端的 ENABLE_VIRTUAL_TERMINAL_PROCESSING 模式
func RestoreVirtualTerminalProcessing(h windows.Handle, enable bool) (err error) {
	var mode uint32
	if err = windows.GetConsoleMode(h, &mode); err != nil {
		return err
	}

	// 不需要作出改变
	old := (mode & windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING) != 0
	if enable == old {
		return nil
	}

	if enable {
		return windows.SetConsoleMode(h, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}
	return windows.SetConsoleMode(h, mode&^windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
