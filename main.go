package main

import (
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root"
	"github.com/ekota-space/zero/pkgs/root/db"

	_ "ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	common.SetupEnvironmentVars()
	db.SetupDatabaseConnection()
	root.SetupRoutes()
}
