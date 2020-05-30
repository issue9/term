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
}

// New 新建一个 Colorize
func New(t Type, foreground, background Color) Colorize {
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
	}
}

// Print 等同于 fmt.Print()
func (c Colorize) Print(v ...interface{}) (int, error) {
	return fmt.Print(c.sgr + fmt.Sprint(v...))
}

// Println 等同于 fmt.Println()
func (c Colorize) Println(v ...interface{}) (int, error) {
	return fmt.Println(c.sgr + fmt.Sprint(v...))
}

// Printf 等同于 fmt.Printf()
func (c Colorize) Printf(format string, v ...interface{}) (int, error) {
	return fmt.Printf(c.sgr + fmt.Sprintf(format, v...))
}

// Fprint 等同于 fmt.Fprint()
//
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprint(w io.Writer, v ...interface{}) (int, error) {
	return fmt.Fprint(w, c.sgr+fmt.Sprint(v...))
}

// Fprintln 等同于 fmt.Fprintln()
//
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprintln(w io.Writer, v ...interface{}) (int, error) {
	return fmt.Fprintln(w, c.sgr+fmt.Sprint(v...))
}

// Fprintf 等同于 fmt.Fprintf()
//
// 若 w 不指向控制台，则颜色值被忽略。
func (c Colorize) Fprintf(w io.Writer, format string, v ...interface{}) (int, error) {
	return fmt.Fprintf(w, c.sgr+fmt.Sprintf(format, v...))
}
