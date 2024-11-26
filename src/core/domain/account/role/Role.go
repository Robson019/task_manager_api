package role

import (
	"strings"

	"github.com/google/uuid"
)

type Role struct {
	id   uuid.UUID
	name string
}

func (instance *Role) ID() uuid.UUID {
	return instance.id
}

func (instance *Role) Name() string {
	return strings.ToUpper(instance.name)
}

func (instance *Role) IsZero() bool {
	return instance == &Role{}
}
