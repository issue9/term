// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package logs 向控制台输出一条日志信息。
// 日志信息带有一个彩色的前缀文字，在部分轻量的命令行程序可以使用。
package logs

import (
	"io"
	"os"

	"github.com/issue9/term/colors"
)

// 几个预定义的输出通道
var (
	Error   = New(os.Stderr, "[ERRO] ", colors.Red)
	Warn    = New(os.Stderr, "[WARN] ", colors.Magenta)
	Success = New(os.Stdout, "[SUCC] ", colors.Green)
	Info    = New(os.Stdout, "[INFO] ")
)

// Logger 控制台的日志输出
type Logger struct {
	w             io.Writer
	Prefix        string
	prefixPrinter colors.Colorize
	textPrinter   colors.Colorize
}

// New 声明 Logger
func New(w io.Writer, prefix string, color ...colors.Color) *Logger {
	l := &Logger{
		w:      w,
		Prefix: prefix,
	}

	switch len(color) {
	case 0:
		l.prefixPrinter = colors.New(colors.Default, colors.Default)
		l.textPrinter = colors.New(colors.Default, colors.Default)
	case 1:
		l.prefixPrinter = colors.New(color[0], colors.Default)
		l.textPrinter = colors.New(colors.Default, colors.Default)
	case 2:
		l.prefixPrinter = colors.New(color[0], color[1])
		l.textPrinter = colors.New(colors.Default, colors.Default)
	case 3:
		l.prefixPrinter = colors.New(color[0], color[1])
		l.textPrinter = colors.New(color[2], colors.Default)
	default:
		l.prefixPrinter = colors.New(color[0], color[1])
		l.textPrinter = colors.New(color[2], color[3])
	}

	return l
}

func (l *Logger) Print(v ...interface{}) (int, error) {
	n, err := l.prefixPrinter.Fprint(l.w, l.Prefix)
	if err != nil {
		return n, err
	}

	return l.textPrinter.Fprint(l.w, v...)
}

func (l *Logger) Printf(format string, v ...interface{}) (int, error) {
	n, err := l.prefixPrinter.Fprint(l.w, l.Prefix)
	if err != nil {
		return n, err
	}

	return l.textPrinter.Fprintf(l.w, format, v...)
}

func (l *Logger) Println(v ...interface{}) (int, error) {
	n, err := l.prefixPrinter.Fprint(l.w, l.Prefix)
	if err != nil {
		return n, err
	}

	return l.textPrinter.Fprintln(l.w, v...)
}
