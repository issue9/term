// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// 输出ansi控制码，不支持windows系统
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
// ansi的相关文档，可参考以下内容：
//  http://en.wikipedia.org/wiki/ANSI_escape_code
//  http://www.mudpedia.org/mediawiki/index.php/ANSI_colors
package ansi

// 当前库的版本
const Version = "0.1.1.141223"
