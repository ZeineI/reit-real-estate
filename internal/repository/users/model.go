package users

import (
	"github.com/google/uuid"
	"time"
)

type user struct {
	id        uuid.UUID
	login     string
	role      string
	createdAt time.Time
}
