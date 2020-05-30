// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"io"

	"github.com/issue9/term/v2/ansi"
)

// Colorize 指定了颜色的输出工具
//
// 默认输出到 os.Stdout，若要指定其它，可以使用 Colorize.Fprintf 系列函数。
type Colorize struct {
	Type       Type
	Foreground Color
	Background Color
	sgr        string
	reset      string
}

// New 新建一个 Colorize
func New(t Type, foreground, background Color) Colorize {
	if !isValidType(t) {
		panic("无效的参数 t")
	}

	codes := make([]int, 0, 10)
	if t != Normal {
		codes = append(codes, int(t))
	}
	codes = append(codes, foreground.fColorCode()...)
	codes = append(codes, background.bColorCode()...)

	return Colorize{
		Type:       t,
		Foreground: foreground,
		Background: background,
		sgr:        string(ansi.SGR(codes...)),
		reset:      string(ansi.SGR(ansi.ResetCode)),
	}
}

// Print 等同于 fmt.Print()
func (c Colorize) Print(v ...interface{}) (int, error) {
	return fmt.Print(c.sgr, fmt.Sprint(v...), c.reset)
}

// Println 等同于 fmt.Println()
func (c Colorize) Println(v ...interface{}) (int, error) {
	return fmt.Println(c.sgr, fmt.Sprint(v...), c.reset)
}

// Printf 等同于 fmt.Printf()
func (c Colorize) Printf(format string, v ...interface{}) (int, error) {
	return fmt.Print(c.sgr, fmt.Sprintf(format, v...), c.reset)
}

// Fprint 等同于 fmt.Fprint()
//
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprint(w io.Writer, v ...interface{}) (int, error) {
	if !isConsole(w) {
		return fmt.Fprint(w, v...)
	}
	return fmt.Fprint(w, c.sgr, fmt.Sprint(v...), c.reset)
}

// Fprintln 等同于 fmt.Fprintln()
//
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprintln(w io.Writer, v ...interface{}) (int, error) {
	if !isConsole(w) {
		return fmt.Fprintln(w, v...)
	}
	return fmt.Fprintln(w, c.sgr, fmt.Sprint(v...), c.reset)
}

// Fprintf 等同于 fmt.Fprintf()
//
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprintf(w io.Writer, format string, v ...interface{}) (int, error) {
	if !isConsole(w) {
		return fmt.Fprintf(w, format, v...)
	}
	return fmt.Fprint(w, c.sgr, fmt.Sprintf(format, v...), c.reset)
}
