package auth

import auth "github.com/ekota-space/zero/pkgs/auth/models"

type AuthTokenResponseDao struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         *auth.Users `json:"user,omitempty"`
}
