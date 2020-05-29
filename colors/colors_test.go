// SPDX-License-Identifier: MIT

package colors

import (
	"testing"

	"github.com/issue9/assert"

	"github.com/issue9/term/v2/ansi"
)

func TestColor_String(t *testing.T) {
	a := assert.New(t)

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

func TestColor_FColor(t *testing.T) {
	a := assert.New(t)

	c := Red
	a.Equal(c.FColor(), ansi.ESC("\033[31m"))

	c = Default
	a.Equal(c.FColor(), ansi.ESC("\033[39m"))

	c = BrightRed
	a.Equal(c.FColor(), ansi.ESC("\033[91m"))

	c = 255
	a.Equal(c.FColor(), ansi.ESC("\033[38;5;255m"))

	c = RGB(100, 200, 100)
	a.Equal(c.FColor(), ansi.ESC("\033[38;2;100;200;100m"))
}

func TestColor_BColor(t *testing.T) {
	a := assert.New(t)

	c := Red
	a.Equal(c.BColor(), ansi.ESC("\033[41m"))

	c = Default
	a.Equal(c.BColor(), ansi.ESC("\033[49m"))

	c = BrightRed
	a.Equal(c.BColor(), ansi.ESC("\033[101m"))

	c = 255
	a.Equal(c.BColor(), ansi.ESC("\033[48;5;255m"))

	c = RGB(100, 200, 100)
	a.Equal(c.BColor(), ansi.ESC("\033[48;2;100;200;100m"))
}

func TestColor_RGB(t *testing.T) {
	a := assert.New(t)

	c := Red
	a.Panic(func() {
		c.RGB()
	})

	r, g, b := uint8(100), uint8(200), uint8(95)
	c = RGB(r, g, b)
	r1, g1, b1 := c.RGB()
	a.Equal(r1, r).Equal(g1, g).Equal(b1, b)
}
