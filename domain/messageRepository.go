package domain

import (
	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(id uuid.UUID, from uuid.UUID, to uuid.UUID, message string) error
}
