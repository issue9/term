// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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

	c.SetColor(Red, Black)
	_, err = c.Println("Colorize.Println:: foreground:", Red, ";background:", Black)
	a.NotError(err)

	c.SetColor(Black, Cyan)
	_, err = c.Printf("Colorize.Printf:: foreground:%v;background:%v\n", Black, Cyan)
	a.NotError(err)
}
