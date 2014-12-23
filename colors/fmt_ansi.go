// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build !windows

package colors

import (
	"errors"
	"fmt"
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

// 功能同fmt.Print。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Print(out int, foreground, background Color, v ...interface{}) (size int, err error) {
	switch out {
	case Stderr:
		fmt.Fprint(os.Stderr, foreTables[foreground], backTables[background])
		size, err = fmt.Fprint(os.Stderr, v...)
		fmt.Fprint(os.Stderr, foreTables[Default], backTables[Default])
	case Stdout:
		fmt.Fprint(os.Stdout, foreTables[foreground], backTables[background])
		size, err = fmt.Fprint(os.Stdout, v...)
		fmt.Fprint(os.Stdout, foreTables[Default], backTables[Default])
	default:
		return 0, errors.New("out值只能是Stderr或Stdout")
	}
	return
}

// 功能同fmt.Println。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Println(out int, foreground, background Color, v ...interface{}) (size int, err error) {
	switch out {
	case Stderr:
		fmt.Fprint(os.Stderr, foreTables[foreground], backTables[background])
		size, err = fmt.Fprintln(os.Stderr, v...)
		fmt.Fprint(os.Stderr, foreTables[Default], backTables[Default])
	case Stdout:
		fmt.Fprint(os.Stdout, foreTables[foreground], backTables[background])
		size, err = fmt.Fprintln(os.Stdout, v...)
		fmt.Fprint(os.Stdout, foreTables[Default], backTables[Default])
	default:
		return 0, errors.New("out值只能是Stderr或Stdout")
	}
	return
}

// 功能同fmt.Printf。但是输出源可以通过out指定为Stderr或是Stdout。
// foreground，background为输出文字的前景和背景色。
func Printf(out int, foreground, background Color, format string, v ...interface{}) (size int, err error) {
	switch out {
	case Stderr:
		fmt.Fprint(os.Stderr, foreTables[foreground], backTables[background])
		size, err = fmt.Fprintf(os.Stderr, format, v...)
		fmt.Fprint(os.Stderr, foreTables[Default], backTables[Default])
	case Stdout:
		fmt.Fprintln(os.Stdout, foreTables[foreground], backTables[background])
		size, err = fmt.Fprintf(os.Stdout, format, v...)
		fmt.Fprint(os.Stdout, foreTables[Default], backTables[Default])
	default:
		return 0, errors.New("out值只能是Stderr或Stdout")
	}
	return
}
