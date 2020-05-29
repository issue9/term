// SPDX-License-Identifier: MIT

// Package colors 带色彩的控制台文本输出包
//
// 兼容 windows 平台。
//  // 输出段蓝底红字：foreground:Red;background:Blue
//  colors.Printf(colors.Red, colors.Blue, "foreground:%v;background:%v", colors.Red, colors.Blue)
//
//  // 功能同上，但是可以重复调用 Print* 系列函数输出内容。
//  c := colors.New(colors.Red, colors.Yellow)
//  c.Print("foreground:%v;background:%v")
//  c.Print(colors.Red, colors.Blue)
package colors

import (
	"strconv"

	"github.com/issue9/term/ansi"
)

// Color 定义了控制台能接受的所有颜色值
//
// 具体颜色值在不同的平台上可能有一定的差异。
type Color uint32

// 颜色值定义
const (
	Black         Color = iota // 黑色
	Red                        // 红色
	Green                      // 绿色
	Yellow                     // 黄色
	Blue                       // 蓝色
	Magenta                    // 洋红色
	Cyan                       // 青色
	White                      // 白色
	BrightBlack                // 亮黑
	BrightRed                  // 亮红色
	BrightGreen                // 亮绿色
	BrightYellow               // 亮黄色
	BrightBlue                 // 亮蓝色
	BrightMagenta              // 亮洋红色
	BrightCyan                 // 亮青色
	BrightWhite                // 亮白色
	max

	Default Color = 0

	brightStart   = BrightBlack
	start256Color = 256
)

const (
	redMask   uint32 = 0xf000
	greenMask uint32 = 0x0f00
	blueMask  uint32 = 0x00f0
)

var colorsMap = map[Color]string{
	Black:         "Black",
	Red:           "Red",
	Green:         "Green",
	Yellow:        "Yellow",
	Blue:          "Blue",
	Magenta:       "Magenta",
	Cyan:          "Cyan",
	White:         "White",
	BrightBlack:   "BrightBlack",
	BrightRed:     "BrightRed",
	BrightGreen:   "BrightGreen",
	BrightYellow:  "BrightYellow",
	BrightBlue:    "BrightBlue",
	BrightMagenta: "BrightMagenta",
	BrightCyan:    "BrightCyan",
	BrightWhite:   "BrightWhite",
}

func (c Color) String() string {
	if name, found := colorsMap[c]; found {
		return name
	}

	if c < start256Color {
		return strconv.Itoa(int(c))
	}

	r := (uint32(c) & redMask) >> 24
	g := (uint32(c) & greenMask) >> 16
	b := (uint32(c) & blueMask) >> 8
	return "(" + strconv.Itoa(int(r)) + "," + strconv.Itoa(int(g)) + "," + strconv.Itoa(int(b)) + ")"
}

// RGB 根据 RBG 生成真色彩
func RGB(r, g, b uint8) Color {
	return Color(uint32(r)<<24 + uint32(g)<<16 + uint32(b)<<8)
}

// FColor 转换成前景色的 ansi.ESC
func (c Color) FColor() ansi.ESC {
	switch {
	case c < brightStart:
		return ansi.FColor(uint8(c))
	case c < max:
		return ansi.FBrightColor(uint8(c - brightStart))
	case c < start256Color:
		return ansi.F256Color(uint8(c))
	default:
		r := (uint32(c) & redMask) >> 24
		g := (uint32(c) & greenMask) >> 16
		b := (uint32(c) & blueMask) >> 8
		return ansi.FTrueColor(uint8(r), uint8(g), uint8(b))
	}
}

// BColor 转换成前景色的 ansi.ESC
func (c Color) BColor() ansi.ESC {
	switch {
	case c < brightStart:
		return ansi.BColor(uint8(c))
	case c < max:
		return ansi.BBrightColor(uint8(c - brightStart))
	case c < start256Color:
		return ansi.B256Color(uint8(c))
	default:
		r := (uint32(c) & redMask) >> 24
		g := (uint32(c) & greenMask) >> 16
		b := (uint32(c) & blueMask) >> 8
		return ansi.BTrueColor(uint8(r), uint8(g), uint8(b))
	}
}
