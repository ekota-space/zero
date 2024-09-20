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
	"github.com/gin-gonic/gin"
)

func Initialize() (*gin.Engine, *sql.DB) {
	gin.SetMode(gin.ReleaseMode)
	common.SetupTestEnvironmentVars()
	db := db.SetupTestDatabaseConnection()

	ql.InitLayer(db)

	return root.SetupRoutes(), db
}

func CreateRequest(router *gin.Engine, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(method, url, body)

	router.ServeHTTP(w, req)

	return w
}
