package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	Id          uuid.UUID
	Title       string
	Description string
	Price       float32
	CreatedDate time.Time
}
