package main

import (
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root"
)

func main() {
	common.SetupEnvironmentVars()

	root.SetupRoutes()
}
