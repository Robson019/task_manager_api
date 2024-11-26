package handlers

import (
	"net/http"
	"task_manager/src/api/handlers/dto/response"
	"task_manager/src/api/handlers/utils/token"
	"task_manager/src/core/interfaces/primary"

	"github.com/labstack/echo/v4"
)

type AccountHandlers struct {
	service primary.AccountManager
}

// FindProfile
// @ID FindProfile
// @Summary Pesquisar dados do perfil do usuário logado
// @Description Rota que retorna todas as informações do usuário logado.
// @Security bearerAuth
// @Tags Conta do usuário
// @Produce json
// @Success 200 {object} response.Account "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /account/profile [get]
func (instance AccountHandlers) FindProfile(context echo.Context) error {
	accountID, authError := token.GetAccountIDFromAuthorization(context)
	if authError != nil {
		return getHttpHandledErrorResponse(context, authError)
	}
	accountInstance, fetchError := instance.service.FetchProfileByID(*accountID)
	if fetchError != nil {
		return getHttpHandledErrorResponse(context, fetchError)
	}
	return context.JSON(http.StatusOK, response.NewAccount(*accountInstance))
}

func NewAccountHandlers(service primary.AccountManager) *AccountHandlers {
	return &AccountHandlers{service: service}
}
