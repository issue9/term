// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ansi

import (
	"fmt"
	"io"
)

// ansi 控制器的 io.Writer 接口.
//
//  a := NewWriter(os.Stdout)
//
//  a.Left(5)
//  a.ClearLine(2)
//  a.SGR(term.SGRFRed,term.SGRBGreen)
//  a.Print("abc")
//
//  fmt.Fprintf(a, "%v", term.SGRFBBlue)
type AnsiWriter struct {
	w io.Writer
}

var _ io.Writer = &AnsiWriter{}

func NewWriter(w io.Writer) *AnsiWriter {
	return &AnsiWriter{w: w}
}

// implements io.Writer
func (a *AnsiWriter) Write(b []byte) (int, error) {
	return a.w.Write(b)
}

func (a *AnsiWriter) WriteString(str string) (int, error) {
	return a.Write([]byte(str))
}

// 左移n个字符光标
func (a *AnsiWriter) Left(n int) (int, error) {
	return a.WriteString(Left(n))
}

// 右移n个字符光标
func (a *AnsiWriter) Right(n int) (int, error) {
	return a.WriteString(Right(n))
}

// 上移n行光标
func (a *AnsiWriter) Up(n int) (int, error) {
	return a.WriteString(Up(n))
}

// 下移n行光标
func (a *AnsiWriter) Down(n int) (int, error) {
	return a.WriteString(Down(n))
}

// 清除屏幕。
// n为0时，清除从当前光标到屏幕尾的所有字符；
// n为1时，清除从当前光标到屏幕头的所有字符；
// n为2时，清除当前屏幕的所有字符。
// 当n为其它值时，将触发panic
func (a *AnsiWriter) Erase(n int) (int, error) {
	return a.WriteString(Erase(n))
}

// 清除行。
// n为0时，清除从当前光标到行尾的所有字符；
// n为1时，清除从当前光标到行头的所有字符；
// n为2时，清除当前行的所有字符。
// 当n为其它值时，将触发panic
func (a *AnsiWriter) EraseLine(n int) (int, error) {
	return a.WriteString(EraseLine(n))
}

// 移动光标到x,y的位置
func (a *AnsiWriter) Move(x, y int) (int, error) {
	return a.WriteString(Move(x, y))
}

func (a *AnsiWriter) SaveCursor() (int, error) {
	return a.WriteString(SaveCursor)
}

func (a *AnsiWriter) RestoreCursor() (int, error) {
	return a.WriteString(RestoreCursor)
}

func (a *AnsiWriter) HideCursor() (int, error) {
	return a.WriteString(HideCursor)
}

func (a *AnsiWriter) ShowCursor() (int, error) {
	return a.WriteString(ShowCursor)
}

func (a *AnsiWriter) SGR(sgr ...string) (int, error) {
	return a.WriteString(SGR(sgr...))
}

func (a *AnsiWriter) FColor256(color int) (int, error) {
	return a.WriteString(FColor256(color))
}

func (a *AnsiWriter) BColor256(color int) (int, error) {
	return a.WriteString(BColor256(color))
}

func (a *AnsiWriter) Color256(f, b int) (int, error) {
	size, err := a.WriteString(FColor256(f))
	if size == 0 || err != nil {
		return size, err
	}

	return a.WriteString(BColor256(b))
}

func (a *AnsiWriter) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(a.w, format, args...)
}

func (a *AnsiWriter) Print(args ...interface{}) (int, error) {
	return fmt.Fprint(a.w, args...)
}

func (a *AnsiWriter) Println(args ...interface{}) (int, error) {
	return fmt.Fprintln(a.w, args...)
}
