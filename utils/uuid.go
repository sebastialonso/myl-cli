package utils

import (
	"github.com/gofrs/uuid"
)

func NewUUID4() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
