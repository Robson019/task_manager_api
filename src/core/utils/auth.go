package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"task_manager/src/core"
	"task_manager/src/core/domain/authorization"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"

	"github.com/golang-jwt/jwt"
)

var logger = core.Logger()

func ExtractAuthorizationAccountRole(authHeader string) (string, bool) {
	authType, authToken := ExtractToken(authHeader)
	if authType == "" || authToken == "" {
		return authorization.AnonymousRoleCode, true
	} else if claims, ok := authorizationIsValid(authType, authToken); !ok {
		return authorization.AnonymousRoleCode, false
	} else {
		return claims.RoleCode, true
	}
}

func ValidateRefreshToken(token string) errors.Error {
	publicKey := os.Getenv("SERVER_REFRESH_TOKEN_PUBLIC_KEY")
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.NewMissingInformationError(messages.InvalidRefreshTokenErrorMessage)
	}

	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], verifyKey)
	if err != nil {
		return errors.NewUnauthorizedError(messages.InvalidRefreshTokenErrorMessage, err)
	}
	return nil
}

func ExtractToken(authHeader string) (authType string, token string) {
	authorization := strings.Split(strings.TrimSpace(authHeader), " ")
	if len(authorization) < 2 {
		return "", ""
	}
	authType = authorization[0]
	token = authorization[1]
	return authType, token
}

func authorizationIsValid(authType, authToken string) (*authorization.Claims, bool) {
	secret := os.Getenv("SERVER_SECRET")
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		logger.Error().Msg("error parsing the provided token on (signature is invalid?)")
		return nil, false
	}
	if !token.Valid || token.Claims.Valid() != nil {
		logger.Error().Msg("the provided token is invalid or expired")
		return nil, false
	}
	claims, err := ExtractTokenClaims(authToken)
	if err != nil {
		return nil, false
	}
	if strings.ToLower(claims.Type) != strings.ToLower(authType) {
		logger.Error().Msg(fmt.Sprintf("the used authorization type \"%s\" is not supported", authType))
		return nil, false
	}
	return claims, true
}

func ExtractTokenClaims(authToken string) (*authorization.Claims, error) {
	parts := strings.Split(authToken, ".")
	payload := parts[1]
	payloadBytes, err := jwt.DecodeSegment(payload)
	if err != nil {
		logger.Error().Msg("an error occurred when decoding the token payload: " + err.Error())
		return nil, err
	}
	var claims authorization.Claims
	err = json.Unmarshal(payloadBytes, &claims)
	if err != nil {
		logger.Error().Msg("an error occurred when unmarshalling the token payload: " + err.Error())
		return nil, err
	}
	return &claims, nil
}
