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

func fprint(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprint(foreground, background, v...))
}

func fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, sprint(foreground, background, v...))
}

func fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprintf(foreground, background, format, v...))
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

func sprint(foreground, background Color, v ...interface{}) string {
	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprint(v...) +
		ansi.Reset
}

func sprintln(foreground, background Color, v ...interface{}) string {
	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprint(v...) +
		ansi.Reset +
		"\n"
}

func sprintf(foreground, background Color, format string, v ...interface{}) string {
	return ansiForeTables[foreground] + ansiBackTables[background] +
		fmt.Sprintf(format, v...) +
		ansi.Reset
}
