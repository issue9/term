// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package colors

type Color int

// 颜色值定义
const (
	Default Color = iota // 默认色，由控制台决定具体颜色值
	Black                // 黑色
	Red                  // 红色
	Green                // 绿色
	Yellow               // 黄色
	Blue                 // 蓝色
	Magenta              // 洋红色
	Cyan                 // 青色
	White                // 白色
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
