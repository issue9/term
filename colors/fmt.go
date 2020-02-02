// SPDX-License-Identifier: MIT

// +build !windows

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

// Fprint 带色彩输出的 fmt.Fprint。
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprint(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprint(!isConsole(w), foreground, background, v...))
}

// Fprintln 带色彩输出的 fmt.Fprintln。
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, sprint(!isConsole(w), foreground, background, v...))
}

// Fprintf 带色彩输出的 fmt.Fprintf。
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	if !isConsole(w) {
		return fmt.Fprintf(w, format, v...)
	}

	return fmt.Fprint(w, ansiForeTables[foreground]+ansiBackTables[background]+
		fmt.Sprintf(format, v...)+
		ansi.Reset)
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

func sprint(ignoreAnsi bool, foreground, background Color, v ...interface{}) string {
	if ignoreAnsi {
		return fmt.Sprint(v...)
	}

	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprint(v...) +
		ansi.Reset
}

// 判断 w 是否为 stderr、stdout、stdin 三者之一
func isConsole(out io.Writer) bool {
	return out == os.Stdout || out == os.Stderr || out == os.Stdin
}
