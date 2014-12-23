// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// colors 带色彩的控制台文本输出包，兼容windows平台。
package colors

// 包的版本号
const Version = "0.1.0.141222"

type Color int

// 颜色值定义
const (
	Default Color = iota
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// 输出方向，只能是Stderr和Stdout，
// 若系统对这这两个做了重定向，则输出内容可能出错。
const (
	Stderr = iota
	Stdout
)

// fmt.Stringer.String()
func (c Color) String() string {
	switch c {
	case Default:
		return "Default"
	case Black:
		return "Black"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Blue:
		return "Blue"
	case Magenta:
		return "Magenta"
	case Cyan:
		return "Cyan"
	case White:
		return "White"
	default:
		return "<unknown>"
	}
}
