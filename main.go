package main

import (
	"fmt"

	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/ql"
)

func main() {
	common.SetupEnvironmentVars()
	db := db.SetupDatabaseConnection()

	ql.InitLayer(db)

	router := root.SetupRoutes()

	router.Run(fmt.Sprintf("localhost:%d", common.Env.Port))
}
