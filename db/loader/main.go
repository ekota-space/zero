package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	authModels "github.com/ekota-space/zero/pkgs/auth/models"
	organizationModels "github.com/ekota-space/zero/pkgs/organizations/models"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&authModels.Users{}, &organizationModels.Organizations{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
