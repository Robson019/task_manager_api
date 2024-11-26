package response

import (
	"task_manager/src/core/domain/authorization"
)

type Authorization struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthorization(authorization authorization.Authorization) *Authorization {
	return &Authorization{
		Token:        authorization.Token(),
		RefreshToken: authorization.RefreshToken(),
	}
}
