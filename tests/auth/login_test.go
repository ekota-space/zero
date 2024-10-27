package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ekota-space/zero/pkgs/auth"
	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/model"
	"github.com/ekota-space/zero/pkgs/root/db/zero/public/table"
	"github.com/ekota-space/zero/pkgs/root/ql"
	"github.com/ekota-space/zero/tests/fake"
	"github.com/ekota-space/zero/tests/test"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func postAuthLogin(router *fiber.App, body authDao.LoginDao) *httptest.ResponseRecorder {
	bodyJson, _ := json.Marshal(body)

	w := test.CreateRequest(router, "POST", "/auth/login", strings.NewReader(string(bodyJson)))

	return w
}

func TestPostLogin(t *testing.T) {
	body := fake.GenerateRandomRegisterDao()
	password, err := auth.HashPassword(body.Password)

	if err != nil {
		t.Error(err)
	}
	rawPassword := body.Password
	body.Password = password

	user := model.Users{}

	stmt := table.Users.INSERT(
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
	).
		MODEL(body).
		RETURNING(table.Users.AllColumns)

	err = stmt.Query(ql.GetDB(), &user)

	if err != nil {
		t.Error(err)
	}

	t.Run("TestPostRegister_Success", func(t *testing.T) {
		w := postAuthLogin(router, authDao.LoginDao{
			Email:    body.Email,
			Password: rawPassword,
		})

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "expirationDurationSeconds")

		cookie := w.Result().Cookies()
		assert.GreaterOrEqual(t, 2, len(cookie))

		var acc_t *http.Cookie
		var ref_t *http.Cookie

		for _, c := range cookie {
			if c.Name == "acc_t" {
				acc_t = c
			} else if c.Name == "ref_t" {
				ref_t = c
			}
		}

		assert.NotNil(t, acc_t)
		assert.NotNil(t, ref_t)
	})

	t.Run("TestPostRegister_InvalidEmail", func(t *testing.T) {
		w := postAuthLogin(router, authDao.LoginDao{
			Email:    "test@t.com",
			Password: rawPassword,
		})

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid email or password")
	})

	t.Run("TestPostRegister_InvalidPassword", func(t *testing.T) {
		w := postAuthLogin(router, authDao.LoginDao{
			Email:    body.Email,
			Password: "wrongpassword",
		})

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid email or password")
	})
}
