// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package prompt 简单的终端交互界面
package prompt

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const delim = '\n'

// Prompt 终端交互对象
type Prompt struct {
	reader *bufio.Reader
	output io.Writer
	err    error
}

// New 声明 Prompt 变量
func New(input io.Reader, output io.Writer) *Prompt {
	return &Prompt{
		reader: bufio.NewReader(input),
		output: output,
	}
}

func (p *Prompt) println(v ...interface{}) {
	if p.err == nil {
		_, p.err = fmt.Fprintln(p.output, v...)
	}
}

func (p *Prompt) print(v ...interface{}) {
	if p.err == nil {
		_, p.err = fmt.Fprint(p.output, v...)
	}
}

func (p *Prompt) printf(format string, v ...interface{}) {
	if p.err == nil {
		_, p.err = fmt.Fprintf(p.output, format, v...)
	}
}

// 从输入端读取一行内容
func (p *Prompt) read() (v string) {
	if p.err != nil {
		return
	}

	if v, p.err = p.reader.ReadString(delim); p.err == nil {
		return v[:len(v)-1]
	}
	return
}

// String 输出问题，并获取用户的回答内容
//
// q 显示的问题内容；
// def 表示默认值。
func (p *Prompt) String(q, def string) (string, error) {
	p.print(q)
	if def != "" {
		p.print("(", def, ")")
	}
	p.print(":")

	v := p.read()

	if p.err != nil {
		return "", p.err
	}

	if v == "" {
		v = def
	}
	return v, nil
}

// Bool 输出 bool 问题，并获取用户的回答内容
func (p *Prompt) Bool(q string, def bool) (bool, error) {
	str := "Y"
	if !def {
		str = "N"
	}
	p.printf("%s(%s)\n", q, str)

	val := p.read()

	if p.err != nil {
		return false, p.err
	}

	switch strings.ToLower(val) {
	case "yes", "y":
		return true, nil
	case "no", "n":
		return false, nil
	default:
		return def, nil
	}
}

// Slice 输出一个选择性问题，并获取用户的选择项
//
// q 表示问题内容；
// slice 表示可选的问题列表；
// def 表示默认项的索引，必须在 slice 之内。
func (p *Prompt) Slice(q string, slice []string, def ...int) (selected []int, err error) {
	p.println(q)
	for i, v := range slice {
		p.printf("(%d) %s\n", i, v)
	}
	p.print("请输入你的选择项:")

	val := p.read()

	if p.err != nil {
		return nil, p.err
	}

	if val == "" {
		return def, nil
	}

	for _, v := range strings.Split(val, ",") {
		vv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		selected = append(selected, vv)
	}
	return selected, nil
}

// Map 输出一个单选问题，并获取用户的选择项
//
// q 表示问题内容；
// maps 表示可选的问题列表；
// def 表示默认项的索引，必须在 maps 之内。
func (p *Prompt) Map(q string, maps map[string]string, def ...string) (selected []string, err error) {
	p.println(q)
	for k, v := range maps {
		p.printf("(%s) %s", k, v)
	}
	p.print("请输入你的选择项:")

	val := p.read()

	if p.err != nil {
		return nil, p.err
	}

	if val == "" {
		return def, nil
	}
	return strings.Split(val, ","), nil
}
