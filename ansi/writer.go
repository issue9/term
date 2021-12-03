// SPDX-License-Identifier: MIT

package ansi

import (
	"io"

	"github.com/issue9/errwrap"
)

// Writer ansi 控制码的 io.Writer 接口
//
//  a := NewWriter(os.Stdout)
//
//  a.Left(5)
//  a.ClearLine(2)
//  a.SGR(term.SGRFRed,term.SGRBGreen)
//  a.Print("abc")
//
//  fmt.Fprintf(a, "%v", term.SGRFBBlue)
type Writer struct {
	w errwrap.Writer
}

// NewWriter 声明一个 Writer 结构体
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: errwrap.Writer{Writer: w}}
}

// WriteESC 输出字符串
func (w *Writer) WriteESC(esc ESC) *Writer { return w.WBytes([]byte(esc)) }

// Writer 暴露原始的 io.Writer 接口
//
// 此接口的出错误信息会直接返回，并不会记录在 Writer.Err 之中。
func (w *Writer) Write(bs []byte) (int, error) { return w.w.Write(bs) }

// Left 左移 n 个字符光标
func (w *Writer) Left(n int) *Writer { return w.WriteESC(CUB(n)) }

// Right 右移 n 个字符光标
func (w *Writer) Right(n int) *Writer { return w.WriteESC(CUF(n)) }

// Up 上移 n 行光标
func (w *Writer) Up(n int) *Writer { return w.WriteESC(CUU(n)) }

// Down 下移 n 行光标
func (w *Writer) Down(n int) *Writer { return w.WriteESC(CUD(n)) }

// Erase 清除屏幕
//
// n==0 时，清除从当前光标到屏幕尾的所有字符；
// n==1 时，清除从当前光标到屏幕头的所有字符；
// n==2 时，清除当前屏幕的所有字符；
// 当 n 为其它值时，将触发 panic
func (w *Writer) Erase(n int) *Writer { return w.WriteESC(ED(n)) }

// EraseLine 清除行
//
// n==0 时，清除从当前光标到行尾的所有字符；
// n==1 时，清除从当前光标到行头的所有字符；
// n==2 时，清除当前行的所有字符；
// 当 n 为其它值时，将触发 panic
func (w *Writer) EraseLine(n int) *Writer { return w.WriteESC(EL(n)) }

// Move 移动光标到 x,y 的位置
func (w *Writer) Move(x, y int) *Writer { return w.WriteESC(CUP(x, y)) }

// SaveCursor 保存光标位置
func (w *Writer) SaveCursor() *Writer { return w.WriteESC(SCP()) }

// RestoreCursor 恢复光标位置
func (w *Writer) RestoreCursor() *Writer { return w.WriteESC(RCP()) }

// SGR 输出 SGR 指令
func (w *Writer) SGR(sgr ...int) *Writer { return w.WriteESC(SGR(sgr...)) }

// FTrueColor 输出 24 色彩前景色
func (w *Writer) FTrueColor(r, g, b uint8) *Writer {
	return w.WriteESC(FTrueColor(r, g, b))
}

// BTrueColor 输出 24 色彩背景色
func (w *Writer) BTrueColor(r, g, b uint8) *Writer {
	return w.WriteESC(BTrueColor(r, g, b))
}

// F256Color 输出 256 色的背景颜色信息
func (w *Writer) F256Color(color uint8) *Writer {
	return w.WriteESC(F256Color(color))
}

// B256Color 输出 256 色的背景颜色信息
func (w *Writer) B256Color(color uint8) *Writer {
	return w.WriteESC(B256Color(color))
}

// TrueColor 输出 24 位色彩的颜色信息
//
// fr、fb 和 fb 表示前景色的 RGB 颜色值；
// br、bb 和 bb 表示背景色的 RGB 颜色值；
func (w *Writer) TrueColor(fr, fg, fb, br, bg, bb uint8) *Writer {
	w.WriteESC(FTrueColor(fr, fg, fb))
	return w.WriteESC(BTrueColor(br, bg, bb))
}

// Color256 输出 256 色的颜色信息
func (w *Writer) Color256(f, b uint8) *Writer {
	w.WriteESC(F256Color(f))
	return w.WriteESC(B256Color(b))
}

// Printf 相当于 fmt.Printf
func (w *Writer) Printf(format string, args ...interface{}) *Writer {
	w.w.Printf(format, args...)
	return w
}

// Print 相当于 fmt.Print
func (w *Writer) Print(args ...interface{}) *Writer {
	w.w.Print(args...)
	return w
}

// Println 相当于 fmt.Println
func (w *Writer) Println(args ...interface{}) *Writer {
	w.w.Println(args...)
	return w
}

// WString 写入字符串
func (w *Writer) WString(str string) *Writer {
	w.w.WString(str)
	return w
}

// WByte 写入单个字节内容
func (w *Writer) WByte(b byte) *Writer { return w.WBytes([]byte{b}) }

// WBytes 写入字节内容
func (w *Writer) WBytes(data []byte) *Writer {
	w.w.WBytes(data)
	return w
}

// Err 返回写入过程中的错误
//
// 以链式的方式调用，中途如果出错，后续不会再写入，但是不会中断链式的调用。
// 可以通过此值判断写入途中是否存在错误。
func (w *Writer) Err() error { return w.w.Err }
