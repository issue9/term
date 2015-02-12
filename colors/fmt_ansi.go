// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build !windows

package colors

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/issue9/term/ansi"
)

// 前景色对照表
var foreTables = [...]string{
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
var backTables = [...]string{
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

// 根据out获取对应的writer
func getW(out int) (io.Writer, error) {
	switch out {
	case Stderr:
		return os.Stderr, nil
	case Stdout:
		return os.Stdout, nil
	default:
		return nil, errors.New("getW:out值只能是Stderr或Stdout")
	}
}

// 功能同fmt.Print。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Print(out int, foreground, background Color, v ...interface{}) (size int, err error) {
	w, err := getW(out)
	if err != nil {
		return 0, err
	}

	f := foreTables[foreground] // 前景色
	b := backTables[background] // 背景色
	if size, err = fmt.Fprint(w, f, b); err != nil {
		return
	}
	if size, err = fmt.Fprint(w, v...); err != nil {
		return
	}
	return fmt.Fprint(w, ansi.Reset)
}

// 功能同fmt.Println。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Println(out int, foreground, background Color, v ...interface{}) (size int, err error) {
	w, err := getW(out)
	if err != nil {
		return 0, err
	}

	f := foreTables[foreground] // 前景色
	b := backTables[background] // 背景色
	if size, err = fmt.Fprint(w, f, b); err != nil {
		return
	}
	if size, err = fmt.Fprintln(w, v...); err != nil {
		return
	}
	return fmt.Fprint(w, ansi.Reset)
}

// 功能同fmt.Printf。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Printf(out int, foreground, background Color, format string, v ...interface{}) (size int, err error) {
	w, err := getW(out)
	if err != nil {
		return 0, err
	}

	f := foreTables[foreground] // 前景色
	b := backTables[background] // 背景色
	if size, err = fmt.Fprint(w, f, b); err != nil {
		return
	}
	if size, err = fmt.Fprintf(w, format, v...); err != nil {
		return
	}
	return fmt.Fprint(w, ansi.Reset)
}
