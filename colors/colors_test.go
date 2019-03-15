// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package colors

import (
	"bytes"
	"os"
	"testing"

	"github.com/issue9/assert"
)

func TestColor(t *testing.T) {
	a := assert.New(t)

	c := Red
	a.True(c.IsValid()).
		Equal(c.String(), "Red")

	c = 100
	a.False(c.IsValid()).
		Equal(c.String(), "<unknown>")
}

func TestIsConsole(t *testing.T) {
	a := assert.New(t)

	a.True(isConsole(os.Stderr))
	a.True(isConsole(os.Stdout))
	a.True(isConsole(os.Stdin))

	a.False(isConsole(new(bytes.Buffer)))
	a.False(isConsole(nil))
}
