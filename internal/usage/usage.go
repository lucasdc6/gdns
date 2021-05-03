// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package usage implements the function used by the help
// module
package usage

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lucasdc6/gdns/pkg/errors"
)

// Usage - Print help usage for the differents modules
func Usage(module string) {
	switch module {
	case "file-syntax":
		file, err := ioutil.ReadFile("./docs/configuration.md")

		if err != nil {
			fmt.Printf("Error reading documentation for %s", module)
			os.Exit(errors.ModuleDocumentationNotFound)
		}

		fmt.Print(string(file))
		return
	default:
		fmt.Printf("Error module %s not found", module)
		os.Exit(errors.ModuleNotFound)
	}
}
