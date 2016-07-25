// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build !windows

package colors

import (
	"io"
	"os"
)

// Fprint 带色彩输出的 fmt.Fprint，颜色值被转换成 ANSI 码一起写入到 w 中。
func Fprint(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fprint(w, foreground, background, v...)
}

// Fprintln 带色彩输出的 fmt.Fprintln，颜色值被转换成 ANSI 码一起写入到 w 中。
func Fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fprintln(w, foreground, background, v...)
}

// Fprintf 带色彩输出的 fmt.Fprintf，颜色值被转换成 ANSI 码一起写入到 w 中。
func Fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	return fprint(w, foreground, background, format, v...)
}

// Print 带色彩输出的 fmt.Print，输出到 os.Stdout。
func Print(foreground, background Color, v ...interface{}) (int, error) {
	return print(os.Stdout, foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println，输出到 os.Stdout。
func Println(foreground, background Color, v ...interface{}) (int, error) {
	return println(os.Stdout, foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf，输出到 os.Stdout。
func Printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return printf(os.Stdout, foreground, background, format, v...)
}

// Sprint 带色彩输出的 fmt.Sprint，返回的字符，颜色值被转换成 ANSI 代码与字符中返回。
func Sprint(foreground, background Color, v ...interface{}) string {
	return sprint(foreground, background, v...)
}

// Sprintln 带色彩输出的 fmt.Sprintln，颜色值被转换成 ANSI 代码与字符中返回。
func Sprintln(foreground, background Color, v ...interface{}) string {
	return sprintln(foreground, background, v...)
}

// Sprintf 带色彩输出的 fmt.Sprintf，颜色值被转换成 ANSI 代码与字符中返回。
func Sprintf(foreground, background Color, format string, v ...interface{}) string {
	return sprintf(foreground, background, format, v...)
}
