package infra

import (
	"fmt"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
)

type InMemoryAdRepository struct {
	ads map[string]domain.Ad
}

func (repository InMemoryAdRepository) Search(maxNumber int) ([]domain.Ad, error) {
	panic("implement me")
}

func NewInMemoryAdRepository() *InMemoryAdRepository {
	return &InMemoryAdRepository{
		ads: make(map[string]domain.Ad),
	}
}

func (repository InMemoryAdRepository) Save(ad domain.Ad) {
	fmt.Println("💾 saving ...", ad)
	repository.ads[ad.Id.String()] = ad
}

func (repository InMemoryAdRepository) Find(uuid uuid.UUID) (domain.Ad, bool) {
	fmt.Println("🔎 finding ...", uuid.String())
	val, ok := repository.ads[uuid.String()]
	return val, ok
}
