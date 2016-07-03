// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// colors 带色彩的控制台文本输出包，兼容 windows 平台。
//  // 输出段蓝底红字：foreground:Red;background:Blue
//  colors.Printf(Stdout, colors.Red, colors.Blue, "foreground:%v;background:%v", colors.Red, colors.Blue)
//
//  // 功能同上，但是可以重复调用Print*系列函数输出内容。
//  c := colors.New(colors.Stdout, colors.Red, colors.Yellow)
//  c.Print("foreground:%v;background:%v")
//  c.Print(colors.Red, colors.Blue)
//
// 为了兼容 windows 平台，只使用了最基本的几种颜色值，
// 而不是 ansi 控制台的 256 色。若不需要考虑 windows 平台，
// 可以直接使用 term/ansi 包，那里有对 ansi 包更好的支持。
package colors
