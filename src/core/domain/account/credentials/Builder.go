package credentials

import (
	"strings"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"
	"task_manager/src/utils/validator"
)

type builder struct {
	credentials   *Credentials
	invalidFields []errors.InvalidField
}

func NewBuilder() *builder {
	return &builder{credentials: &Credentials{}}
}

func (instance *builder) WithEmail(email string) *builder {
	email = strings.TrimSpace(email)
	if !validator.IsEmailValid(email) {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.AccountEmail,
			Description: messages.InvalidAccountEmailErrorMessage,
		})
	}
	instance.credentials.email = email
	return instance
}

func (instance *builder) WithPassword(password string) *builder {
	if !validator.IsPasswordValid(password) {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.AccountPassword,
			Description: messages.InvalidAccountPasswordErrorMessage,
		})
	}
	instance.credentials.password = password
	return instance
}

func (instance *builder) WithUnvalidatedPassword(password string) *builder {
	instance.credentials.password = password
	return instance
}

func (instance *builder) Build() (*Credentials, errors.Error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.InvalidAccountErrorMessage, instance.invalidFields...)
	}
	return instance.credentials, nil
}
