// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ansi

import (
	"testing"

	"github.com/issue9/assert"
)

func TestFColor256(t *testing.T) {
	a := assert.New(t)

	// panic
	a.Panic(func() { FColor256(256) })
	a.Panic(func() { FColor256(-1) })

	// 临界点
	a.Equal(FColor256(0), "\033[38;5;0m")
	a.Equal(FColor256(255), "\033[38;5;255m")

	// 正常测试
	a.Equal(FColor256(211), "\033[38;5;211m")

}

func TestBColor256(t *testing.T) {
	a := assert.New(t)

	// panic
	a.Panic(func() { BColor256(256) })
	a.Panic(func() { BColor256(-1) })

	// 临界点
	a.Equal(BColor256(0), "\033[48;5;0m")
	a.Equal(BColor256(255), "\033[48;5;255m")

	// 正常测试
	a.Equal(BColor256(211), "\033[48;5;211m")
}

func TestLeftRightUpDown(t *testing.T) {
	a := assert.New(t)

	a.Equal(Left(5), "\033[5D")
	a.Equal(Right(5), "\033[5C")
	a.Equal(Up(5), "\033[5A")
	a.Equal(Down(5), "\033[5B")
}

func TestErase(t *testing.T) {
	a := assert.New(t)

	a.Panic(func() { Erase(-1) })
	a.Panic(func() { Erase(3) })

	a.Equal(Erase(0), "\033[0J")
	a.Equal(Erase(2), "\033[2J")
	a.Equal(Erase(1), "\033[1J")
}

func TestEraseLine(t *testing.T) {
	a := assert.New(t)

	a.Panic(func() { EraseLine(-1) })
	a.Panic(func() { EraseLine(3) })

	a.Equal(EraseLine(0), "\033[0K")
	a.Equal(EraseLine(2), "\033[2K")
	a.Equal(EraseLine(1), "\033[1K")
}

func TestMove(t *testing.T) {
	a := assert.New(t)

	a.Equal(Move(3, 2), "\033[3;2H")
}

func TestColor(t *testing.T) {
	t.Logf("%v%vFRed, BDefault%v\n", FRed, BDefault, Reset)
	t.Logf("%v%vFGreen, BWhite%v\n", FGreen, BWhite, Reset)
	t.Logf("%v%vFYellow, BCyan%v\n", FYellow, BCyan, Reset)
	t.Logf("%v%vFBlue, BMagenta%v\n", FBlue, BMagenta, Reset)
	t.Logf("%v%vFMagenta, BBlue%v\n", FMagenta, BBlue, Reset)
	t.Logf("%v%vFCyan, BYellow%v\n", FCyan, BYellow, Reset)
	t.Logf("%v%vFWhite, BGreen%v\n", FWhite, BGreen, Reset)
	t.Logf("%v%vFDefault, BRed%v\n", FDefault, BRed, Reset)

	for i := 0; i < 256; i += 10 {
		t.Logf("%v%v字体颜色%d, 背景颜色%d%v\n", FColor256(i), BColor256(255-i), i, 255-i, Reset)
	}
}
