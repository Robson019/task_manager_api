package account

import (
	"strings"
	"task_manager/src/core/domain/account/role"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"
	"task_manager/src/utils/validator"

	"github.com/google/uuid"
)

type builder struct {
	account       *Account
	invalidFields []errors.InvalidField
}

func NewBuilder() *builder {
	return &builder{account: &Account{}}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.AccountID,
			Description: messages.InvalidAccountIDErrorMessage,
		})
		return instance
	}
	instance.account.id = id
	return instance
}

func (instance *builder) WithEmail(email string) *builder {
	email = strings.TrimSpace(email)
	if !validator.IsEmailValid(email) {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.AccountEmail,
			Description: messages.InvalidAccountEmailErrorMessage,
		})
		return instance
	}
	instance.account.email = email
	return instance
}

func (instance *builder) WithPassword(password string) *builder {
	instance.account.password = password
	return instance
}

func (instance *builder) WithHash(hash string) *builder {
	instance.account.hash = hash
	return instance
}

func (instance *builder) WithRole(role role.Role) *builder {
	if role.IsZero() {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.Role,
			Description: messages.InvalidRoleErrorMessage,
		})
		return instance
	}
	instance.account.roleInstance = &role
	return instance
}

func (instance *builder) Build() (*Account, errors.Error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.InvalidAccountErrorMessage, instance.invalidFields...)
	}
	return instance.account, nil
}
