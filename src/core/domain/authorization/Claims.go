package authorization

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.Claims `json:"c,omitempty"`
	AccountID  string `json:"sub"`
	RoleCode   string `json:"section"`
	Expiry     int64  `json:"exp"`
	Type       string `json:"typ"`
}

func newClaims(accountID string, roleCode string, typ string, exp int64) *Claims {
	return &Claims{
		AccountID: accountID,
		RoleCode:  roleCode,
		Type:      typ,
		Expiry:    exp,
	}
}

func newRefreshClaims(accountID string, typ string, exp int64) *Claims {
	return &Claims{
		AccountID: accountID,
		Type:      typ,
		Expiry:    exp,
	}
}
