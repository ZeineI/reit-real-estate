package wallets

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type wallet struct {
	id        uuid.UUID
	userID    uuid.UUID
	address   string
	createdAt time.Time
}

var ErrNotFound = errors.New("wallet not found")
