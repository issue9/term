// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package colors 带色彩的控制台文本输出包，兼容 windows 平台。
//  // 输出段蓝底红字：foreground:Red;background:Blue
//  colors.Printf(colors.Red, colors.Blue, "foreground:%v;background:%v", colors.Red, colors.Blue)
//
//  // 功能同上，但是可以重复调用 Print* 系列函数输出内容。
//  c := colors.New(colors.Red, colors.Yellow)
//  c.Print("foreground:%v;background:%v")
//  c.Print(colors.Red, colors.Blue)
//
// 为了兼容 windows 平台，只使用了最基本的几种颜色值，
// 而不是 ansi 控制台的 256 色。若不需要考虑 windows 平台，
// 可以直接使用 term/ansi 包，那里有对 ansi 包更好的支持。
//
// 兼容 mingw 等软件。
package colors

import (
	"io"
	"os"
)

// Color 定义了控制台能接受的所有颜色值。
// 具体颜色值在不同的平台上可能有一定的差异。
type Color int8

// 颜色值定义
const (
	Default Color = iota // 默认色，由控制台决定具体颜色值
	Black                // 黑色
	Red                  // 红色
	Green                // 绿色
	Yellow               // 黄色
	Blue                 // 蓝色
	Magenta              // 洋红色
	Cyan                 // 青色
	White                // 白色
	max
)

// IsValid 检测是否为一个有效的 Color 值。
func (c Color) IsValid() bool {
	return c >= Default && c < max
}

func (c Color) String() string {
	switch c {
	case Default:
		return "Default"
	case Black:
		return "Black"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Blue:
		return "Blue"
	case Magenta:
		return "Magenta"
	case Cyan:
		return "Cyan"
	case White:
		return "White"
	default:
		return "<unknown>"
	}
}

// 判断 w 是否为 stderr、stdout、stdin 三者之一
func isConsole(out io.Writer) bool {
	return out == os.Stdout || out == os.Stderr || out == os.Stdin
}
