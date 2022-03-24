// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/issue9/assert/v2"
)

func TestFprint(t *testing.T) {
	a := assert.New(t, false)
	path := "./fprint.test"

	f, err := os.Create(path)
	a.NotError(err).NotNil(f)
	_, err = Fprint(f, Bold, Red, Green, "abc")
	a.NotError(err)
	a.NotError(f.Close())

	data, err := ioutil.ReadFile(path)
	a.NotError(err).NotNil(data)
	a.Contains(string(data), "\033[") // 也包含控制符
}

func TestFprintf(t *testing.T) {
	a := assert.New(t, false)

	a.Panic(func() {
		Fprintf(os.Stderr, -100, Red, Green, "test")
	})

	Fprintf(os.Stderr, Normal, Red, Green, "test")

	Fprintf(os.Stderr, Blink, Red, Green, "test")
}

func TestFprintln(t *testing.T) {
	a := assert.New(t, false)

	a.Panic(func() {
		Fprintln(os.Stderr, -100, Red, Green, "test")
	})

	Fprintln(os.Stderr, Normal, Red, Green, "test")

	Fprintln(os.Stderr, Blink, Red, Green, "test")
}

func TestPrintf(t *testing.T) {
	a := assert.New(t, false)

	// named colors
	fmt.Printf("named colors\n")
	fmt.Printf("foreground:%s\n", Default)
	_, err := Printf(Italic, Default, Default, "%s\t", Default.String())
	a.NotError(err)

	for bColor := Black; bColor < maxNamedColor; bColor++ {
		_, err := Printf(Italic, Default, bColor, "%s\t", bColor.String())
		a.NotError(err)
	}
	fmt.Println()
	fmt.Println()

	for fColor := Black; fColor < maxNamedColor; fColor++ {
		fmt.Printf("foreground:%s\n", fColor)
		_, err := Printf(Italic, fColor, Default, "%s\t", Default.String())
		a.NotError(err)

		for bColor := Black; bColor < maxNamedColor; bColor++ {
			_, err := Printf(Italic, fColor, bColor, "%s\t", bColor.String())
			a.NotError(err)
		}
		fmt.Println()
		fmt.Println()
	}

	// 256
	fmt.Printf("\n\n256 colors\n")
	for i := maxNamedColor; i < end256Color; i++ {
		_, err := Fprintf(os.Stdout, Bold, Color(i), Default, "%d\t", i)
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
		_, err := Printf(Italic, rgb, Default, "%s\t", rgb)
		a.NotError(err)
	}
	fmt.Println()
}
