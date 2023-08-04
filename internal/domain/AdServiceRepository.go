package domain

import "github.com/google/uuid"

type AdServiceRepository interface {
	Save(ad Ad) error
	Find(uuid uuid.UUID) (Ad, bool)
	Search(maxNumber int) ([]Ad, error)
	Delete(id uuid.UUID) bool
}

//go:generate mockery --case=snake --outpkg=storagemocks --name=AdServiceRepository --output=../infra/storage/storagemocks
