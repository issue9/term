// SPDX-License-Identifier: MIT

package colors

import (
	"testing"

	"github.com/issue9/assert"
)

func TestColor(t *testing.T) {
	a := assert.New(t)

	c := Red
	a.Equal(c.String(), "Red")

	c = 100
	a.Equal(c.String(), "100")
}
