// SPDX-License-Identifier: MIT

package ansi

import (
	"fmt"
	"io"
)

// Writer ansi 控制码的 io.Writer 接口
//
//  a := NewWriter(os.Stdout)
//
//  a.Left(5)
//  a.ClearLine(2)
//  a.SGR(term.SGRFRed,term.SGRBGreen)
//  a.Print("abc")
//
//  fmt.Fprintf(a, "%v", term.SGRFBBlue)
type Writer struct {
	io.Writer
}

// NewWriter 声明一个 Writer 结构体
func NewWriter(w io.Writer) *Writer {
	return &Writer{Writer: w}
}

// WriteESC 输出字符串
func (a *Writer) WriteESC(esc ESC) (int, error) {
	return a.Write([]byte(esc))
}

// Left 左移 n 个字符光标
func (a *Writer) Left(n int) (int, error) {
	return a.WriteESC(CUB(n))
}

// Right 右移 n 个字符光标
func (a *Writer) Right(n int) (int, error) {
	return a.WriteESC(CUF(n))
}

// Up 上移 n 行光标
func (a *Writer) Up(n int) (int, error) {
	return a.WriteESC(CUU(n))
}

// Down 下移 n 行光标
func (a *Writer) Down(n int) (int, error) {
	return a.WriteESC(CUD(n))
}

// Erase 清除屏幕
//
// n==0 时，清除从当前光标到屏幕尾的所有字符；
// n==1 时，清除从当前光标到屏幕头的所有字符；
// n==2 时，清除当前屏幕的所有字符；
// 当 n 为其它值时，将触发 panic
func (a *Writer) Erase(n int) (int, error) {
	return a.WriteESC(ED(n))
}

// EraseLine 清除行
//
// n==0 时，清除从当前光标到行尾的所有字符；
// n==1 时，清除从当前光标到行头的所有字符；
// n==2 时，清除当前行的所有字符；
// 当 n 为其它值时，将触发 panic
func (a *Writer) EraseLine(n int) (int, error) {
	return a.WriteESC(EL(n))
}

// Move 移动光标到 x,y 的位置
func (a *Writer) Move(x, y int) (int, error) {
	return a.WriteESC(CUP(x, y))
}

// SaveCursor 保存光标位置
func (a *Writer) SaveCursor() (int, error) {
	return a.WriteESC(SCP())
}

// RestoreCursor 恢复光标位置
func (a *Writer) RestoreCursor() (int, error) {
	return a.WriteESC(RCP())
}

// SGR 输出 SGR 指令
func (a *Writer) SGR(sgr ...int) (int, error) {
	return a.WriteESC(SGR(sgr...))
}

// FTrueColor 输出 24 色彩前景色
func (a *Writer) FTrueColor(r, g, b uint8) (int, error) {
	return a.WriteESC(FTrueColor(r, g, b))
}

// BTrueColor 输出 24 色彩背景色
func (a *Writer) BTrueColor(r, g, b uint8) (int, error) {
	return a.WriteESC(BTrueColor(r, g, b))
}

// F256Color 输出 256 色的背景颜色信息
func (a *Writer) F256Color(color uint8) (int, error) {
	return a.WriteESC(F256Color(color))
}

// B256Color 输出 256 色的背景颜色信息
func (a *Writer) B256Color(color uint8) (int, error) {
	return a.WriteESC(B256Color(color))
}

// TrueColor 输出 24 位色彩的颜色信息
//
// fr、fb 和 fb 表示前景色的 RGB 颜色值；
// br、bb 和 bb 表示背景色的 RGB 颜色值；
func (a *Writer) TrueColor(fr, fg, fb, br, bg, bb uint8) (int, error) {
	size, err := a.WriteESC(FTrueColor(fr, fg, fb))
	if size == 0 || err != nil {
		return size, err
	}

	return a.WriteESC(BTrueColor(br, bg, bb))
}

// Color256 输出 256 色的颜色信息
func (a *Writer) Color256(f, b uint8) (int, error) {
	size, err := a.WriteESC(F256Color(f))
	if size == 0 || err != nil {
		return size, err
	}

	return a.WriteESC(B256Color(b))
}

// Printf 相当于 fmt.Printf
func (a *Writer) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(a, format, args...)
}

// Print 相当于 fmt.Print
func (a *Writer) Print(args ...interface{}) (int, error) {
	return fmt.Fprint(a, args...)
}

// Println 相当于 fmt.Println
func (a *Writer) Println(args ...interface{}) (int, error) {
	return fmt.Fprintln(a, args...)
}
