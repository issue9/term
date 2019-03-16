// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package prompt 简单的终端交互界面
package prompt

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/issue9/term/colors"
)

// Prompt 终端交互对象
type Prompt struct {
	reader       *bufio.Reader
	output       io.Writer
	delim        byte
	defaultColor colors.Color
	err          error
}

// New 声明 Prompt 变量
//
// delim 从 input 读取内容时的分隔符，如果为空，则采用 \n；
// defaultColor 默认值的颜色，如果该值无效，则会 panic。
func New(delim byte, input io.Reader, output io.Writer, defaultColor colors.Color) *Prompt {
	if delim == 0 {
		delim = '\n'
	}

	if !defaultColor.IsValid() {
		panic("无效的颜色值 defaultColor")
	}

	return &Prompt{
		reader:       bufio.NewReader(input),
		output:       output,
		delim:        delim,
		defaultColor: defaultColor,
	}
}

func (p *Prompt) println(c colors.Color, v ...interface{}) {
	if p.err == nil {
		_, p.err = colors.Fprintln(p.output, c, colors.Default, v...)
	}
}

func (p *Prompt) print(c colors.Color, v ...interface{}) {
	if p.err == nil {
		_, p.err = colors.Fprint(p.output, c, colors.Default, v...)
	}
}

func (p *Prompt) printf(c colors.Color, format string, v ...interface{}) {
	if p.err == nil {
		_, p.err = colors.Fprintf(p.output, c, colors.Default, format, v...)
	}
}

// 从输入端读取一行内容
func (p *Prompt) read() (v string) {
	if p.err != nil {
		return
	}

	if v, p.err = p.reader.ReadString(p.delim); p.err == nil {
		return v[:len(v)-1]
	}
	return
}

// String 输出问题，并获取用户的回答内容
//
// q 显示的问题内容；
// def 表示默认值。
func (p *Prompt) String(q, def string) (string, error) {
	p.print(colors.Default, q)
	if def != "" {
		p.print(p.defaultColor, "（", def, "）")
	}
	p.print(colors.Default, "：")

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
	p.print(colors.Default, q)
	str := "Y"
	if !def {
		str = "N"
	}
	p.print(p.defaultColor, "（", str, "）：")
	p.print(colors.Default, "：")

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
	p.println(colors.Default, q)
	for i, v := range slice {
		c := colors.Default
		if inIntSlice(i, def) {
			c = p.defaultColor
		}
		p.printf(c, "（%d）", i)
		p.printf(colors.Default, "%s\n", v)
	}
	p.print(colors.Default, "请输入你的选择项，多项请用半角逗号（,）分隔：")

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
	p.println(colors.Default, q)
	for k, v := range maps {
		c := colors.Default
		if inStringSlice(k, def) {
			c = p.defaultColor
		}
		p.printf(c, "（%s）", k)
		p.printf(colors.Default, "%s\n", v)
	}
	p.print(colors.Default, "请输入你的选择项，多项请用半角逗号（,）分隔：")

	val := p.read()

	if p.err != nil {
		return nil, p.err
	}

	if val == "" {
		return def, nil
	}
	return strings.Split(val, ","), nil
}

func inIntSlice(v int, vals []int) bool {
	for _, val := range vals {
		if val == v {
			return true
		}
	}

	return false
}

func inStringSlice(v string, vals []string) bool {
	for _, val := range vals {
		if val == v {
			return true
		}
	}

	return false
}
