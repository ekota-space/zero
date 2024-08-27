package root

import (
	"fmt"

	"github.com/ekota-space/zero/pkgs/common"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	r.Run(fmt.Sprintf(":%d", common.Env.Port))
}
