// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"io"
	"os"

	"github.com/issue9/term/v3/ansi"
)

// Fprint 带色彩输出的 [fmt.Fprint]
func Fprint(w io.Writer, t Type, foreground, background Color, v ...any) (int, error) {
	return fmt.Fprint(w, Sprint(t, foreground, background, v...))
}

// Fprintln 带色彩输出的 [fmt.Fprintln]
func Fprintln(w io.Writer, t Type, foreground, background Color, v ...any) (int, error) {
	return fmt.Fprint(w, Sprintln(t, foreground, background, v...))
}

// Fprintf 带色彩输出的 [fmt.Fprintf]
func Fprintf(w io.Writer, t Type, foreground, background Color, format string, v ...any) (int, error) {
	return fmt.Fprint(w, Sprintf(t, foreground, background, format, v...))
}

// Print 带色彩输出的 [fmt.Print]
func Print(t Type, foreground, background Color, v ...any) (int, error) {
	return Fprint(os.Stdout, t, foreground, background, v...)
}

// Println 带色彩输出的 [fmt.Println]
func Println(t Type, foreground, background Color, v ...any) (int, error) {
	return Fprintln(os.Stdout, t, foreground, background, v...)
}

// Printf 带色彩输出的 [fmt.Printf]
func Printf(t Type, foreground, background Color, format string, v ...any) (int, error) {
	return Fprintf(os.Stdout, t, foreground, background, format, v...)
}

func Sprint(t Type, foreground, background Color, v ...any) string {
	if !isValidType(t) {
		panic("无效的参数 t")
	}

	return string(sgr(t, foreground, background)) + fmt.Sprint(v...) + string(ansi.SGR())
}

func Sprintln(t Type, foreground, background Color, v ...any) string {
	if !isValidType(t) {
		panic("无效的参数 t")
	}

	return string(sgr(t, foreground, background)) + fmt.Sprint(v...) + string(ansi.SGR()) + "\n"
}

func Sprintf(t Type, foreground, background Color, format string, v ...any) string {
	if !isValidType(t) {
		panic("无效的参数 t")
	}

	color := string(sgr(t, foreground, background))
	return fmt.Sprintf(color + fmt.Sprintf(format, v...) + string(ansi.SGR()))
}
