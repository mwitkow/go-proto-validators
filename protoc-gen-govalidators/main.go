// Copyright Improbable Worlds Ltd.. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"google.golang.org/protobuf/compiler/protogen"

<<<<<<< HEAD
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	validator_plugin "github.com/maanasasubrahmanyam-sd/go-proto-validators/plugin"
=======
	"github.com/mwitkow/go-proto-validators/internal/generator"
>>>>>>> origin/protobuf-v2
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
