// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ansi

import (
	"fmt"
	"io"
)

// Writer ansi 控制码的 io.Writer 接口.
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
	w io.Writer
}

// NewWriter 声明一个 Writer 结构体
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (a *Writer) Write(b []byte) (int, error) {
	return a.w.Write(b)
}

// WriteString 输出字符串
func (a *Writer) WriteString(str string) (int, error) {
	return a.Write([]byte(str))
}

// Left 左移 n 个字符光标
func (a *Writer) Left(n int) (int, error) {
	return a.WriteString(Left(n))
}

// Right 右移 n 个字符光标
func (a *Writer) Right(n int) (int, error) {
	return a.WriteString(Right(n))
}

// Up 上移 n 行光标
func (a *Writer) Up(n int) (int, error) {
	return a.WriteString(Up(n))
}

// Down 下移 n 行光标
func (a *Writer) Down(n int) (int, error) {
	return a.WriteString(Down(n))
}

// Erase 清除屏幕。
//
// n==0 时，清除从当前光标到屏幕尾的所有字符；
// n==1 时，清除从当前光标到屏幕头的所有字符；
// n==2 时，清除当前屏幕的所有字符；
// 当 n 为其它值时，将触发 panic
func (a *Writer) Erase(n int) (int, error) {
	return a.WriteString(Erase(n))
}

// EraseLine 清除行。
//
// n==0 时，清除从当前光标到行尾的所有字符；
// n==1 时，清除从当前光标到行头的所有字符；
// n==2 时，清除当前行的所有字符；
// 当 n 为其它值时，将触发 panic
func (a *Writer) EraseLine(n int) (int, error) {
	return a.WriteString(EraseLine(n))
}

// Move 移动光标到 x,y 的位置
func (a *Writer) Move(x, y int) (int, error) {
	return a.WriteString(Move(x, y))
}

// SaveCursor 保存光标位置
func (a *Writer) SaveCursor() (int, error) {
	return a.WriteString(SaveCursor)
}

// RestoreCursor 恢复光标位置
func (a *Writer) RestoreCursor() (int, error) {
	return a.WriteString(RestoreCursor)
}

// HideCursor 隐藏光标
func (a *Writer) HideCursor() (int, error) {
	return a.WriteString(HideCursor)
}

// ShowCursor 显示光标
func (a *Writer) ShowCursor() (int, error) {
	return a.WriteString(ShowCursor)
}

// SGR 输出 SGR 指令
func (a *Writer) SGR(sgr ...string) (int, error) {
	return a.WriteString(SGR(sgr...))
}

// FColor256 输出 256 色的背景颜色信息
func (a *Writer) FColor256(color int) (int, error) {
	return a.WriteString(FColor256(color))
}

// BColor256 输出 256 色的背景颜色信息
func (a *Writer) BColor256(color int) (int, error) {
	return a.WriteString(BColor256(color))
}

// Color256 输出 256 色的颜色信息
func (a *Writer) Color256(f, b int) (int, error) {
	size, err := a.WriteString(FColor256(f))
	if size == 0 || err != nil {
		return size, err
	}

	return a.WriteString(BColor256(b))
}

// Printf 输出普通内容，相当于 fmt.Printf
func (a *Writer) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(a.w, format, args...)
}

// Print 输出普通内容，相当于 fmt.Print
func (a *Writer) Print(args ...interface{}) (int, error) {
	return fmt.Fprint(a.w, args...)
}

// Println 输出普通内容，相当于 fmt.Println
func (a *Writer) Println(args ...interface{}) (int, error) {
	return fmt.Fprintln(a.w, args...)
}
