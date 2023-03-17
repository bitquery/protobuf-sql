package main

import (
	"github.com/bitquery/protobuf-sql/compiler/protogen"
	"github.com/bitquery/protobuf-sql/generators"
)

func main() {

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			// if f.GeneratedFilenamePrefix != "tables" {
			// 	continue
			// }

			for _, m := range f.Messages {
				generators.GenerateFileForMessage(gen, m)
			}
		}
		return nil
	})
}
