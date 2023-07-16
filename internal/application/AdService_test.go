package application

import (
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/ferminhg/learning-go/internal/infra"
	"reflect"
	"testing"
)

func TestPostAd(t *testing.T) {
	tests := map[string]struct {
		title       string
		description string
		price       float32
	}{
		"post simple ad": {
			title:       "title",
			description: "A description",
			price:       1.23,
		},
	}
	for name, tt := range tests {
		inMemoryAdRepository := infra.NewInMemoryAdRepository()
		service := AdService{Repository: inMemoryAdRepository}
		t.Run(name, func(t *testing.T) {
			ad, _ := service.Post(tt.title, tt.description, tt.price)
			if ad.Title != tt.title {
				t.Errorf("Expected %s -> got %s", tt.title, ad.Title)
			}
			if ad.Description != tt.description {
				t.Errorf("Expected %s -> got %s", tt.title, ad.Title)
			}
			if ad.Price != tt.price {
				t.Errorf("Expected %s -> got %s", tt.title, ad.Title)
			}

			_, ok := inMemoryAdRepository.Find(ad.Id)
			if !ok {
				t.Errorf("Ad {%s} not found on repository", ad.Id)
			}
		})
	}
}

func TestFindAd(t *testing.T) {
	tests := map[string]struct {
		id string
		ok bool
	}{
		"not found ad": {
			id: "not valid",
			ok: false,
		},
	}
	for name, tt := range tests {
		inMemoryAdRepository := infra.NewInMemoryAdRepository()
		service := AdService{Repository: inMemoryAdRepository}
		t.Run(name, func(t *testing.T) {
			_, ok := service.Find(tt.id)
			if ok != tt.ok {
				t.Errorf("Expected %v, got %v", tt.ok, ok)
			}
		})
	}
}

func TestFindValidAd(t *testing.T) {
	inMemoryAdRepository := infra.NewInMemoryAdRepository()
	service := AdService{Repository: inMemoryAdRepository}
	ad, _ := domain.NewAd("t", "d", 1)
	inMemoryAdRepository.Save(ad)

	actualAd, _ := service.Find(ad.Id.String())

	if !reflect.DeepEqual(ad, actualAd) {
		t.Errorf("Expected %v, got %v", ad, actualAd)
	}
}

func TestFindRandomAds(t *testing.T) {
	inMemoryAdRepository := infra.NewInMemoryAdRepository()
	service := AdService{Repository: inMemoryAdRepository}
	ad1, _ := domain.NewAd("t1", "d", 1)
	inMemoryAdRepository.Save(ad1)
	ad2, _ := domain.NewAd("t2", "d", 1)
	inMemoryAdRepository.Save(ad2)

	smallAds, _ := service.FindRandom()
	if len(smallAds) != 2 {
		t.Errorf("Expected 2, got %v", len(smallAds))
	}

	ad3, _ := domain.NewAd("t3", "d", 1)
	inMemoryAdRepository.Save(ad3)
	ad4, _ := domain.NewAd("t4", "d", 1)
	inMemoryAdRepository.Save(ad4)
	ad5, _ := domain.NewAd("t5", "d", 1)
	inMemoryAdRepository.Save(ad5)

	allAds, _ := service.FindRandom()

	if len(allAds) != 5 {
		t.Errorf("Expected 5, got %v", len(allAds))
	}

	ad6, _ := domain.NewAd("t6", "d", 1)
	inMemoryAdRepository.Save(ad6)

	bigAds, _ := service.FindRandom()

	if len(bigAds) != 5 {
		t.Errorf("Expected 5, got %v", len(bigAds))
	}

}
