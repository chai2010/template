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

	// Output:
	// Hello, Neo
}
