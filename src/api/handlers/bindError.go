package handlers

import (
	"net/http"
	"task_manager/src/api/handlers/dto/response"
	"task_manager/src/core/errors"

	"github.com/labstack/echo/v4"
)

func treatError(err errors.Error) response.ErrorMessage {
	dto := response.ErrorMessage{}

	if err != nil {
		dto.Message = err.FriendlyMessage()
	}

	switch v := err.(type) {
	case *errors.ConflictError:
		dto.StatusCode = http.StatusConflict
		dto.DuplicatedFields = v.ConflictingFields()
	case *errors.UnexpectedError:
		dto.StatusCode = http.StatusInternalServerError
	case *errors.UnauthorizedError:
		dto.StatusCode = http.StatusUnauthorized
	case *errors.NotAcceptableError:
		dto.StatusCode = http.StatusNotAcceptable
	case *errors.MissingInformationError:
		dto.StatusCode = http.StatusBadRequest
	case *errors.UnavailableServiceError:
		dto.StatusCode = http.StatusServiceUnavailable
	case *errors.NotFoundError:
		dto.StatusCode = http.StatusNotFound
	case *errors.ValidationError:
		dto.StatusCode = http.StatusUnprocessableEntity
		fields := v.InvalidFields()
		for _, f := range fields {
			dto.InvalidFields = append(dto.InvalidFields, response.InvalidField{
				FieldName:   f.Name,
				Description: f.Description,
			})
		}
	}

	return dto
}

func getDefaultBadRequestResponse(context echo.Context, fields ...response.InvalidField) error {
	return context.JSON(http.StatusBadRequest, response.ErrorMessage{
		StatusCode:    http.StatusBadRequest,
		Message:       "Não foi possível processar a solicitação",
		InvalidFields: fields,
	})
}

func getHttpHandledErrorResponse(context echo.Context, err errors.Error) error {
	handledError := treatError(err)
	return context.JSON(handledError.StatusCode, handledError)
}
