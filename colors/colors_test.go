// SPDX-License-Identifier: MIT

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
