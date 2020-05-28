// SPDX-License-Identifier: MIT

// Package ansi 输出 ansi 控制码，windows 系统默认情况下不支持 ansi 控制码。
// 若仅仅是需要输出彩色字符到控制台，请使用 term/colors 包。
//
//  fmt.Printf("%v这是红色的字", term.FRed)
//  fmt.Printf("%v这是红色字，绿色背景", term.SGR(term.SGRFRed,term.SGRBGreen))
//  fmt.Printf("%v%v这是红色字，绿色背景", term.FRed,term.BGreen)
//
//  // 包装一个stderr。
//  w := term.NewWriter(os.Stderr)
//  w.Left(5)
//  w.SGR(term.SGRFRed)
//  w.Printf("%s", "string")
//  w.Move(1,1)
//
// ansi 的相关文档，可参考以下内容：
//  http://en.wikipedia.org/wiki/ANSI_escape_code
//  http://www.mudpedia.org/mediawiki/index.php/ANSI_colors
package ansi

import (
	"fmt"
	"strconv"
)

// ANSI 码的定义
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

// FColor256 获取扩展的文本颜色值控制码，当 color 的值超出[0,255]时，将触发 panic
func FColor256(color uint8) string {
	return "\033[38;5;" + strconv.Itoa(int(color)) + "m"
}

// BColor256 获取扩展的背景颜色值控制码，当 color 的值超出[0,255]时，将触发 panic
func BColor256(color uint8) string {
	return "\033[48;5;" + strconv.Itoa(int(color)) + "m"
}

// Left 返回左移 N 个字符的控制符
func Left(n int) string {
	return "\033[" + strconv.Itoa(n) + "D"
}

// Right 返回右移 N 个字符的控制符
func Right(n int) string {
	return "\033[" + strconv.Itoa(n) + "C"
}

// Up 返回上移N行的控制符
func Up(n int) string {
	return "\033[" + strconv.Itoa(n) + "A"
}

// Down 返回下移N行的控制符
func Down(n int) string {
	return "\033[" + strconv.Itoa(n) + "B"
}

// Erase 返回清除屏幕的控制符。
//
// n == 0 时，清除从当前光标到屏幕尾的所有字符；
// n == 1 时，清除从当前光标到屏幕头的所有字符；
// n == 2 时，清除当前屏幕的所有字符；
// 当 n 为其它值时，将触发 panic
func Erase(n int) string {
	if n < 0 || n > 2 {
		panic(fmt.Sprintf("n值[%v]必须介于[0,2]", n))
	}
	return "\033[" + strconv.Itoa(n) + "J"
}

// EraseLine 获取清除行的控制符。
//
// n == 0 时，清除从当前光标到行尾的所有字符；
// n == 1 时，清除从当前光标到行头的所有字符；
// n == 2 时，清除当前行的所有字符。
// 当 n 为其它值时，将触发 panic
func EraseLine(n int) string {
	if n < 0 || n > 2 {
		panic(fmt.Sprintf("n值[%v]必须介于[0,2]", n))
	}
	return "\033[" + strconv.Itoa(n) + "K"
}

// Move 移动光标到 x,y 的位置
func Move(x, y int) string {
	//与x;yf相同？
	return "\033[" + strconv.Itoa(x) + ";" + strconv.Itoa(y) + "H"
}
