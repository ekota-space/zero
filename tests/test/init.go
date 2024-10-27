package test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/ekota-space/zero/pkgs/common"
	"github.com/ekota-space/zero/pkgs/root"
	"github.com/ekota-space/zero/pkgs/root/db"
	"github.com/ekota-space/zero/pkgs/root/ql"
	"github.com/gofiber/fiber/v2"
)

func Initialize() (*fiber.App, *sql.DB) {
	common.SetupTestEnvironmentVars()
	db := db.SetupDatabaseConnection()

	ql.InitLayer(db)

	return root.SetupRoutes(), db
}

func CreateRequest(router *fiber.App, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(method, url, body)

	// router.ServeHTTP(w, req)

	return w
}
