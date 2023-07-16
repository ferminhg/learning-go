package application

import (
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
)

const MaxSearch = 5

type AdService struct {
	repository domain.AdServiceRepository
}

func (service AdService) post(title string, description string, price float32) domain.Ad {
	ad, _ := domain.NewAd(title, description, price)
	service.repository.Save(ad)
	return ad
}

func (service AdService) find(adId string) (domain.Ad, bool) {
	uuidAdId, err := uuid.Parse(adId)
	if err != nil {
		return domain.Ad{}, false
	}

	return service.repository.Find(uuidAdId)
}
