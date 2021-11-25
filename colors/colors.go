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
	"math"
	"strconv"

	"github.com/issue9/term/v2/ansi"
)

// Color 定义了控制台能接受的所有颜色值
//
// 具体颜色值在不同的平台上可能有一定的差异。
//
// 颜色定义分为以下几种：
// 默认色： math.MaxInt32
// 基本色： 0-7
// 增强色： 8-15
// 256 色： 0-256，其中 0-15 的数据会被转换成以上的色彩；
// 真彩色： 负数；
//
// 以上这样设置，可以最大限度地保证兼容性。
//
// 默认色、增强色和 256 色基本上所有的终端都支持，
// 而 24 位真彩色则未必所有终端都支持，
// 比如 macOS 自带的终端对该色彩支持并不好。
//
// 关于颜色的具体定义，可参考以下文章：
// https://en.wikipedia.org/wiki/ANSI_escape_code
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

	Default Color = math.MaxInt32

	brightStart = BrightBlack
	end256Color = 256
)

// Type 字符的各类显示方式
type Type int

// 各类字体控制属性
//
// NOTE: 并不是所有的终端都支持这些所有特性。
const (
	Bold         Type = iota + 1
	Faint             // 弱化
	Italic            // 斜体
	Underline         // 下划线
	Blink             // 闪烁
	RapidBlink        // 快速闪烁
	ReverseVideo      // 反显
	Conceal           // 隐藏
	Delete            // 删除线
	maxType

	Normal Type = -1 // 正常显示
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

	if c < 0 { // 负数除了默认值就是 RGB
		r, g, b := c.RGB()
		return "(" + strconv.Itoa(int(r)) + "," + strconv.Itoa(int(g)) + "," + strconv.Itoa(int(b)) + ")"
	}

	if c < end256Color {
		return strconv.Itoa(int(c))
	}

	return "<unknown>"
}

// RGB 根据 RBG 生成真色彩
func RGB(r, g, b uint8) Color {
	return Color(-(int32(r)<<16 + int32(g)<<8 + int32(b)))
}

// HEX 以 16 进制的形式转换成颜色
//
// 可以由以下形式：
//  HEX("#aaa") ==> RGB(0xaa, 0xaa, 0xaa)
//  HEX("aaa") ==> RGB(0xaa, 0xaa, 0xaa)
//  HEX("ababab") ==> RGB(0xab, 0xab, 0xab)
func HEX(hex string) Color {
	if len(hex) == 0 {
		panic(fmt.Sprintf("无效的参数 hex"))
	}

	if hex[0] == '#' {
		hex = hex[1:]
	}

	switch len(hex) {
	case 3:
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	case 6:
	default:
		panic(fmt.Sprintf("无效的参数 hex"))
	}

	c, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	return Color(-c)
}

// FColor 转换成前景色的 ansi.ESC
func (c Color) FColor() ansi.ESC { return ansi.CSI('m', c.fColorCode()...) }

// BColor 转换成前景色的 ansi.ESC
func (c Color) BColor() ansi.ESC { return ansi.CSI('m', c.bColorCode()...) }

// RGB 转换成 RGB 三原色
func (c Color) RGB() (r, g, b uint8) {
	if c > 0 || c == Default {
		panic("这不是 RGB 的颜色标记")
	}

	return uint8((uint32(-c) & redMask) >> 16),
		uint8((uint32(-c) & greenMask) >> 8),
		uint8((uint32(-c) & blueMask))
}

// fColorCode 前景色的 ansi 代码
func (c Color) fColorCode() []int {
	switch {
	case c == Default:
		return []int{39}
	case c < 0:
		r, g, b := c.RGB()
		return []int{38, 2, int(r), int(g), int(b)}
	case c < brightStart:
		return []int{30 + int(c)}
	case c < maxNamedColor:
		return []int{90 + int(c-brightStart)}
	case c < end256Color:
		return []int{38, 5, int(c)}
	default:
		panic("unreached")
	}
}

// bColorCode 前景色的 ansi 代码
func (c Color) bColorCode() []int {
	switch {
	case c == Default:
		return []int{49}
	case c < 0:
		r, g, b := c.RGB()
		return []int{48, 2, int(r), int(g), int(b)}
	case c < brightStart:
		return []int{40 + int(c)}
	case c < maxNamedColor:
		return []int{100 + int(c-brightStart)}
	case c < end256Color:
		return []int{48, 5, int(c)}
	default:
		panic("unreached")
	}
}

func isValidType(t Type) bool {
	return t == Normal || (t >= Bold && t < maxType)
}
