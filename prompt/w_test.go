// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

package prompt

import (
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
	w := &w{}
	p := New(0, r, io.Discard, colors.Red)
	a.NotNil(p)

	w.print(r, colors.Default, "print")
	a.Contains(r.String(), "print")

	r.Reset()
	w.println(r, colors.Default, "println")
	a.Contains(r.String(), "println")

	r.Reset()
	w.printf(r, colors.Default, "printf %s", "printf")
	a.Contains(r.String(), "printf printf")
}

func TestW_read(t *testing.T) {
	a := assert.New(t, false)

	r := new(bytes.Buffer)
	w := &w{}
	p := New(0, r, io.Discard, colors.Red)
	a.NotNil(p)

	r.WriteString("hello\nworld\n\n")
	a.Equal(w.read(p), "hello")
	a.Equal(w.read(p), "world")
	a.Equal(w.read(p), "")
	a.Equal(w.read(p), "")
	a.NotNil(w.err)

	// 没有读到指定分隔符，则读取所有
	r.Reset()
	w.err = nil
	p = New('x', r, io.Discard, colors.Red)
	a.NotNil(p)
	r.WriteString("hello\nworld\n\n")
	a.Equal(w.read(p), "").
		NotNil(w.err)

	// 返回错误信息
	r.Reset()
	w.err = nil
	p = New(0, iotest.TimeoutReader(r), io.Discard, colors.Red)
	a.NotNil(p)
	r.WriteString("hello")
	a.Equal(w.read(p), "").
		NotNil(w.err)
	r.WriteString("world\n\n")
	a.Equal(w.read(p), "").
		NotNil(w.err)
}
