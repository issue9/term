// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// colors 带色彩的控制台文本输出包，兼容windows平台。
//  // 输出段蓝底红字：foreground:Red;background:Blue
//  colors.Printf(Stdout, colors.Red, colors.Blue, "foreground:%v;background:%v", colors.Red, colors.Blue)
//
//  // 功能同上，但是可以重复调用Print*系列函数输出内容。
//  c := colors.New(colors.Stdout, colors.Red, colors.Yellow)
//  c.Print("foreground:%v;background:%v")
//  c.Print(colors.Red, colors.Blue)
//
// 为了兼容windows平台，只使用了最基本的几种颜色值，
// 而不是ansi控制台的256色。若不需要考虑windows平台，
// 可以直接使用term/ansi包，那里有对ansi包更好的支持。
package colors

// 包的版本号
const Version = "0.1.2.150211"
