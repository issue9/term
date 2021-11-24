// SPDX-License-Identifier: MIT

package ansi

import (
	"math"
	"testing"

	"github.com/issue9/assert/v2"
)

func TestF256Color(t *testing.T) {
	a := assert.New(t, false)

	// 临界点
	a.Equal(F256Color(0), "\033[38;5;0m")
	a.Equal(F256Color(255), "\033[38;5;255m")

	// 正常测试
	a.Equal(F256Color(211), "\033[38;5;211m")
}

func TestB256Color(t *testing.T) {
	a := assert.New(t, false)

	// 临界点
	a.Equal(B256Color(0), "\033[48;5;0m")
	a.Equal(B256Color(255), "\033[48;5;255m")

	// 正常测试
	a.Equal(B256Color(211), "\033[48;5;211m")
}

func TestCursor(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(CUB(5), "\033[5D")
	a.Equal(CUF(5), "\033[5C")
	a.Equal(CUU(5), "\033[5A")
	a.Equal(CUD(5), "\033[5B")
}

func TestED(t *testing.T) {
	a := assert.New(t, false)

	a.Panic(func() { ED(-1) })
	a.Panic(func() { ED(3) })

	a.Equal(ED(0), "\033[0J")
	a.Equal(ED(2), "\033[2J")
	a.Equal(ED(1), "\033[1J")
}

func TestEL(t *testing.T) {
	a := assert.New(t, false)

	a.Panic(func() { EL(-1) })
	a.Panic(func() { EL(3) })

	a.Equal(EL(0), "\033[0K")
	a.Equal(EL(2), "\033[2K")
	a.Equal(EL(1), "\033[1K")
}

func TestCUP(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(CUP(3, 2), "\033[3;2H")
}

func TestColor(t *testing.T) {
	for i := uint8(0); i < math.MaxUint8; i++ {
		t.Logf("%v%v字体颜色%d, 背景颜色%d%v\n", F256Color(i), B256Color(255-i), i, 255-i, CSI(ResetCode))
	}
}
