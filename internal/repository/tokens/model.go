package tokens

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type token struct {
	id         uuid.UUID
	propertyID uuid.UUID
	symbol     string
	price      float64
	createdAt  time.Time
}

var ErrNotFound = errors.New("token not found")
