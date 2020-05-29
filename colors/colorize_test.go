// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"testing"

	"github.com/issue9/assert"
)

func TestColorize(t *testing.T) {
	a := assert.New(t)

	for b := Default; b < maxNamedColor; b++ {
		for f := Default; f < maxNamedColor; f++ {
			c := New(Italic, f, b)
			_, err := c.Printf("%s:%s\t", b.String(), f.String())
			a.NotError(err)
		}
		fmt.Println()
	}
}
