// SPDX-License-Identifier: MIT

package ansi

import (
	"fmt"
	"strconv"
)

// ESC 表示 ansi 转码序列
type ESC string

// FColor 生成标准前景色彩值
func FColor(color uint8) ESC {
	if color > 7 {
		panic("参数 color 不参大于 7")
	}

	return CSI('m', 30+int(color))
}

// BColor 生成标准背景色彩值
func BColor(color uint8) ESC {
	if color > 7 {
		panic("参数 color 不参大于 7")
	}

	return CSI('m', 40+int(color))
}

// FBrightColor 生成高强度前景色彩值
func FBrightColor(color uint8) ESC {
	if color > 7 {
		panic("参数 color 不参大于 7")
	}

	return CSI('m', 90+int(color))
}

// BBrightColor 生成高强度背景色彩值
func BBrightColor(color uint8) ESC {
	if color > 7 {
		panic("参数 color 不参大于 7")
	}

	return CSI('m', 100+int(color))
}

// F256Color 获取扩展的前景颜色值控制码
func F256Color(color uint8) ESC {
	return CSI('m', 38, 5, int(color))
}

// B256Color 获取扩展的背景颜色值控制码
func B256Color(color uint8) ESC {
	return CSI('m', 48, 5, int(color))
}

// FTrueColor 返回真色彩的前景颜色值
func FTrueColor(r, g, b uint8) ESC {
	return CSI('m', 38, 2, int(r), int(g), int(b))
}

// BTrueColor 返回真色彩的背景颜色值
func BTrueColor(r, g, b uint8) ESC {
	return CSI('m', 48, 2, int(r), int(g), int(b))
}

// CUB 光标后移
func CUB(n int) ESC {
	return CSI('D', n)
}

// CUF 光标前移
func CUF(n int) ESC {
	return CSI('C', n)
}

// CUU 光标上移 n
func CUU(n int) ESC {
	return CSI('A', n)
}

// CUD 光标下移 n
func CUD(n int) ESC {
	return CSI('B', n)
}

// CNL 光标下移 n 行并至行首
func CNL(n int) ESC {
	return CSI('E', n)
}

// CPL 光标上移 n 行并至行首
func CPL(n int) ESC {
	return CSI('F', n)
}

// CHA 光标移至 n 列位置
func CHA(n int) ESC {
	return CSI('G', n)
}

// CUP 移动光标到 x,y 的位置
func CUP(x, y int) ESC {
	return CSI('H', x, y)
}

// SU 向上滚动 n 行并至行尾
func SU(n int) ESC {
	return CSI('S', n)
}

// SD 向下滚动 n 行并至行尾
func SD(n int) ESC {
	return CSI('T', n)
}

// DSR 设备状态报告
func DSR() ESC {
	return CSI('n', 6)
}

// SCP 保存光标位置
func SCP() ESC {
	return CSI('s')
}

// RCP 恢复光标位置
func RCP() ESC {
	return CSI('u')
}

// ED 返回清除屏幕的控制符
//
// n == 0 时，清除从当前光标到屏幕尾的所有字符；
// n == 1 时，清除从当前光标到屏幕头的所有字符；
// n == 2 时，清除当前屏幕的所有字符；
// 当 n 为其它值时，将触发 panic
func ED(n int) ESC {
	if n < 0 || n > 2 {
		panic(fmt.Sprintf("参数 n [%v]必须介于 [0,2]", n))
	}
	return CSI('J', n)
}

// EL 获取清除行的控制符
//
// n == 0 时，清除从当前光标到行尾的所有字符；
// n == 1 时，清除从当前光标到行头的所有字符；
// n == 2 时，清除当前行的所有字符。
// 当 n 为其它值时，将触发 panic
func EL(n int) ESC {
	if n < 0 || n > 2 {
		panic(fmt.Sprintf("参数 n [%v]必须介于 [0,2]", n))
	}
	return CSI('K', n)
}

// CSI 生成 CSI 指令
func CSI(end byte, v ...int) ESC {
	if len(v) == 0 {
		return ESC("\033[" + string(end))
	}

	esc := "\033["
	for _, item := range v {
		esc += strconv.Itoa(item) + ";"
	}
	return ESC(esc[:len(esc)-1] + string(end))
}
