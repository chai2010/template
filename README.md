# tmplate helper

[![Build Status](https://travis-ci.org/chai2010/template.svg)](https://travis-ci.org/chai2010/template)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/template)](https://goreportcard.com/report/github.com/chai2010/template)
[![GoDoc](https://godoc.org/github.com/chai2010/template?status.svg)](https://godoc.org/github.com/chai2010/template)
[![License](http://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/chai2010/template/blob/master/LICENSE)


## Example

```go
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

	// Output:
	// Hello, Neo
	// Hello, chai2010
	// Hello, CHAI2010
	// 《Go语言高级编程》《WebAssembly标准入门》《C/C++面向WebAssembly编程》
}
```

## BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
