package auth

import (
	"testing"

	"github.com/ekota-space/zero/tests/test"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(t *testing.M) {
	r, db := test.Initialize()
	defer db.Close()

	router = r

	t.Run()
}
