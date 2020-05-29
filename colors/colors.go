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
	"fmt"
	"strconv"

	"github.com/issue9/term/ansi"
)

// Color 定义了控制台能接受的所有颜色值
//
// 具体颜色值在不同的平台上可能有一定的差异。
type Color int32

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
	maxNamedColor

	Default Color = -1

	brightStart = BrightBlack
	end256Color = 256
)

// Type 字符的各类显示方式
type Type int

// 各类字体控制属性
const (
	Normal Type = -1 // 正常显示

	Bold  Type = iota + 1
	Faint      // 弱化
	Italic
	Underline
	Blink // 闪烁
	RapidBlink
	ReverseVideo // 反显
	Conceal      // 隐藏
	Delete       // 删除线
	maxType
)

const (
	redMask   uint32 = 0x00_ff_00_00
	greenMask uint32 = 0x00_00_ff_00
	blueMask  uint32 = 0x00_00_00_ff
)

var colorsMap = map[Color]string{
	Default:       "Default",
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

	if c < end256Color {
		return strconv.Itoa(int(c))
	}

	r, g, b := c.RGB()
	return "(" + strconv.Itoa(int(r)) + "," + strconv.Itoa(int(g)) + "," + strconv.Itoa(int(b)) + ")"
}

// RGB 根据 RBG 生成真色彩
func RGB(r, g, b uint8) Color {
	return Color(int32(r)<<16 + int32(g)<<8 + int32(b))
}

// FColor 转换成前景色的 ansi.ESC
func (c Color) FColor() ansi.ESC {
	return ansi.CSI('m', c.fColorCode()...)
}

// BColor 转换成前景色的 ansi.ESC
func (c Color) BColor() ansi.ESC {
	return ansi.CSI('m', c.bColorCode()...)
}

// RGB 转换成 RGB 三原色
func (c Color) RGB() (r, g, b uint8) {
	if c < end256Color {
		panic(fmt.Sprintf("颜色值只有大于 %d 的才能转换成 RGB", end256Color))
	}

	r = uint8((uint32(c) & redMask) >> 16)
	g = uint8((uint32(c) & greenMask) >> 8)
	b = uint8((uint32(c) & blueMask))
	return
}

// fColorCode 前景色的 ansi 代码
func (c Color) fColorCode() []int {
	switch {
	case c == -1:
		return []int{39}
	case c < brightStart:
		return []int{30 + int(c)}
	case c < maxNamedColor:
		return []int{90 + int(c-brightStart)}
	case c < end256Color:
		return []int{38, 5, int(c)}
	default:
		r, g, b := c.RGB()
		return []int{38, 2, int(r), int(g), int(b)}
	}
}

// bColorCode 前景色的 ansi 代码
func (c Color) bColorCode() []int {
	switch {
	case c == -1:
		return []int{49}
	case c < brightStart:
		return []int{40 + int(c)}
	case c < maxNamedColor:
		return []int{100 + int(c-brightStart)}
	case c < end256Color:
		return []int{48, 5, int(c)}
	default:
		r, g, b := c.RGB()
		return []int{48, 2, int(r), int(g), int(b)}
	}
}

func isValidType(t Type) bool {
	return t == Normal || (t >= Bold && t < maxType)
}
