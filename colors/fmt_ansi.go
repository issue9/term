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

// NOTE: fmt_windows 下的 Msys 模式会调用此文件下的函数。

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

// 判断 w 是否为 stderr、stdout、stdin 三者之一
func isConsole(out io.Writer) bool {
	o, ok := out.(*os.File)
	if !ok {
		return false
	}

	return o == os.Stdout || o == os.Stderr || o == os.Stdin
}

func fprint(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprint(!isConsole(w), foreground, background, v...))
}

func fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, sprint(!isConsole(w), foreground, background, v...))
}

func fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprintf(!isConsole(w), foreground, background, format, v...))
}

func print(foreground, background Color, v ...interface{}) (int, error) {
	return fprint(os.Stdout, foreground, background, v...)
}

func println(foreground, background Color, v ...interface{}) (int, error) {
	return fprintln(os.Stdout, foreground, background, v...)
}

func printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return fprintf(os.Stdout, foreground, background, format, v...)
}

func sprint(ignoreAnsi bool, foreground, background Color, v ...interface{}) string {
	if ignoreAnsi {
		return fmt.Sprint(v...)
	}

	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprint(v...) +
		ansi.Reset
}

func sprintln(ignoreAnsi bool, foreground, background Color, v ...interface{}) string {
	if ignoreAnsi {
		return fmt.Sprintln(v...)
	}

	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprint(v...) +
		ansi.Reset +
		"\n"
}

func sprintf(ignoreAnsi bool, foreground, background Color, format string, v ...interface{}) string {
	if ignoreAnsi {
		return fmt.Sprintf(format, v...)
	}

	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprintf(format, v...) +
		ansi.Reset
}
