// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package colors

import (
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
