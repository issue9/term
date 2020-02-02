// SPDX-License-Identifier: MIT

package colors

import (
	"testing"

	"github.com/issue9/assert"
)

func TestColorize(t *testing.T) {
	a := assert.New(t)

	c := New(Green, White)
	_, err := c.Print("Colorize.Print:: foreground:", Green, ";background:", White, "\n")
	a.NotError(err)

	c.Background = Red
	c.Background = Black
	_, err = c.Println("Colorize.Println:: foreground:", Red, ";background:", Black)
	a.NotError(err)

	c.Background = Black
	c.Background = Cyan
	_, err = c.Printf("Colorize.Printf:: foreground:%v;background:%v\n", Black, Cyan)
	a.NotError(err)
}
