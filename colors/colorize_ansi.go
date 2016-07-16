// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build !windows

package colors

import "io"

type Colorize struct {
	foreground Color
	background Color
}

// 新建一个 Colorize
func New(foreground, background Color) Colorize {
	return Colorize{
		foreground: foreground,
		background: background,
	}
}

// SetColor 更改颜色值。foreground 文字颜色；background 背景色。
func (c Colorize) SetColor(foreground, background Color) {
	c.foreground = foreground
	c.background = background
}

// Print 等同于 Print()，颜色由 Colorize 指定
func (c Colorize) Print(v ...interface{}) (int, error) {
	return Print(c.foreground, c.background, v...)
}

// Println 等同于 Println()，颜色由 Colorize 指定
func (c Colorize) Println(v ...interface{}) (int, error) {
	return Println(c.foreground, c.background, v...)
}

// Printf 等同于 Printf()，颜色由 Colorize 指定
func (c Colorize) Printf(format string, v ...interface{}) (int, error) {
	return Printf(c.foreground, c.background, format, v...)
}

// Fprint 等同于 Fprint()，颜色由 Colorize 指定，
// 若 w 不指赂控制台，则颜色值以 ansi 值的形式出现在字符串中。
func (c Colorize) Fprint(w io.Writer, v ...interface{}) (int, error) {
	return Fprint(w, c.foreground, c.background, v...)
}

// Fprintln 等同于 Fprintln()，颜色由 Colorize 指定，
// 若 w 不指赂控制台，则颜色值以 ansi 值的形式出现在字符串中。
func (c Colorize) Fprintln(w io.Writer, v ...interface{}) (int, error) {
	return Fprintln(w, c.foreground, c.background, v...)
}

// Fprintf 等同于 Fprintf()，颜色由 Colorize 指定，
// 若 w 不指赂控制台，则颜色值以 ansi 值的形式出现在字符串中。
func (c Colorize) Fprintf(w io.Writer, format string, v ...interface{}) (int, error) {
	return Fprintf(w, c.foreground, c.background, format, v...)
}

// Sprint 等同于 Sprint()，颜色由 Colorize 指定，
// 颜色值以 ansi 值的形式出现在字符串中。
func (c Colorize) Sprint(w io.Writer, v ...interface{}) string {
	return Sprint(c.foreground, c.background, v...)
}

// Sprintln 等同于 Sprintln()，颜色由 Colorize 指定，
// 颜色值以 ansi 值的形式出现在字符串中。
func (c Colorize) Sprintln(w io.Writer, v ...interface{}) string {
	return Sprintln(c.foreground, c.background, v...)
}

// Sprintf 等同于 Sprintf()，颜色由 Colorize 指定，
// 颜色值以 ansi 值的形式出现在字符串中。
func (c Colorize) Sprintf(w io.Writer, format string, v ...interface{}) string {
	return Sprintf(c.foreground, c.background, format, v...)
}
