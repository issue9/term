// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package colors

import (
	"fmt"
	"io"
	"os"

	"github.com/issue9/term/ansi"
)

// 前景色对照表
var ansiForeTables = []string{
	Default: ansi.FDefault,
	Black:   ansi.FBlack,
	Red:     ansi.FRed,
	Green:   ansi.FGreen,
	Yellow:  ansi.FYellow,
	Blue:    ansi.FBlue,
	Magenta: ansi.FMagenta,
	Cyan:    ansi.FCyan,
	White:   ansi.FWhite,
}

// 背景色对照表
var ansiBackTables = []string{
	Default: ansi.BDefault,
	Black:   ansi.BBlack,
	Red:     ansi.BRed,
	Green:   ansi.BGreen,
	Yellow:  ansi.BYellow,
	Blue:    ansi.BBlue,
	Magenta: ansi.BMagenta,
	Cyan:    ansi.BCyan,
	White:   ansi.BWhite,
}

// Fprint 带色彩输出的 fmt.Fprint，颜色值被转换成 ANSI 码一起写入到 w 中。
func fprint(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprint(foreground, background, v...))
}

// Fprintln 带色彩输出的 fmt.Fprintln，颜色值被转换成 ANSI 码一起写入到 w 中。
func fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, sprint(foreground, background, v...))
}

// Fprintf 带色彩输出的 fmt.Fprintf，颜色值被转换成 ANSI 码一起写入到 w 中。
func fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprintf(foreground, background, format, v...))
}

// Print 带色彩输出的 fmt.Print，输出到 os.Stdout。
func print(foreground, background Color, v ...interface{}) (int, error) {
	return fprint(os.Stdout, foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println，输出到 os.Stdout。
func println(foreground, background Color, v ...interface{}) (int, error) {
	return fprintln(os.Stdout, foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf，输出到 os.Stdout。
func printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return fprintf(os.Stdout, foreground, background, format, v...)
}

// Sprint 带色彩输出的 fmt.Sprint，返回的字符，颜色值被转换成 ANSI 代码与字符中返回。
func sprint(foreground, background Color, v ...interface{}) string {
	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprint(v...) +
		ansi.Reset
}

// Sprintln 带色彩输出的 fmt.Sprintln，颜色值被转换成 ANSI 代码与字符中返回。
func sprintln(foreground, background Color, v ...interface{}) string {
	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprint(v...) +
		ansi.Reset +
		"\n"
}

// Sprintf 带色彩输出的 fmt.Sprintf，颜色值被转换成 ANSI 代码与字符中返回。
func sprintf(foreground, background Color, format string, v ...interface{}) string {
	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprintf(format, v...) +
		ansi.Reset
}
