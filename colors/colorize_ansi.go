// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// +build !windows

package colors

import (
	"errors"
)

type Colorize struct {
	out        int
	foreground Color
	background Color
}

// 新建一个Colorize
func New(out int, foreground, background Color) *Colorize {
	return &Colorize{
		out:        out,
		foreground: foreground,
		background: background,
	}
}

// 更改颜色值。
// foreground 文字颜色；
// background 背景色。
func (c *Colorize) SetColor(foreground, background Color) {
	c.foreground = foreground
	c.background = background
}

// 更改输出方向。
// out参数只能是Stdout和Stderr两种类型，其它值为返回错误内容。
func (c *Colorize) SetOut(out int) error {
	if out != Stdout && out != Stderr {
		return errors.New("out参数只能是Stdout或是Stderr")
	}

	c.out = out
	return nil
}

// 等同于fmt.Print()，颜色和输出方向由Colorize指定
func (c *Colorize) Print(v ...interface{}) (int, error) {
	return Print(c.out, c.foreground, c.background, v...)
}

// 等同于fmt.Println()，颜色和输出方向由Colorize指定
func (c *Colorize) Println(v ...interface{}) (int, error) {
	return Println(c.out, c.foreground, c.background, v...)
}

// 等同于fmt.Printf()，颜色和输出方向由Colorize指定
func (c *Colorize) Printf(format string, v ...interface{}) (int, error) {
	return Printf(c.out, c.foreground, c.background, format, v...)
}
