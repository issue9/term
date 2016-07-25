// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package colors

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"unsafe"
)

// windows 预定义的颜色值
const (
	fBlue      = 1
	fGreen     = 2
	fRed       = 4
	fIntensity = 8

	bBlue      = 16
	bGreen     = 32
	bRed       = 64
	bIntensity = 123

	// 增强前景色
	fYellow  = fRed + fGreen
	fCyan    = fGreen + fBlue
	fMagenta = fBlue + fRed
	fWhite   = fRed + fBlue + fGreen
	fDefault = fWhite

	// 增强背景色
	bYellow  = bRed + bGreen
	bCyan    = bGreen + bBlue
	bMagenta = bBlue + bRed
	bWhite   = bRed + bBlue + bGreen
	bDefault = 0 // 默认背景为黑

	defColor = fDefault + bDefault
)

// 前景色对照表
var foreTables = []uint16{
	Default: fWhite,
	Black:   0,
	Red:     fRed,
	Green:   fGreen,
	Yellow:  fYellow,
	Blue:    fBlue,
	Magenta: fMagenta,
	Cyan:    fCyan,
	White:   fWhite,
}

// 背景色对照表
var backTables = []uint16{
	Default: 0,
	Black:   0,
	Red:     bRed,
	Green:   bGreen,
	Yellow:  bYellow,
	Blue:    bBlue,
	Magenta: bMagenta,
	Cyan:    bCyan,
	White:   bWhite,
}

type coord struct {
	X, Y int16
}

type small_rect struct {
	Left, Top, Right, Bottom int16
}

type console_screen_buffer_info struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            small_rect
	DwMaximumWindowSize coord
}

var (
	kernel32                   = syscall.NewLazyDLL("kernel32.dll")
	setConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
	getConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

// 设置控制台颜色。对 SetConsoleTextAttribute() 的简单包装，
// 使参数更符合 Go 的风格。
func setColor(h syscall.Handle, attr uint16) error {
	r1, _, err := setConsoleTextAttribute.Call(uintptr(h), uintptr(attr))
	if int(r1) == 0 { // setConsoleTextAttribute 返回 BOOL，而不是 bool
		return err
	}

	return nil
}

// 获取颜色值
func getColor(h syscall.Handle) (uint16, error) {
	var csbi console_screen_buffer_info
	r1, _, err := getConsoleScreenBufferInfo.Call(uintptr(h), uintptr(unsafe.Pointer(&csbi)))
	if int(r1) == 0 { // getConsoleScreenBufferInfo 返回 BOOL，而不是 bool
		return 0, err
	}
	return csbi.WAttributes, nil
}

// 根据 out 获取与之相对应的 Handler 以及是否可以使用颜色
func getHW(out io.Writer) (syscall.Handle, bool) {
	o, ok := out.(*os.File)
	if !ok {
		return 0, false
	}

	switch o {
	case os.Stderr:
		return syscall.Stderr, true
	case os.Stdout:
		return syscall.Stdout, true
	case os.Stdin:
		return syscall.Stdin, true
	default:
		return 0, false
	}
}

// Fprint 带色彩输出的 fmt.Fprint。
//
// foreground，background 为输出文字的前景和背景色。
func Fprint(w io.Writer, foreground, background Color, v ...interface{}) (size int, err error) {
	h, ok := getHW(w)
	if !ok {
		return fmt.Fprint(w, v...)
	}

	origin, err := getColor(h)
	if err != nil {
		return 0, err
	}

	attr := foreTables[foreground] + backTables[background]
	if err = setColor(h, attr); err != nil {
		return 0, err
	}

	if size, err = fmt.Fprint(w, v...); err != nil {
		return size, err
	}

	if err = setColor(h, origin); err != nil {
		return 0, err
	}

	return
}

// Fprintln 带色彩输出的 fmt.Fprintln。
//
// foreground，background 为输出文字的前景和背景色。
func Fprintln(w io.Writer, foreground, background Color, v ...interface{}) (size int, err error) {
	h, ok := getHW(w)
	if !ok {
		return fmt.Fprintln(w, v...)
	}

	origin, err := getColor(h)
	if err != nil {
		return 0, err
	}

	attr := foreTables[foreground] + backTables[background]
	if err = setColor(h, attr); err != nil {
		return 0, err
	}

	if size, err = fmt.Fprintln(w, v...); err != nil {
		return size, err
	}

	if err = setColor(h, origin); err != nil {
		return 0, err
	}

	return
}

// Fprintf 带色彩输出的 fmt.Fprintf。
//
// foreground，background 为输出文字的前景和背景色。
func Fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (size int, err error) {
	h, ok := getHW(w)
	if !ok {
		return fmt.Fprintf(w, format, v...)
	}

	origin, err := getColor(h)
	if err != nil {
		return 0, err
	}

	attr := foreTables[foreground] + backTables[background]
	if err = setColor(h, attr); err != nil {
		return 0, err
	}

	if size, err = fmt.Fprintf(w, format, v...); err != nil {
		return size, err
	}

	if err = setColor(h, origin); err != nil {
		return 0, err
	}

	return
}

// Print 带色彩输出的 fmt.Print，输出到 os.Stdout。
func Print(foreground, background Color, v ...interface{}) (int, error) {
	return Fprint(os.Stdout, foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println，输出到 os.Stdout。
func Println(foreground, background Color, v ...interface{}) (int, error) {
	return Fprintln(os.Stdout, foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf，输出到 os.Stdout。
func Printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return Fprintf(os.Stdout, foreground, background, format, v...)
}

// Print 带色彩输出的 fmt.Print，在 windows 下会忽略颜色值的定义。
func Sprint(foreground, background Color, v ...interface{}) string {
	return fmt.Sprint(v...)
}

// Println 带色彩输出的 fmt.Println，在 windows 下会忽略颜色值的定义。
func Sprintln(foreground, background Color, v ...interface{}) string {
	return fmt.Sprintln(v...)
}

// Printf 带色彩输出的 fmt.Printf，在 windows 下会忽略颜色值的定义。
func Sprintf(foreground, background Color, format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}
