// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"strings"

	"github.com/chai2010/template"
)

func main() {
	fmt.Println(
		template.MustRender(`Hello, {{.}}`, "Neo"),
	)
	fmt.Println(
		template.MustRender(`Hello, {{index . 0}}`, []string{"Go"}),
	)
	fmt.Println(
		template.MustRender(`Hello, {{index . "Name"}}`,
			map[string]string{"Name": "凹(Wa)"},
		),
	)

	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, map[string]string{
			"Name": "Lua",
		}),
	)
	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, struct{ Name string }{
			Name: "Ruby",
		}),
	)

	fmt.Println(
		template.MustRender(
			`Hello, {{upper .Name}}`, struct{ Name string }{Name: "chai2010"},
			template.FuncMap{"upper": strings.ToUpper},
		),
	)

	fmt.Println(
		template.MustRender(
			`{{range $i, $v := .}}{{$v.Book}}{{end}}`,
			[]struct{ Name, Book string }{
				{Name: "chai2010", Book: "《Go语言高级编程》"},
				{Name: "chai2010 & ending", Book: "《WebAssembly标准入门》"},
				{Name: "ending & chai2010", Book: "《C/C++面向WebAssembly编程》"},
			},
		),
	)

	fmt.Println(
		template.MustRenderWithDelims(`Hello, {{<<.))}}`, `<<`, `))`, "Neo"),
	)

	fmt.Println(
		template.MustRender(`Hello, {{A}}-{{B}}-{{C}}`, "",
			template.FuncMap{
				"A": func() bool { return true },
				"B": func() int { return 9527 },
				"C": func() string { return "C-Value" },
			},
		),
	)

	// Output:
	// Hello, Neo
	// Hello, Go
	// Hello, 凹(Wa)
	// Hello, Lua
	// Hello, Ruby
	// Hello, CHAI2010
	// 《Go语言高级编程》《WebAssembly标准入门》《C/C++面向WebAssembly编程》
	// Hello, {{Neo}}
	// Hello, true-9527-C-Value
}
