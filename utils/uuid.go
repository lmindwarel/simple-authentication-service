package utils

import (
	"github.com/lmindwarel/authentificator/models"
	uuid "github.com/satori/go.uuid"
)

// NewUUID return new UUID
func NewUUID() models.UUID {
	return models.UUID(uuid.NewV4().String())
}
