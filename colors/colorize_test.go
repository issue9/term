// SPDX-FileCopyrightText: 2014-2024 caixw
//
// SPDX-License-Identifier: MIT

package colors

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/issue9/assert/v4"
)

func TestColorize_Color(t *testing.T) {
	a := assert.New(t, false)

	buf := new(bytes.Buffer)
	c := New(buf)
	c.Color(Normal, Red, Blue)
	c.Println(buf, "test")
	a.Contains(buf.String(), "[31;44m") // 包含控制符
	a.NotError(c.Err())

	c = New(os.Stdout)
	a.Panic(func() {
		c.Color(-100, Red, Red)
	})

	// named colors
	fmt.Printf("named colors\n")
	fmt.Printf("foreground:%s\n", Default)
	c.Color(Italic, Default, Default)
	c.Printf("%s\t", Default.String())
	a.NotError(c.Err())

	for bColor := Black; bColor < maxNamedColor; bColor++ {
		c.Color(Italic, Default, bColor).WString(bColor.String()).WString("\t").Reset()
		a.NotError(c.Err())
	}
	fmt.Println()
	fmt.Println()

	for fColor := Black; fColor < maxNamedColor; fColor++ {
		fmt.Printf("foreground:%s\n", fColor)
		c.Color(Italic, fColor, Default).Printf("%s\t", Default.String()).Reset()
		a.NotError(c.Err())

		for bColor := Black; bColor < maxNamedColor; bColor++ {
			c.Color(Italic, fColor, bColor).Printf("%s\t", bColor.String()).Reset()
			a.NotError(c.Err())
		}
		fmt.Println()
		fmt.Println()
	}

	// 256
	fmt.Printf("\n\n256 colors\n")
	for i := maxNamedColor; i < end256Color; i++ {
		c.Color(Bold, i, Default).Printf("%d\t", i).Reset()
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
		c.Color(Italic, rgb, Default).Printf("%s\t", rgb).Reset()
		a.NotError(c.Err())
	}
	fmt.Println()
}
