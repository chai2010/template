// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"reflect"
)

type HTMLFuncMap = template.FuncMap

func RenderHTML(tmpl string, data interface{}, funcMap ...HTMLFuncMap) (string, error) {
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

func RenderHTMLMap(tmpl string, m map[string]interface{}) (string, error) {
	var (
		dataMap = make(map[string]interface{})
		fnMap   = make(HTMLFuncMap)
	)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() != reflect.Func {
			dataMap[k] = v
		} else {
			fnMap[k] = v
		}
	}

	return RenderHTML(tmpl, dataMap, fnMap)
}

func RenderHTMLWithDelims(tmpl, left, right string, data interface{}, funcMap ...HTMLFuncMap) (string, error) {
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

func RenderHTMLFile(filename string, data interface{}, funcMap ...HTMLFuncMap) (string, error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return RenderHTML(string(s), data, funcMap...)
}

func RenderHTMLFileWithDelims(filename, left, right string, data interface{}, funcMap ...HTMLFuncMap) (string, error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return RenderHTMLWithDelims(string(s), left, right, data, funcMap...)
}

func MustRenderHTML(tmpl string, data interface{}, funcMap ...HTMLFuncMap) string {
	s, err := RenderHTML(tmpl, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderHTMLMap(tmpl string, m map[string]interface{}) string {
	s, err := RenderHTMLMap(tmpl, m)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderHTMLWithDelims(tmpl, left, right string, data interface{}, funcMap ...HTMLFuncMap) string {
	s, err := RenderHTMLWithDelims(tmpl, left, right, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderHTMLFile(filename string, data interface{}, funcMap ...HTMLFuncMap) string {
	s, err := RenderHTMLFile(filename, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderHTMLFileWithDelims(filename, left, right string, data interface{}, funcMap ...HTMLFuncMap) string {
	s, err := RenderHTMLFileWithDelims(filename, left, right, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}
