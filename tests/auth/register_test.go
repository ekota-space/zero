package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	authDao "github.com/ekota-space/zero/pkgs/auth/dao"
	"github.com/ekota-space/zero/tests/fake"
	"github.com/ekota-space/zero/tests/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func postAuthRegister(router *gin.Engine, body authDao.RegisterDao) *httptest.ResponseRecorder {
	bodyJson, _ := json.Marshal(body)

	w := test.CreateRequest(router, "POST", "/auth/register", strings.NewReader(string(bodyJson)))

	return w
}

func TestPostRegister(t *testing.T) {
	router, db := test.Initialize()

	defer db.Close()

	t.Run("TestPostRegister_Success", func(t *testing.T) {
		body := fake.GenerateRandomRegisterDao()
		w := postAuthRegister(router, body)

		assert.Equal(t, 201, w.Code)
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

	t.Run("TestPostRegister_UsernameExists", func(t *testing.T) {
		body := fake.GenerateRandomRegisterDao()

		postAuthRegister(router, body)

		body.Email = fake.GenerateRandomRegisterDao().Email
		w := postAuthRegister(router, body)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "Username already exists")
	})

	t.Run("TestPostRegister_EmailExists", func(t *testing.T) {
		body := fake.GenerateRandomRegisterDao()
		postAuthRegister(router, body)

		body.Username = fake.GenerateRandomRegisterDao().Username
		w := postAuthRegister(router, body)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "Email already exists")
	})
}
