package users

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type user struct {
	id        uuid.UUID
	login     string
	role      string
	createdAt time.Time
}

var ErrNotFound = errors.New("user not found")
