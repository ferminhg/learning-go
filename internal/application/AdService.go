package application

import (
	"errors"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
	"math/rand"
)

type AdService struct {
	Repository domain.AdServiceRepository
	generator  domain.DescriptionGenerator
}

func NewAdService(
	repository domain.AdServiceRepository,
	generator domain.DescriptionGenerator,
) AdService {
	return AdService{
		Repository: repository,
		generator:  generator,
	}
}

func (s AdService) Post(title string, description string, price float32) (domain.Ad, error) {
	ad, err := domain.NewAd(title, description, price)
	if err != nil {
		return domain.Ad{}, err
	}

	err = s.Repository.Save(ad)
	return ad, err
}

func (s AdService) Find(adId string) (domain.Ad, bool) {
	uuidAdId, err := uuid.Parse(adId)
	if err != nil {
		return domain.Ad{}, false
	}

	return s.Repository.Find(uuidAdId)
}

func (s AdService) FindRandom() ([]domain.Ad, error) {
	return s.Repository.Search(5)
}

func (s AdService) DescriptionGenerator(title string) (domain.RandomDescription, error) {
	descriptions, err := s.generator.Run(title)
	if err != nil {
		return domain.RandomDescription{}, err
	}

	switch len(descriptions) {
	case 0:
		return domain.RandomDescription{}, errors.New("timeout generating descriptions")
	case 1:
		return descriptions[0], nil
	default:
		r := rand.Intn(len(descriptions) - 1)
		return descriptions[r], nil

	}
}
