// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build !windows

package colors

import "io"

// Fprint 带色彩输出的 fmt.Fprint。
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprint(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fprint(w, foreground, background, v...)
}

// Fprintln 带色彩输出的 fmt.Fprintln。
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprintln(w io.Writer, foreground, background Color, v ...interface{}) (int, error) {
	return fprintln(w, foreground, background, v...)
}

// Fprintf 带色彩输出的 fmt.Fprintf。
// 颜色值只在 w 不为 os.Stderr、os.Stdin、os.Stdout 中的一个时才启作用，否则只向 w 输出普通字符串。
func Fprintf(w io.Writer, foreground, background Color, format string, v ...interface{}) (int, error) {
	return fprintf(w, foreground, background, format, v...)
}

// Print 带色彩输出的 fmt.Print，输出到 os.Stdout。
func Print(foreground, background Color, v ...interface{}) (int, error) {
	return print(foreground, background, v...)
}

// Println 带色彩输出的 fmt.Println，输出到 os.Stdout。
func Println(foreground, background Color, v ...interface{}) (int, error) {
	return println(foreground, background, v...)
}

// Printf 带色彩输出的 fmt.Printf，输出到 os.Stdout。
func Printf(foreground, background Color, format string, v ...interface{}) (int, error) {
	return printf(foreground, background, format, v...)
}
