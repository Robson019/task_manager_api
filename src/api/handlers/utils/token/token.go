package token

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"task_manager/src/core/domain/authorization"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"
	"task_manager/src/core/utils"
)

func getAuthClaims(authHeader string) (*authorization.Claims, errors.Error) {
	_, token := utils.ExtractToken(authHeader)
	authClaims, err := utils.ExtractTokenClaims(token)
	if err != nil {
		invalidFields := []errors.InvalidField{{Name: messages.Token, Description: messages.InvalidAccessTokenErrorMessage}}
		return nil, errors.NewValidationError(messages.InvalidTokenErrorMessage, invalidFields...)
	}
	return authClaims, nil
}

func GetAccountIDFromAuthorization(ctx echo.Context) (*uuid.UUID, errors.Error) {
	claims, err := getAuthClaims(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	if accountID, parseError := uuid.Parse(claims.AccountID); parseError != nil {
		invalidFields := []errors.InvalidField{{Name: messages.IDFromToken, Description: messages.InvalidIDFromAccessTokenErrorMessage}}
		return nil, errors.NewValidationError(messages.InvalidIDFromTokenErrorMessage, invalidFields...)
	} else {
		return &accountID, nil
	}
}
