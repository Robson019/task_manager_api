package account

import (
	"github.com/google/uuid"
	"task_manager/src/core/domain/account/role"
)

type Account struct {
	id       uuid.UUID
	email    string
	password string
	hash     string

	roleInstance *role.Role
}

func (instance *Account) ID() uuid.UUID {
	return instance.id
}

func (instance *Account) Email() string {
	return instance.email
}

func (instance *Account) Password() string {
	return instance.password
}

func (instance *Account) Hash() string {
	return instance.hash
}

func (instance *Account) Role() *role.Role {
	return instance.roleInstance
}

func (instance *Account) IsZero() bool {
	return instance == &Account{}
}
