// SPDX-License-Identifier: MIT

package colors

import "io"

// Colorize 指定了颜色的输出工具。
//
// 默认输出到 os.Stdout，若要指定其它，可以使用 Colorize.Fprintf 系列函数。
type Colorize struct {
	Foreground Color
	Background Color
}

// New 新建一个 Colorize
func New(foreground, background Color) Colorize {
	return Colorize{
		Foreground: foreground,
		Background: background,
	}
}

// Print 等同于 Print()，颜色由 Colorize 指定
func (c Colorize) Print(v ...interface{}) (int, error) {
	return Print(c.Foreground, c.Background, v...)
}

// Println 等同于 Println()，颜色由 Colorize 指定
func (c Colorize) Println(v ...interface{}) (int, error) {
	return Println(c.Foreground, c.Background, v...)
}

// Printf 等同于 Printf()，颜色由 Colorize 指定
func (c Colorize) Printf(format string, v ...interface{}) (int, error) {
	return Printf(c.Foreground, c.Background, format, v...)
}

// Fprint 等同于 Fprint()，颜色由 Colorize 指定，
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprint(w io.Writer, v ...interface{}) (int, error) {
	return Fprint(w, c.Foreground, c.Background, v...)
}

// Fprintln 等同于 Fprintln()，颜色由 Colorize 指定，
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprintln(w io.Writer, v ...interface{}) (int, error) {
	return Fprintln(w, c.Foreground, c.Background, v...)
}

// Fprintf 等同于 Fprintf()，颜色由 Colorize 指定，
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprintf(w io.Writer, format string, v ...interface{}) (int, error) {
	return Fprintf(w, c.Foreground, c.Background, format, v...)
}
