// Copyright Improbable Worlds Ltd.. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/mwitkow/go-proto-validators/internal/generator"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				generator.GenerateFile(gen, f)
			}
		}
		return nil
	})
}
