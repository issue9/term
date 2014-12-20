// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package term

import (
	"fmt"
	"strconv"
)

const (
	Reset           = "\033[0m"
	Bold            = "\033[1m"
	Underline       = "\033[4m"
	Blink           = "\033[5m" // 闪烁
	ReverseVideo    = "\033[7m" // 反显
	Conceal         = "\033[8m"
	BoldOff         = "\033[22m"
	UnderlineOff    = "\033[24m"
	BlinkOff        = "\033[25m"
	ReverseVideoOff = "\033[27m"
	ConcealOff      = "\033[28m"

	FBlack   = "\033[30m"
	FRed     = "\033[31m"
	FGreen   = "\033[32m"
	FYellow  = "\033[33m"
	FBlue    = "\033[34m"
	FMagenta = "\033[35m"
	FCyan    = "\033[36m"
	FWhite   = "\033[37m"
	FDefault = "\033[39m" // 默认前景色

	BBlack   = "\033[40m"
	BRed     = "\033[41m"
	BGreen   = "\033[42m"
	BYellow  = "\033[43m"
	BBlue    = "\033[44m"
	BMagenta = "\033[45m"
	BCyan    = "\033[46m"
	BWhite   = "\033[47m"
	BDefault = "\033[49m" // 默认背景色

	SaveCursor    = "\033[s"    // 保存光标位置
	RestoreCursor = "\033[u"    // 恢复光标位置
	HideCursor    = "\033[?25l" // 隐藏光标
	ShowCursor    = "\033[?25h" // 显示光标
)

// 获取扩展的文本颜色值控制码，当color的值超出[0,255]时，将触发panic
func FColor256(color int) string {
	if color < 0 || color > 255 {
		panic(fmt.Sprintf("颜色值color[%v]只能介于[0,255]之间", color))
	}

	return "\033[38;5;" + strconv.Itoa(color) + "m"
}

// 获取扩展的背景颜色值控制码，当color的值超出[0,255]时，将触发panic
func BColor256(color int) string {
	if color < 0 || color > 255 {
		panic(fmt.Sprintf("颜色值color[%v]只能介于[0,255]之间", color))
	}

	return "\033[48;5;" + strconv.Itoa(color) + "m"
}

// 返回左移N个字符的控制符
func Left(n int) string {
	return "\033[" + strconv.Itoa(n) + "D"
}

// 返回右移N个字符的控制符
func Right(n int) string {
	return "\033[" + strconv.Itoa(n) + "C"
}

// 返回上移N行的控制符
func Up(n int) string {
	return "\033[" + strconv.Itoa(n) + "A"
}

// 返回下移N行的控制符
func Down(n int) string {
	return "\033[" + strconv.Itoa(n) + "B"
}

// 返回清除屏幕的控制符。
// n为0时，清除从当前光标到屏幕尾的所有字符；
// n为1时，清除从当前光标到屏幕头的所有字符；
// n为2时，清除当前屏幕的所有字符。
// 当n为其它值时，将触发panic
func Erase(n int) string {
	if n < 0 || n > 2 {
		panic(fmt.Sprintf("n值[%v]必须介于[0,2]", n))
	}
	return "\033[" + strconv.Itoa(n) + "J"
}

// 返回清除行的控制符。
// n为0时，清除从当前光标到行尾的所有字符；
// n为1时，清除从当前光标到行头的所有字符；
// n为2时，清除当前行的所有字符。
// 当n为其它值时，将触发panic
func EraseLine(n int) string {
	if n < 0 || n > 2 {
		panic(fmt.Sprintf("n值[%v]必须介于[0,2]", n))
	}
	return "\033[" + strconv.Itoa(n) + "K"
}

// 移动光标到x,y的位置
func Move(x, y int) string {
	//与x;yf相同？
	return "\033[" + strconv.Itoa(x) + ";" + strconv.Itoa(y) + "H"
}
