package domain

import "github.com/google/uuid"

type AdServiceRepository interface {
	Save(ad Ad)
	Find(uuid uuid.UUID) (Ad, bool)
}
