// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package template_test

import (
	"fmt"

	"github.com/chai2010/template"
)

func Example() {
	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, map[string]string{
			"Name": "Neo",
		}),
	)
	fmt.Println(
		template.MustRender(`Hello, {{.Name}}`, struct{Name string}{
			Name: "chai2010",
		}),
	)

	// Output:
	// Hello, Neo
	// Hello, chai2010
}
