// Copyright 2019 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package prompt

import (
	"bytes"
	"io/ioutil"
	"testing"
	"testing/iotest"

	"github.com/issue9/term/colors"

	"github.com/issue9/assert"
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

func TestPrompt_read(t *testing.T) {
	a := assert.New(t)

	r := new(bytes.Buffer)
	p := New(0, r, ioutil.Discard, colors.Red)
	a.NotNil(p)

	r.WriteString("hello\nworld\n\n")
	a.Equal(p.read(), "hello")
	a.Equal(p.read(), "world")
	a.Equal(p.read(), "")
	a.Equal(p.read(), "")
	a.NotNil(p.err)

	// 没有读到指定分隔符，则读取所有
	r.Reset()
	p = New('x', r, ioutil.Discard, colors.Red)
	a.NotNil(p)
	r.WriteString("hello\nworld\n\n")
	a.Equal(p.read(), "hello\nworld\n\n").
		NotNil(p.err)

	// 返回错误信息
	r.Reset()
	p = New(0, iotest.TimeoutReader(r), ioutil.Discard, colors.Red)
	a.NotNil(p)
	r.WriteString("hello")
	a.Equal(p.read(), "hello")
	r.WriteString("world\n\n")
	a.Equal(p.read(), "").
		NotNil(p.err)
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
