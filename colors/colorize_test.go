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
		New(os.Stdout, -100, Red, Red)
	})

	buf := new(bytes.Buffer)
	c := New(buf, Normal, Red, Blue)
	c.Println(buf, "test")
	a.Contains(buf.String(), "[31;44m") // 包含控制符
	a.NotError(c.Err())

	// named colors
	fmt.Printf("named colors\n")
	fmt.Printf("foreground:%s\n", Default)
	c = New(os.Stdout, Italic, Default, Default)
	c.Printf("%s\t", Default.String())
	a.NotError(c.Err())

	for bColor := Black; bColor < maxNamedColor; bColor++ {
		c := New(os.Stdout, Italic, Default, bColor)
		c.Printf("%s\t", bColor.String())
		a.NotError(c.Err())
	}
	fmt.Println()
	fmt.Println()

	for fColor := Black; fColor < maxNamedColor; fColor++ {
		fmt.Printf("foreground:%s\n", fColor)
		c := New(os.Stdout, Italic, fColor, Default)
		c.Printf("%s\t", Default.String())
		a.NotError(c.Err())

		for bColor := Black; bColor < maxNamedColor; bColor++ {
			c := New(os.Stdout, Italic, fColor, bColor)
			c.Printf("%s\t", bColor.String())
			a.NotError(c.Err())
		}
		fmt.Println()
		fmt.Println()
	}

	// 256
	fmt.Printf("\n\n256 colors\n")
	for i := maxNamedColor; i < end256Color; i++ {
		c := New(os.Stdout, Bold, i, Default).Printf("%d\t", i)
		a.NotError(c.Err())
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
		c := New(os.Stdout, Italic, rgb, Default).Printf("%s\t", rgb)
		a.NotError(c.Err())
	}
	fmt.Println()
}
