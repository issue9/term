// Copyright 2019 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package prompt

import (
	"bytes"
	"io/ioutil"
	"testing"
	"testing/iotest"

	"github.com/issue9/assert"

	"github.com/issue9/term/colors"
)

func TestW_print(t *testing.T) {
	a := assert.New(t)

	r := new(bytes.Buffer)
	w := &w{}
	p := New(0, r, ioutil.Discard, colors.Red)
	a.NotNil(p)

	w.print(r, colors.Default, "print")
	a.Equal(r.String(), "print")

	r.Reset()
	w.println(r, colors.Default, "println")
	a.Equal(r.String(), "println\n")

	r.Reset()
	w.printf(r, colors.Default, "printf %s", "printf")
	a.Equal(r.String(), "printf printf")
}

func TestW_read(t *testing.T) {
	a := assert.New(t)

	r := new(bytes.Buffer)
	w := &w{}
	p := New(0, r, ioutil.Discard, colors.Red)
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
	p = New('x', r, ioutil.Discard, colors.Red)
	a.NotNil(p)
	r.WriteString("hello\nworld\n\n")
	a.Equal(w.read(p), "").
		NotNil(w.err)

	// 返回错误信息
	r.Reset()
	w.err = nil
	p = New(0, iotest.TimeoutReader(r), ioutil.Discard, colors.Red)
	a.NotNil(p)
	r.WriteString("hello")
	a.Equal(w.read(p), "").
		NotNil(w.err)
	r.WriteString("world\n\n")
	a.Equal(w.read(p), "").
		NotNil(w.err)
}
