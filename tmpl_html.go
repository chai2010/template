// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"html/template"
	"io/ioutil"
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

func RenderHTMLFile(filename string, data interface{}, funcMap ...HTMLFuncMap) (string, error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return RenderHTML(string(s), data, funcMap...)
}

func MustRenderHTML(tmpl string, data interface{}, funcMap ...HTMLFuncMap) string {
	s, err := RenderHTML(tmpl, data, funcMap...)
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
