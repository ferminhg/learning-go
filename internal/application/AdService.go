package application

import (
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
)

type AdService struct {
	Repository domain.AdServiceRepository
}

func (service AdService) Post(title string, description string, price float32) domain.Ad {
	ad, _ := domain.NewAd(title, description, price)
	service.Repository.Save(ad)
	return ad
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