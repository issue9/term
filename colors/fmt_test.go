// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package colors

import (
	"testing"

	"github.com/issue9/assert"
)

func TestPrint(t *testing.T) {
	a := assert.New(t)

	_, err := Print(Stderr, Red, Blue, "Print::foreground:Red;background:Blue\n")
	a.NotError(err)

	_, err = Println(Stdout, Cyan, Default, "Println::foreground:Cyan;background:Default")
	a.NotError(err)

	_, err = Printf(Stdout, Red, Blue, "Print:foreground:%v;background:%v\n", Red, Blue)
	a.NotError(err)
}
