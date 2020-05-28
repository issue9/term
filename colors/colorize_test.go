// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"testing"

	"github.com/issue9/assert"
)

func TestColorize(t *testing.T) {
	a := assert.New(t)

	for b := Color(0); b < max; b++ {
		for f := Color(0); f < max; f++ {
			c := New(f, b)
			_, err := c.Printf("%s:%s\t", b.String(), f.String())
			a.NotError(err)
		}
		fmt.Println()
	}
}
