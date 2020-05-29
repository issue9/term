// SPDX-License-Identifier: MIT

package prompt

import (
	"io"

	"github.com/issue9/term/v2/colors"
)

type w struct {
	err error
}

func (w *w) println(output io.Writer, c colors.Color, v ...interface{}) {
	if w.err == nil {
		_, w.err = colors.Fprintln(output, colors.Normal, c, colors.Default, v...)
	}
}

func (w *w) print(output io.Writer, c colors.Color, v ...interface{}) {
	if w.err == nil {
		_, w.err = colors.Fprint(output, colors.Normal, c, colors.Default, v...)
	}
}

func (w *w) printf(output io.Writer, c colors.Color, format string, v ...interface{}) {
	if w.err == nil {
		_, w.err = colors.Fprintf(output, colors.Normal, c, colors.Default, format, v...)
	}
}

// 从输入端读取一行内容
func (w *w) read(p *Prompt) (v string) {
	if w.err != nil {
		return ""
	}

	v, w.err = p.reader.ReadString(p.delim)
	if w.err != nil {
		return ""
	}

	return v[:len(v)-1]
}
