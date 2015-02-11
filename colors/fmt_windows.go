// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package colors

import (
	"errors"
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
var foreTables = [...]uint16{
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
var backTables = [...]uint16{
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

// 设置控制台颜色。对SetConsoleTextAttribute()的简单包装，
// 使参数更符合go的风格。
func setColor(h syscall.Handle, attr uint16) error {
	r1, _, err := setConsoleTextAttribute.Call(uintptr(h), uintptr(attr))
	if int(r1) == 0 { // setConsoleTextAttribute返回BOOL，而不是bool
		return err
	}

	return nil
}

// 获取颜色值
func getColor(h syscall.Handle) (uint16, error) {
	var csbi console_screen_buffer_info
	r1, _, err := getConsoleScreenBufferInfo.Call(uintptr(h), uintptr(unsafe.Pointer(&csbi)))
	if int(r1) == 0 { // getConsoleScreenBufferInfo返回BOOL，而不是bool
		return 0, err
	}
	return csbi.WAttributes, nil
}

// 根据out获取与之相对应的Handler和writer
func getHW(out int) (syscall.Handle, io.Writer, error) {
	switch out {
	case Stderr:
		return syscall.Stderr, os.Stderr, nil
	case Stdout:
		return syscall.Stdout, os.Stdout, nil
	default:
		return 0, nil, errors.New("无效的输出类型")
	}
}

func print1(out int, attr uint16, v ...interface{}) (size int, err error) {
	h, w, err := getHW(out)
	if err != nil {
		return 0, err
	}

	// 保存原始颜色值
	origin, err := getColor(h)
	if err != nil {
		return 0, err
	}

	// 设置新的颜色值
	if err = setColor(h, attr); err != nil {
		return 0, err
	}

	// 输出字符内容
	if size, err = fmt.Fprint(w, v...); err != nil {
		return size, err
	}

	// 还原原始颜色
	if err = setColor(h, origin); err != nil {
		return 0, err
	}

	return
}

func println1(out int, attr uint16, v ...interface{}) (size int, err error) {
	h, w, err := getHW(out)
	if err != nil {
		return 0, err
	}

	origin, err := getColor(h)
	if err != nil {
		return 0, err
	}

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

func printf(out int, attr uint16, format string, v ...interface{}) (size int, err error) {
	h, w, err := getHW(out)
	if err != nil {
		return 0, err
	}

	origin, err := getColor(h)
	if err != nil {
		return 0, err
	}

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

// 功能同fmt.Print。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Print(out int, foreground, background Color, v ...interface{}) (size int, err error) {
	attr := foreTables[foreground] + backTables[background]
	return print1(out, attr, v...)
}

// 功能同fmt.Println。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Println(out int, foreground, background Color, v ...interface{}) (size int, err error) {
	attr := foreTables[foreground] + backTables[background]
	return println1(out, attr, v...)
}

// 功能同fmt.Printf。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Printf(out int, foreground, background Color, format string, v ...interface{}) (size int, err error) {
	attr := foreTables[foreground] + backTables[background]
	return printf(out, attr, format, v...)
}
