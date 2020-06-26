// SPDX-License-Identifier: MIT

package prompt

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/issue9/assert"

	"github.com/issue9/term/v2/colors"
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

func TestInIntSlice(t *testing.T) {
	a := assert.New(t)

	vals := []int{1, 2, 3, 4, 5}
	a.True(inIntSlice(5, vals))
	a.True(inIntSlice(1, vals))
	a.True(inIntSlice(3, vals))
	a.False(inIntSlice(9, vals))
}
