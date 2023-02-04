// SPDX-License-Identifier: MIT

// Package ansi 输出 ansi 控制码
//
//	// 包装一个 stderr。
//	w := ansi.NewWriter(os.Stderr)
//	w.Left(5)
//	w.Printf("%s", "string")
//	w.Move(1,1)
//
// ansi 的相关文档，可参考以下内容 [ansi]、[ANSI转义序列]：
//
// [ansi]: https://en.wikipedia.org/wiki/ANSI_escape_code
// [ANSI转义序列]: https://zh.wikipedia.org/wiki/ANSI%E8%BD%AC%E4%B9%89%E5%BA%8F%E5%88%97
package ansi
