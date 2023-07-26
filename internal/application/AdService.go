package application

import (
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
)

type AdService struct {
	Repository domain.AdServiceRepository
}

func NewAdService(repository domain.AdServiceRepository) AdService {
	return AdService{
		Repository: repository,
	}
}

func (service AdService) Post(title string, description string, price float32) (domain.Ad, error) {
	ad, err := domain.NewAd(title, description, price)
	if err != nil {
		return domain.Ad{}, err
	}

	err = service.Repository.Save(ad)
	return ad, err
}

func (service AdService) Find(adId string) (domain.Ad, bool) {
	uuidAdId, err := uuid.Parse(adId)
	if err != nil {
		return domain.Ad{}, false
	}

	return service.Repository.Find(uuidAdId)
}

func (service AdService) FindRandom() ([]domain.Ad, error) {
	return service.Repository.Search(5)
}
