package main

import (
	"fmt"

	"github.com/ekota-space/zero/docs"
	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/ql"
)

// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	docs.SwaggerInfo.Title = "Zero API"
	docs.SwaggerInfo.Description = "API controller for Ekota"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	common.SetupEnvironmentVars()
	db := db.SetupDatabaseConnection()

	ql.InitLayer(db)

	router := root.SetupRoutes()

	router.Run(fmt.Sprintf("localhost:%d", common.Env.Port))
}
