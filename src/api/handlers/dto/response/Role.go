package response

import (
	"github.com/google/uuid"
	"task_manager/src/core/domain/account/role"
)

type Role struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

func NewRole(role role.Role) *Role {
	roleResponse := &Role{
		ID:   role.ID(),
		Name: role.Name(),
	}

	return roleResponse
}
