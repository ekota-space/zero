package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	authModels "github.com/ekota-space/zero/pkgs/auth/models"
)

func main() {
	statements, err := gormschema.New("postgres").Load(&authModels.Users{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, statements)
}
