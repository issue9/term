// SPDX-License-Identifier: MIT

package colors

import (
	"fmt"
	"testing"

	"github.com/issue9/assert"
)

func TestColorize(t *testing.T) {
	a := assert.New(t)

	for b := Default; b < max; b++ {
		for f := Default; f < max; f++ {
			c := New(f, b)
			_, err := c.Printf("%s:%s\t", b.String(), f.String())
			a.NotError(err)
		}
		fmt.Println()
	}
}
