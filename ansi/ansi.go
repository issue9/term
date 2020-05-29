// SPDX-License-Identifier: MIT

// Package ansi 输出 ansi 控制码
//
// windows 系统默认情况下不支持 ansi 控制码。
// 若仅仅是需要输出彩色字符到控制台，请使用 term/colors 包。
//
//  fmt.Printf("%v这是红色的字", term.FRed)
//  fmt.Printf("%v这是红色字，绿色背景", term.SGR(term.SGRFRed,term.SGRBGreen))
//  fmt.Printf("%v%v这是红色字，绿色背景", term.FRed,term.BGreen)
//
//  // 包装一个stderr。
//  w := term.NewWriter(os.Stderr)
//  w.Left(5)
//  w.SGR(term.SGRFRed)
//  w.Printf("%s", "string")
//  w.Move(1,1)
//
// ansi 的相关文档，可参考以下内容：
//  https://en.wikipedia.org/wiki/ANSI_escape_code
//  https://zh.wikipedia.org/wiki/ANSI%E8%BD%AC%E4%B9%89%E5%BA%8F%E5%88%97
package ansi
