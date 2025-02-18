package mockData

import (
	"github.com/google/uuid"
	"task_manager/src/utils/tests/functions"
	"time"
)

// Common
var (
	Token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjIjpudWxsLCJzdWIiOiI2NjljZGZlYi05YmUzLTRjZGQtYTc0OS0zZjcwNGNkYWE0NjkiLCJzZWN0aW9uIjoiUk9MRV9QUk9GSVNTSU9OQUxfVU5JREFERV9TQVVERSIsImV4cCI6MTY3NTQ2MTMxMywidHlwIjoiYmVhcmVyIn0.nyqyxs1tAveZAMtubZ_iXemhRjKA2PV5V9iykNT9lb4"

	TheStatusCodesDoesNotMatch = "The status codes does not match"
	TheJSONsDoesNotMatch       = "The JSONs does not match"

	ConnectionError = "connection error"
)

// Account
var (
	AccountID    = uuid.MustParse("669cdfeb-9be3-4cdd-a749-3f704cdaa469")
	AccountEmail = "robson@gmail.com"
)

// Role
var (
	RoleID   = uuid.MustParse("5c2b8986-050b-4897-890a-3675dee732ae")
	RoleName = "ROLE_USUARIO"
)

// Task
var (
	TaskID           = uuid.MustParse("2f2f9da3-a08a-4ade-b07a-d63b53508a54")
	TaskTitle        = "Estudar Golang"
	TaskDescription  = "Revisar fundamentos do Golang"
	TaskStatus       = "pending"
	TaskCreatedAt, _ = time.Parse("2006-01-02", "2024-11-25")
	TaskUpdatedAt    = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

	TaskValuesValid = functions.ReadJSON("register-valid-task.json")
)
