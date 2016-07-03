term [![Build Status](https://travis-ci.org/issue9/term.svg?branch=master)](https://travis-ci.org/issue9/term)
======

term 包含了终端控制相关的包：ansi 和 colors。

其中 ansi 定义了 ANSI escape codes 操作的相关内容，windows 默认情况下不支持
ANSI escape codes，但可以通过 ansi.sys 或类似功能来支持 ansi；
colors 则是平台通用的控制台彩色字符串输出包。定义了一些类似fmt包的函数，
用于字符串的输出。

```go
// 向 stderr 终端输出蓝底红字的字符串："colors"
colors.Print(colors.Stdout, colors.Red, colors.Blue, "colors")

// 输出蓝底红字的字符串
c := colors.New(colors.Stdout, colors.Red, colors.Blue)
c.Println("colors")

// 输出黑底绿字的字符串
c.SetColor(colors.Green, colors.Black)
c.Printf("colors:%v,%v", colors.Green, colors.Black)
```

### 安装

```shell
go get github.com/issue9/term
```


### 文档

[![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/issue9/term)
[![GoDoc](https://godoc.org/github.com/issue9/term?status.svg)](https://godoc.org/github.com/issue9/term)


### 版权

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
