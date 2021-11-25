// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"io"

	"github.com/issue9/term/v2/ansi"
)

// Colorize 适合固定颜色的大段内容输出
type Colorize struct {
	Type       Type
	Foreground Color
	Background Color
	sgr        string
	reset      string
	out        io.Writer
}

// New 新建一个 Colorize
func New(out io.Writer, t Type, foreground, background Color) Colorize {
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
		out:        out,
	}
}

func (c Colorize) Print(v ...interface{}) (int, error) {
	return fmt.Fprint(c.out, c.sgr, fmt.Sprint(v...), c.reset)
}

func (c Colorize) Println(v ...interface{}) (int, error) {
	return fmt.Fprintln(c.out, c.sgr, fmt.Sprint(v...), c.reset)
}

func (c Colorize) Printf(format string, v ...interface{}) (int, error) {
	return fmt.Fprint(c.out, c.sgr, fmt.Sprintf(format, v...), c.reset)
}
