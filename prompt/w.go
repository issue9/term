// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

package prompt

import (
	"bufio"
	"io"

	"github.com/issue9/term/v3/colors"
)

type w struct {
	output io.Writer
	err    error
}

func (w *w) println(c colors.Color, v ...any) {
	if w.err == nil {
		_, w.err = colors.Fprintln(w.output, colors.Normal, c, colors.Default, v...)
	}
}

func (w *w) print(c colors.Color, v ...any) {
	if w.err == nil {
		_, w.err = colors.Fprint(w.output, colors.Normal, c, colors.Default, v...)
	}
}

func (w *w) printf(c colors.Color, format string, v ...any) {
	if w.err == nil {
		_, w.err = colors.Fprintf(w.output, colors.Normal, c, colors.Default, format, v...)
	}
}

// 从输入端读取一行内容
func (w *w) read(reader *bufio.Reader, delim byte) (v string) {
	if w.err != nil {
		return ""
	}

	v, w.err = reader.ReadString(delim)
	if w.err != nil {
		return ""
	}

	return v[:len(v)-1]
}
