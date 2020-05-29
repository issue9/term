// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"io"
	"os"

	"github.com/issue9/term/ansi"
)

// Fprint 带色彩输出的 fmt.Fprint
//
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprint(w io.Writer, t Type, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprint(w, sprint(!isConsole(w), t, foreground, background, v...))
}

// Fprintln 带色彩输出的 fmt.Fprintln
//
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprintln(w io.Writer, t Type, foreground, background Color, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, sprint(!isConsole(w), t, foreground, background, v...))
}

// Fprintf 带色彩输出的 fmt.Fprintf
//
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprintf(w io.Writer, t Type, foreground, background Color, format string, v ...interface{}) (int, error) {
	if !isConsole(w) {
		return fmt.Fprintf(w, format, v...)
	}

	if !isValidType(t) {
		panic("无效的参数 t")
	}

	if t == Normal {
		return fmt.Fprint(w, string(foreground.FColor())+string(background.BColor())+
			fmt.Sprint(v...)+
			string(ansi.CSI('m', ansi.ResetCode)))
	}

	return fmt.Fprint(w, string(ansi.CSI('m', int(t))+foreground.FColor())+string(background.BColor())+
		fmt.Sprintf(format, v...)+
		string(ansi.CSI('m', ansi.ResetCode)))
}

// Print 带色彩输出的 fmt.Print
func Print(t Type, foreground, background Color, v ...interface{}) (int, error) {
	return Fprint(os.Stdout, t, foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println
func Println(t Type, foreground, background Color, v ...interface{}) (int, error) {
	return Fprintln(os.Stdout, t, foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf
func Printf(t Type, foreground, background Color, format string, v ...interface{}) (int, error) {
	return Fprintf(os.Stdout, t, foreground, background, format, v...)
}

func sprint(ignoreAnsi bool, t Type, foreground, background Color, v ...interface{}) string {
	if ignoreAnsi {
		return fmt.Sprint(v...)
	}

	if !isValidType(t) {
		panic("无效的参数 t")
	}

	if t == Normal {
		return string(foreground.FColor()) + string(background.BColor()) +
			fmt.Sprint(v...) +
			string(ansi.CSI('m', ansi.ResetCode))
	}

	return string(ansi.CSI('m', int(t))+foreground.FColor()) + string(background.BColor()) +
		fmt.Sprint(v...) +
		string(ansi.CSI('m', ansi.ResetCode))
}

// 判断 w 是否为 stderr、stdout、stdin 三者之一
func isConsole(w io.Writer) bool {
	return w == os.Stderr || w == os.Stdout || w == os.Stdin
}
