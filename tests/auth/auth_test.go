package auth

import (
	"testing"

	"github.com/ekota-space/zero/tests/setup"
	"github.com/ekota-space/zero/tests/test"
	"github.com/gofiber/fiber/v2"
)

var router *fiber.App

func TestMain(t *testing.M) {
	setup.GlobalSetup()
	r, db := test.Initialize()
	defer db.Close()

	router = r

	t.Run()
}
