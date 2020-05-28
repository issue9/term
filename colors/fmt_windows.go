// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/windows"
)

// windows 预定义的颜色值
const (
	fDefault   = 0
	fBlue      = 1
	fGreen     = 2
	fCyan      = 3
	fRed       = 4
	fMagenta   = 5
	fYellow    = 6
	fWhite     = 7
	fIntensity = 8

	bDefault   = 0
	bBlue      = 16
	bGreen     = 32
	bCyan      = 48
	bRed       = 64
	bMagenta   = 80
	bYellow    = 96
	bWhite     = 112
	bIntensity = 123

	defColor = fDefault + bDefault
)

// 前景色对照表
var foreTables = []uint16{
	Default: fDefault,
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

var (
	kernel32                = windows.NewLazySystemDLL("kernel32.dll")
	setConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
)

// 设置控制台颜色。对 SetConsoleTextAttribute() 的简单包装，
// 使参数更符合 Go 的风格。
func setColor(h windows.Handle, attr uint16) error {
	r1, _, err := setConsoleTextAttribute.Call(uintptr(h), uintptr(attr))

	if int(r1) == 0 { // setConsoleTextAttribute 返回 BOOL，而不是 bool
		return err
	}
	return nil
}

// 获取颜色值
func getColor(h windows.Handle) (uint16, error) {
	var info windows.ConsoleScreenBufferInfo
	err := windows.GetConsoleScreenBufferInfo(h, &info)
	return info.Attributes, err
}

// 根据 out 获取与之相对应的 Handler 以及是否可以使用颜色
func getHW(out io.Writer) (windows.Handle, bool) {
	switch out {
	case os.Stderr:
		return windows.Stderr, true
	case os.Stdout:
		return windows.Stdout, true
	case os.Stdin:
		return windows.Stdin, true
	default:
		return 0, false
	}
}

// Fprint 带色彩输出的 fmt.Fprint
//
// 颜色值只在 w 为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
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

// Fprintln 带色彩输出的 fmt.Fprintln
//
// 颜色值只在 w 为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
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

// Fprintf 带色彩输出的 fmt.Fprintf
//
// 颜色值只在 w 为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
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

// Print 带色彩输出的 fmt.Print
func Print(foreground, background Color, v ...interface{}) (int, error) {
	return Fprint(os.Stdout, foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println
func Println(foreground, background Color, v ...interface{}) (int, error) {
	return Fprintln(os.Stdout, foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf
func Printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return Fprintf(os.Stdout, foreground, background, format, v...)
}
