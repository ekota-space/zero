package organizations

import (
	"testing"

	"github.com/ekota-space/zero/tests/setup"
	"github.com/ekota-space/zero/tests/test"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(t *testing.M) {
	setup.GlobalSetup()
	r, db := test.Initialize()
	defer db.Close()

	router = r

	t.Run()
}
