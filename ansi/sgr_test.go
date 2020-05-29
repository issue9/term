// SPDX-License-Identifier: MIT

package ansi

import (
	"strconv"
	"testing"

	"github.com/issue9/assert"
)

func TestSGR(t *testing.T) {
	// 多个
	sgr := SGR(SGRBBlack, SGRBBlue, SGRFRed)
	result := "\033[" + strconv.Itoa(SGRBBlack) + ";" + strconv.Itoa(SGRBBlue) + ";" + strconv.Itoa(SGRFRed) + "m"
	assert.Equal(t, sgr, result)

	// 传递单个值
	sgr = SGR(SGRBBlack)
	result = "\033[" + strconv.Itoa(SGRBBlack) + "m"
	assert.Equal(t, sgr, result)

	// 传递空值，相当于SReset
	sgr = SGR()
	assert.Equal(t, sgr, CSI('m', SGRReset))
}
