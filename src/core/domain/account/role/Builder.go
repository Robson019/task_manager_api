package role

import (
	"strings"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"
	"task_manager/src/utils/validator"

	"github.com/google/uuid"
)

type builder struct {
	role          *Role
	invalidFields []errors.InvalidField
}

func NewBuilder() *builder {
	return &builder{role: &Role{}}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.RoleID,
			Description: messages.InvalidRoleIDErrorMessage,
		})
		return instance
	}
	instance.role.id = id
	return instance
}

func (instance *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        messages.RoleName,
			Description: messages.InvalidRoleNameErrorMessage,
		})
		return instance
	}
	instance.role.name = name
	return instance
}

func (instance *builder) Build() (*Role, errors.Error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.InvalidRoleErrorMessage, instance.invalidFields...)
	}
	return instance.role, nil
}
