package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
	"math/rand"
)

type AdService struct {
	repository   domain.AdServiceRepository
	generator    domain.DescriptionGenerator
	eventHandler domain.EventHandler
}

func NewAdService(
	repository domain.AdServiceRepository,
	generator domain.DescriptionGenerator,
	eventHandler domain.EventHandler,
) AdService {
	return AdService{
		repository:   repository,
		generator:    generator,
		eventHandler: eventHandler,
	}
}

func (s AdService) Post(title string, description string, price float32) (domain.Ad, error) {
	ad, err := domain.NewAd(title, description, price)
	if err != nil {
		return domain.Ad{}, err
	}

	if err := s.repository.Save(ad); err != nil {
		return domain.Ad{}, err
	}

	if err := s.sendAdEvent(ad); err != nil {
		return domain.Ad{}, fmt.Errorf("failed to store your data: %s", err)
	}

	return ad, nil
}

func (s AdService) Find(adId string) (domain.Ad, bool) {
	uuidAdId, err := uuid.Parse(adId)
	if err != nil {
		return domain.Ad{}, false
	}

	return s.repository.Find(uuidAdId)
}

func (s AdService) FindRandom() ([]domain.Ad, error) {
	return s.repository.Search(5)
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

func (s AdService) sendAdEvent(ad domain.Ad) error {
	jsonAd, _ := json.Marshal(ad)
	partition, offset, err := s.eventHandler.SendMessage(domain.NewProducerMessage(domain.AdTopic, string(jsonAd)))
	if err != nil {
		return fmt.Errorf("failed to store your data: %s", err)
	}

	fmt.Printf("Your data is stored with unique identifier important/%d/%d\n", partition, offset)
	return nil
}

func (s AdService) Delete(adId string) (bool, error) {
	uuidAdId, err := uuid.Parse(adId)
	if err != nil {
		return false, err
	}

	ok := s.repository.Delete(uuidAdId)

	return ok, nil
}
