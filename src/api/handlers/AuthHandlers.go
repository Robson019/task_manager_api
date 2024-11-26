package handlers

import (
	"net/http"
	"task_manager/src/api/handlers/dto/request"
	"task_manager/src/api/handlers/dto/response"
	"task_manager/src/api/handlers/utils/converters"
	"task_manager/src/api/handlers/utils/token"
	"task_manager/src/core/domain/account/credentials"
	"task_manager/src/core/errors"
	"task_manager/src/core/interfaces/primary"
	"task_manager/src/core/messages"
	"task_manager/src/core/utils"

	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	service primary.AuthManager
}

// Login
// @ID Login
// @Summary Fazer login no sistema
// @Tags Rotas de autenticação
// @Description Rota que permite que um usuário se autentique no sistema utilizando seu endereço de e-mail e senha.
// @Description | E-mail              | Senha     | Função                                                            |
// @Description |---------------------|-----------|-------------------------------------------------------------------|
// @Description | robson@gmail.com | Test1234! | Usuário do sistema. |
// @Accept json
// @Produce json
// @Param json body request.LoginDTO true "JSON com todos os dados necessários para que o login seja realizado."
// @Success 201 {object} response.Authorization "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /auth/login [post]
func (instance AuthHandlers) Login(context echo.Context) error {
	var dto request.LoginDTO
	bindError := context.Bind(&dto)
	if bindError != nil {
		return getDefaultBadRequestResponse(context)
	}

	builder := credentials.NewBuilder()
	builder.WithEmail(dto.Email).WithUnvalidatedPassword(dto.Password)
	builtCredentials, buildError := builder.Build()
	if buildError != nil {
		return getHttpHandledErrorResponse(context, buildError)
	}

	authorization, loginError := instance.service.Login(*builtCredentials)
	if loginError != nil {
		return getHttpHandledErrorResponse(context, loginError)
	}

	return context.JSON(http.StatusCreated, response.NewAuthorization(authorization))
}

// Refresh
// @ID Refresh
// @Summary Gerar um novo par de tokens para autenticação
// @Tags Rotas de autenticação
// @Accept json
// @Produce json
// @Param json body request.RefreshDTO true "JSON com todos os dados necessários para que o login seja realizado."
// @Success 201 {object} response.Authorization "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /auth/refresh [post]
func (instance AuthHandlers) Refresh(context echo.Context) error {
	var dto request.RefreshDTO
	bindError := context.Bind(&dto)
	if bindError != nil {
		return getDefaultBadRequestResponse(context)
	}

	err := utils.ValidateRefreshToken(dto.RefreshToken)
	if err != nil {
		return getHttpHandledErrorResponse(context, err)
	}

	claims, extractErr := utils.ExtractTokenClaims(dto.RefreshToken)
	if extractErr != nil {
		newErr := errors.NewUnauthorizedError(messages.InvalidRefreshTokenErrorMessage, extractErr)
		return getHttpHandledErrorResponse(context, newErr)
	}

	accountID, err := converters.ConvertFromStringToUUID(claims.AccountID, messages.AccountID, messages.InvalidAccountIDErrorMessage, messages.ConversionErrorMessage)
	if err != nil {
		return getHttpHandledErrorResponse(context, err)
	}

	authorization, err := instance.service.Refresh(dto.RefreshToken, *accountID)
	if err != nil {
		return getHttpHandledErrorResponse(context, err)
	}

	return context.JSON(http.StatusCreated, response.NewAuthorization(authorization))
}

// Logout
// @ID Logout
// @Summary Fazer logout no sistema
// @Tags Rotas de autenticação
// @Description Rota que permite que um usuário faça logout no sistema.
// @Security bearerAuth
// @Produce json
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /auth/logout [delete]
func (instance AuthHandlers) Logout(context echo.Context) error {
	accountID, authError := token.GetAccountIDFromAuthorization(context)
	if authError != nil {
		return getHttpHandledErrorResponse(context, authError)
	}

	err := instance.service.Logout(*accountID)
	if err != nil {
		return getHttpHandledErrorResponse(context, err)
	}

	return context.NoContent(http.StatusNoContent)
}

func NewAuthHandlers(service primary.AuthManager) *AuthHandlers {
	return &AuthHandlers{service: service}
}
