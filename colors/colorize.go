// SPDX-License-Identifier: MIT

package colors

import (
	"io"

	"github.com/issue9/term/v2/ansi"
)

// Colorize 适合固定颜色的大段内容输出
type Colorize struct {
	w *ansi.Writer
}

// New 新建一个 Colorize
func New(w io.Writer) *Colorize {
	if ww, ok := w.(*Colorize); ok {
		return ww
	}
	return &Colorize{w: ansi.NewWriter(w)}
}

// Writer 暴露原始的 io.Writer 接口
//
// 此接口的出错误信息会直接返回，并不会记录在 Writer.Err 之中。
func (c *Colorize) Write(bs []byte) (int, error) { return c.w.Write(bs) }

func (c *Colorize) Color(t Type, foreground, background Color) *Colorize {
	if !isValidType(t) {
		panic("无效的参数 t")
	}

	codes := make([]int, 0, 10)
	if t != Normal {
		codes = append(codes, int(t))
	}
	codes = append(codes, foreground.fColorCode()...)
	codes = append(codes, background.bColorCode()...)

	c.w.WriteESC(ansi.SGR(codes...))

	return c
}

func (c *Colorize) Reset() *Colorize {
	c.w.WriteESC(ansi.SGR(ansi.ResetCode))
	return c
}

func (c *Colorize) Print(v ...interface{}) *Colorize {
	c.w.Print(v...)
	return c
}

func (c *Colorize) Println(v ...interface{}) *Colorize {
	c.w.Println(v...)
	return c
}

func (c *Colorize) Printf(format string, v ...interface{}) *Colorize {
	c.w.Printf(format, v...)
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
