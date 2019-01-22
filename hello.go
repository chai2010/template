// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/template"
)

func main() {
	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, map[string]string{
			"Name": "Neo",
		}),
	)
	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, struct{ Name string }{
			Name: "chai2010",
		}),
	)

	fmt.Println(
		template.MustRender(
			`{{range $i, $v := .}}{{$v.Book}}{{end}}`,
			[]struct{ Name, Book string }{
				{
					Name: "chai2010",
					Book: "《Go语言高级编程》",
				},
				{
					Name: "chai2010",
					Book: "《WebAssembly标准入门》",
				},
				{
					Name: "ending",
					Book: "《C/C++面向WebAssembly编程》",
				},
			},
		),
	)

	// Output:
	// Hello, Neo
	// Hello, chai2010
	// 《Go语言高级编程》《WebAssembly标准入门》《C/C++面向WebAssembly编程》
}
