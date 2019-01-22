// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package template

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func RenderHTML(tmpl string, data interface{}, funcMap ...template.FuncMap) (string, error) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", err
	}

	for _, fnMap := range funcMap {
		if len(fnMap) > 0 {
			t = t.Funcs(fnMap)
		}
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func RenderHTMLFile(filename string, data interface{}, funcMap ...template.FuncMap) (string, error) {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return RenderHTML(string(s), data, funcMap...)
}

func MustRenderHTML(tmpl string, data interface{}, funcMap ...template.FuncMap) string {
	s, err := RenderHTML(tmpl, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}

func MustRenderHTMLFile(filename string, data interface{}, funcMap ...template.FuncMap) string {
	s, err := RenderHTMLFile(filename, data, funcMap...)
	if err != nil {
		panic(err)
	}
	return s
}
