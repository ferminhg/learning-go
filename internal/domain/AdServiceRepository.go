package domain

import "github.com/google/uuid"

type AdServiceRepository interface {
	Save(ad Ad) error
	Find(uuid uuid.UUID) (Ad, bool)
	Search(maxNumber int) ([]Ad, error)
}
