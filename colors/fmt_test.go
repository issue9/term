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

	_, err := Printf(Stdout, Red, Blue, "Print::foreground:%v;background:%v", Red, Blue)
	a.NotError(err)
	_, err = Printf(Stdout, Blue, Red, "Print::foreground:%v;background:%v\n", Blue, Red)
	a.NotError(err)

	_, err = Println(Stdout, Cyan, Default, "Println::foreground:Cyan;background:Default")
	a.NotError(err)

	_, err = Print(Stderr, Red, Blue, "Print::foreground:Red;background:Blue\n\n")
	a.NotError(err)
}
