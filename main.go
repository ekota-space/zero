package main

import (
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root"
	"github.com/ekota-space/zero/pkgs/root/db"
)

func main() {
	common.SetupEnvironmentVars()
	db.SetupDatabaseConnection()
	root.SetupRoutes()
}
