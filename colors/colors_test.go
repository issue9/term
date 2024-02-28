// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

package colors

import (
	"testing"

	"github.com/issue9/assert/v4"
)

func TestHEX(t *testing.T) {
	a := assert.New(t, false)

	r, g, b := HEX("#111").RGB()
	a.Equal(r, 0x11).
		Equal(g, 0x11).
		Equal(b, 0x11)

	r, g, b = HEX("#cc00cc").RGB()
	a.Equal(r, 0xcc).
		Equal(g, 0x00).
		Equal(b, 0xcc)

	r, g, b = HEX("cc00cc").RGB()
	a.Equal(r, 0xcc).
		Equal(g, 0x00).
		Equal(b, 0xcc)

	a.Panic(func() {
		HEX("c0")
	})

	a.Panic(func() {
		HEX("xxx")
	})

	a.Panic(func() {
		HEX("")
	})
}

func TestColor_String(t *testing.T) {
	a := assert.New(t, false)

	c := Red
	a.Equal(c.String(), "Red")

	c = Default
	a.Equal(c.String(), "Default")

	c = 100
	a.Equal(c.String(), "100")

	c = 255
	a.Equal(c.String(), "255")

	c = RGB(100, 100, 100)
	a.Equal(c.String(), "(100,100,100)")
}

func TestColor_RGB(t *testing.T) {
	a := assert.New(t, false)

	c := Red
	a.Panic(func() {
		c.RGB()
	})

	r, g, b := uint8(100), uint8(200), uint8(95)
	c = RGB(r, g, b)
	r1, g1, b1 := c.RGB()
	a.Equal(r1, r).Equal(g1, g).Equal(b1, b)
}
