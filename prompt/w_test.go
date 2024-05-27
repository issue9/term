// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

package prompt

import (
	"bufio"
	"bytes"
	"io"
	"testing"
	"testing/iotest"

	"github.com/issue9/assert/v4"

	"github.com/issue9/term/v3/colors"
)

func TestW_print(t *testing.T) {
	a := assert.New(t, false)

	r := new(bytes.Buffer)
	w := &w{output: r}
	p := New(0, r, io.Discard, colors.Red)
	a.NotNil(p)

	w.print(colors.Default, "print")
	a.Contains(r.String(), "print")

	r.Reset()
	w.println(colors.Default, "println")
	a.Contains(r.String(), "println")

	r.Reset()
	w.printf(colors.Default, "printf %s", "printf")
	a.Contains(r.String(), "printf printf")
}

func TestW_read(t *testing.T) {
	a := assert.New(t, false)

	r := new(bytes.Buffer)
	rr := bufio.NewReader(r)
	w := &w{}

	r.WriteString("hello\nworld\n\n")
	a.Equal(w.read(rr, '\n'), "hello")
	a.Equal(w.read(rr, '\n'), "world")
	a.Equal(w.read(rr, '\n'), "")
	a.Equal(w.read(rr, '\n'), "")
	a.NotNil(w.err)

	// 没有读到指定分隔符，则读取所有
	rr.Reset(r)
	w.err = nil
	r.WriteString("hello\nworld\n\n")
	a.Equal(w.read(rr, 'x'), "").
		NotNil(w.err)

	// 返回错误信息
	rr.Reset(r)
	w.err = nil
	rr = bufio.NewReader(iotest.TimeoutReader(r))
	r.WriteString("hello")
	a.Equal(w.read(rr, '\n'), "").
		NotNil(w.err)
	r.WriteString("world\n\n")
	a.Equal(w.read(rr, '\n'), "").
		NotNil(w.err)
}
