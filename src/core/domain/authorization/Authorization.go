package authorization

import (
	"encoding/base64"
	"os"
	"task_manager/src/core"
	"task_manager/src/core/domain/account"
	"task_manager/src/core/errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var logger = core.Logger()

const (
	RefreshTokenTimeout = time.Hour
	AccessTokenTimeout  = time.Minute * 15
	AnonymousRoleCode   = "anonymous"
)

type Authorization interface {
	Token() string
	RefreshToken() string
}

type authorization struct {
	token        string
	refreshToken string
}

func New() Authorization {
	return &authorization{}
}

func NewFromAccount(acc account.Account) (Authorization, errors.Error) {
	instance := &authorization{}
	if err := instance.GenerateToken(acc); err != nil {
		return nil, err
	}
	return instance, nil
}

func (instance *authorization) Token() string {
	return instance.token
}

func (instance *authorization) RefreshToken() string {
	return instance.refreshToken
}

func (instance *authorization) GenerateToken(account account.Account) errors.Error {
	secret := os.Getenv("SERVER_SECRET")
	privateKey := os.Getenv("SERVER_REFRESH_TOKEN_PRIVATE_KEY")

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpectedError(err.Error(), err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpectedError(err.Error(), err)
	}

	accesstoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(
		account.ID().String(),
		account.Role().Name(),
		"bearer",
		time.Now().Add(AccessTokenTimeout).Unix(),
	)).SignedString([]byte(secret))
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpectedError(err.Error(), err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, newRefreshClaims(
		account.ID().String(),
		"bearer",
		time.Now().Add(RefreshTokenTimeout).Unix(),
	)).SignedString(signKey)
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpectedError(err.Error(), err)
	}

	instance.token = accesstoken
	instance.refreshToken = refreshToken
	return nil
}
