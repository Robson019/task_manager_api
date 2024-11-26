package converters

import (
	"github.com/google/uuid"
	"task_manager/src/core/errors"
)

func ConvertFromStringToUUID(value, fieldName, fieldErrorDescription, errorMessage string) (*uuid.UUID, errors.Error) {
	convertedValue, err := uuid.Parse(value)
	if err != nil {
		return nil, getSingleValidationErr(fieldName, fieldErrorDescription, errorMessage)
	}

	return &convertedValue, nil
}

func getSingleValidationErr(fieldName, fieldErrorDescription, errorMessage string) *errors.ValidationError {
	invalidField := errors.InvalidField{Name: fieldName, Description: fieldErrorDescription}
	return errors.NewValidationError(errorMessage, invalidField)
}
