package inmemory

import (
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type InMemoryAdRepository struct {
	ads map[string]domain.Ad
}

func NewInMemoryAdRepository() *InMemoryAdRepository {
	return &InMemoryAdRepository{
		ads: make(map[string]domain.Ad),
	}
}

func (repository InMemoryAdRepository) Search(maxNumber int) ([]domain.Ad, error) {
	var ads []domain.Ad

	if len(repository.ads) <= maxNumber {
		for _, value := range repository.ads {
			ads = append(ads, value)
		}
		return ads, nil
	}

	keys := make([]string, 0, len(repository.ads))
	for k := range repository.ads {
		keys = append(keys, k)
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < maxNumber && len(keys) > 0; i++ {
		randomIndex := r.Intn(len(keys))
		ads = append(ads, repository.ads[keys[randomIndex]])
		keys = append(keys[:randomIndex], keys[randomIndex+1:]...)
	}

	//fmt.Println(ads)
	return ads, nil
}

func (repository InMemoryAdRepository) Save(ad domain.Ad) error {
	//fmt.Println("💾 saving ...", ad)
	repository.ads[ad.Id.String()] = ad
	return nil
}

func (repository InMemoryAdRepository) Find(uuid uuid.UUID) (domain.Ad, bool) {
	//fmt.Println("🔎 finding ...", uuid.String())
	val, ok := repository.ads[uuid.String()]
	return val, ok
}
