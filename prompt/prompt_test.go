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

func TestNew(t *testing.T) {
	a := assert.New(t)

	p := New(0, new(bytes.Buffer), ioutil.Discard, colors.Red)
	a.NotNil(p)
	a.Equal(p.delim, '\n').
		Equal(p.defaultColor, colors.Red)

	p = New('x', new(bytes.Buffer), ioutil.Discard, colors.Red)
	a.NotNil(p)
	a.Equal(p.delim, 'x')

	a.Panic(func() {
		New(0, new(bytes.Buffer), ioutil.Discard, colors.Color(123))
	})
}

func TestPrompt_String(t *testing.T) {
	a := assert.New(t)

	r := new(bytes.Buffer)
	w := new(bytes.Buffer)
	p := New(0, r, w, colors.Red)
	a.NotNil(p)

	r.WriteString("v1\n")
	v, err := p.String("string", "def")
	a.NotError(err)
	a.Equal(w.String(), "string（def）：")
	a.Equal(v, "v1")
}

func TestPrompt_Bool(t *testing.T) {
	a := assert.New(t)

	r := new(bytes.Buffer)
	w := new(bytes.Buffer)
	p := New(0, r, w, colors.Red)
	a.NotNil(p)

	r.WriteString("Y\n")
	v, err := p.Bool("string", true)
	a.NotError(err)
	a.Equal(w.String(), "string（Y）：")
	a.Equal(v, true)
}

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

func TestInIntSlice(t *testing.T) {
	a := assert.New(t)

	vals := []int{1, 2, 3, 4, 5}
	a.True(inIntSlice(5, vals))
	a.True(inIntSlice(1, vals))
	a.True(inIntSlice(3, vals))
	a.False(inIntSlice(9, vals))
}

func TestInStringSlice(t *testing.T) {
	a := assert.New(t)

	vals := []string{"1", "2", "3", "4", "5"}
	a.True(inStringSlice("5", vals))
	a.True(inStringSlice("1", vals))
	a.True(inStringSlice("3", vals))
	a.False(inStringSlice("9", vals))
}
