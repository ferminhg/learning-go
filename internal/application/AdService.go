package application

import (
	"github.com/ferminhg/learning-go/internal/domain"
)

type AdService struct {
	repository domain.AdServiceRepository
}

func (service AdService) post(title string, description string, price float32) domain.Ad {
	ad, _ := domain.NewAd(title, description, price)
	service.repository.Save(ad)
	return ad
}
