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
	return fmt.Fprint(w, Sprint(foreground, background, v...))
}

// Fprintln 带色彩输出的 fmt.Fprintln，颜色值被转换成 ANSI 码一起写入到 w 中。
func fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, Sprint(foreground, background, v...))
}

// Fprintf 带色彩输出的 fmt.Fprintf，颜色值被转换成 ANSI 码一起写入到 w 中。
func fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	return fmt.Fprint(w, Sprintf(foreground, background, format, v...))
}

// Print 带色彩输出的 fmt.Print，输出到 os.Stdout。
func print(foreground, background Color, v ...interface{}) (int, error) {
	return Fprint(os.Stdout, foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println，输出到 os.Stdout。
func println(foreground, background Color, v ...interface{}) (int, error) {
	return Fprintln(os.Stdout, foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf，输出到 os.Stdout。
func printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return Fprintf(os.Stdout, foreground, background, format, v...)
}

// Sprint 带色彩输出的 fmt.Sprint，返回的字符，颜色值被转换成 ANSI 代码与字符中返回。
func sprint(foreground, background Color, v ...interface{}) string {
	buf := fmt.Sprint(ansiForeTables[foreground], ansiBackTables[background])
	buf += fmt.Sprint(v...)
	return buf + fmt.Sprint(ansi.Reset)
}

// Sprintln 带色彩输出的 fmt.Sprintln，颜色值被转换成 ANSI 代码与字符中返回。
func sprintln(foreground, background Color, v ...interface{}) string {
	buf := fmt.Sprint(ansiForeTables[foreground], ansiBackTables[background])
	buf += fmt.Sprint(v...)
	return buf + fmt.Sprintln(ansi.Reset)
}

// Sprintf 带色彩输出的 fmt.Sprintf，颜色值被转换成 ANSI 代码与字符中返回。
func sprintf(foreground, background Color, format string, v ...interface{}) string {
	buf := fmt.Sprint(ansiForeTables[foreground], ansiBackTables[background])
	buf += fmt.Sprintf(format, v...)
	return buf + fmt.Sprint(ansi.Reset)
}
