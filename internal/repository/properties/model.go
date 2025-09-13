package properties

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type property struct {
	id         uuid.UUID
	ownerID    uuid.UUID
	name       string
	tokenTotal int64
	createdAt  time.Time
}

var ErrNotFound = errors.New("property not found")
