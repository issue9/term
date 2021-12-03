// SPDX-License-Identifier: MIT

package colors

import (
	"io"

	"github.com/issue9/term/v2/ansi"
)

// Colorize 适合固定颜色的大段内容输出
type Colorize struct {
	sgr   ansi.ESC
	reset ansi.ESC
	w     *ansi.Writer
}

// New 新建一个 Colorize
func New(w io.Writer, t Type, foreground, background Color) *Colorize {
	if !isValidType(t) {
		panic("无效的参数 t")
	}

	codes := make([]int, 0, 10)
	if t != Normal {
		codes = append(codes, int(t))
	}
	codes = append(codes, foreground.fColorCode()...)
	codes = append(codes, background.bColorCode()...)

	return &Colorize{
		sgr:   ansi.SGR(codes...),
		reset: ansi.SGR(ansi.ResetCode),
		w:     ansi.NewWriter(w),
	}
}

func (c *Colorize) Print(v ...interface{}) *Colorize {
	c.w.WriteESC(c.sgr).Print(v...).WriteESC(c.reset)
	return c
}

func (c *Colorize) Println(v ...interface{}) *Colorize {
	c.w.WriteESC(c.sgr).Println(v...).WriteESC(c.reset)
	return c
}

func (c *Colorize) Printf(format string, v ...interface{}) *Colorize {
	c.w.WriteESC(c.sgr).Printf(format, v...).WriteESC(c.reset)
	return c
}

// WString 写入字符串
func (c *Colorize) WString(str string) *Colorize {
	c.w.WString(str)
	return c
}

// WByte 写入单个字节内容
func (c *Colorize) WByte(b byte) *Colorize { return c.WBytes([]byte{b}) }

// WBytes 写入字节内容
func (c *Colorize) WBytes(data []byte) *Colorize {
	c.w.WBytes(data)
	return c
}

func (c *Colorize) Err() error { return c.w.Err() }
