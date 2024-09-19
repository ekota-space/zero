package main

import (
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/ql"
)

func main() {
	common.SetupEnvironmentVars()
	db := db.SetupDatabaseConnection()

	ql.InitLayer(db)

	root.SetupRoutes()
}
