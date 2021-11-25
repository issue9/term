// SPDX-License-Identifier: MIT

package colors

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/issue9/assert/v2"
)

func TestColorize(t *testing.T) {
	a := assert.New(t, false)

	a.Panic(func() {
		New(-100, Red, Red)
	})

	buf := new(bytes.Buffer)
	c := New(Normal, Red, Blue)
	_, err := c.Fprintln(buf, "test")
	a.NotError(err).
		Contains(buf.String(), "[31;44m") // 包含控制符

	// named colors
	fmt.Printf("named colors\n")
	fmt.Printf("foreground:%s\n", Default)
	c = New(Italic, Default, Default)
	_, err = c.Printf("%s\t", Default.String())
	a.NotError(err)

	for bColor := Black; bColor < maxNamedColor; bColor++ {
		c := New(Italic, Default, bColor)
		_, err := c.Printf("%s\t", bColor.String())
		a.NotError(err)
	}
	fmt.Println()
	fmt.Println()

	for fColor := Black; fColor < maxNamedColor; fColor++ {
		fmt.Printf("foreground:%s\n", fColor)
		c := New(Italic, fColor, Default)
		_, err := c.Printf("%s\t", Default.String())
		a.NotError(err)

		for bColor := Black; bColor < maxNamedColor; bColor++ {
			c := New(Italic, fColor, bColor)
			_, err := c.Printf("%s\t", bColor.String())
			a.NotError(err)
		}
		fmt.Println()
		fmt.Println()
	}

	// 256
	fmt.Printf("\n\n256 colors\n")
	for i := maxNamedColor; i < end256Color; i++ {
		_, err := (New(Bold, i, Default)).Fprintf(os.Stdout, "%d\t", i)
		a.NotError(err)
	}
	fmt.Println()

	// RGB
	fmt.Printf("\n\nRGB colors\n")
	var b int
	r := end256Color - 1
	for i := 0; i < end256Color; i++ {
		b = i + 5
		if b >= end256Color {
			b = 1
		}

		r -= 5
		if r <= 0 {
			r = end256Color - 1
		}

		rgb := RGB(uint8(r), uint8(end256Color-i), uint8(b))
		_, err := (New(Italic, rgb, Default)).Printf("%s\t", rgb)
		a.NotError(err)
	}
	fmt.Println()
}
