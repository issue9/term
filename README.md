term [![Build Status](https://travis-ci.org/issue9/term.svg?branch=master)](https://travis-ci.org/issue9/term)
======

term包含了终端控制相关的包：ansi和colors。

其中ansi定义了ANSI escape codes操作的相关内容，windows默认情况下不支持
ANSI escape codes，但可以通过ansi.sys或类似功能来支持ansi；
colors则是平台通用的控制台彩色字符串输出包。定义了一些类似fmt包的函数，
用于字符串的输出。

```go
// 向stderr终端输出蓝底红字的字符串："colors"
colors.Print(colors.Stdout, colors.Red, colors.Blue, "colors")

// 输出蓝底红字的字符串
c := colors.New(colors.Stdout, Colors.Red, colors.Blue)
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

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/term)
[![GoDoc](https://godoc.org/github.com/issue9/term?status.svg)](https://godoc.org/github.com/issue9/term)


### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
