package root

import (
	"fmt"

	"github.com/ekota-space/zero/pkgs/common"
	root "github.com/ekota-space/zero/pkgs/root/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	r.GET("/", root.GetRoot)

	r.Run(fmt.Sprintf("localhost:%d", common.Env.Port))
}
