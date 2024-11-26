package response

import (
	"github.com/google/uuid"
	"task_manager/src/core/domain/account"
)

type Account struct {
	ID    *uuid.UUID `json:"id"`
	Email string     `json:"email,omitempty"`
	Role  Role       `json:"role"`
}

func NewAccount(account account.Account) *Account {
	id := account.ID()
	return &Account{
		ID:    &id,
		Email: account.Email(),
		Role:  *NewRole(*account.Role()),
	}
}
