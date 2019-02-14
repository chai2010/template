// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"text/template"
)

type FuncMap = template.FuncMap

func Render(tmpl string, data interface{}, funcMap ...FuncMap) (string, error) {
	t := template.New("")

	for _, fnMap := range funcMap {
		if len(fnMap) > 0 {
			t = t.Funcs(fnMap)
		}
	}

	t, err := t.Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func RenderMap(tmpl string, m map[string]interface{}) (string, error) {
	var (
		dataMap = make(map[string]interface{})
		fnMap   = make(FuncMap)
	)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() != reflect.Func {
			dataMap[k] = v
		} else {
			fnMap[k] = v
		}
	}

	return Render(tmpl, dataMap, fnMap)
}

func RenderWithDelims(tmpl, left, right string, data interface{}, funcMap ...FuncMap) (string, error) {
	t := template.New("").Delims(left, right)

	for _, fnMap := range funcMap {
		if len(fnMap) > 0 {
			t = t.Funcs(fnMap)
		}
	}

	t, err := t.Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func RenderFile(filename string, data interface{}, funcMap ...FuncMap) (string, error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return Render(string(s), data, funcMap...)
}

func RenderFileWithDelims(filename, left, right string, data interface{}, funcMap ...FuncMap) (string, error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return RenderWithDelims(string(s), left, right, data, funcMap...)
}

func MustRender(tmpl string, data interface{}, funcMap ...FuncMap) string {
	s, err := Render(tmpl, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderMap(tmpl string, m map[string]interface{}) string {
	s, err := RenderMap(tmpl, m)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderWithDelims(tmpl, left, right string, data interface{}, funcMap ...FuncMap) string {
	s, err := RenderWithDelims(tmpl, left, right, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderFile(filename string, data interface{}, funcMap ...FuncMap) string {
	s, err := RenderFile(filename, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderFileWithDelims(filename, left, right string, data interface{}, funcMap ...FuncMap) string {
	s, err := RenderFileWithDelims(filename, left, right, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}
