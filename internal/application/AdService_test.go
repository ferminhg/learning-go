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
		service := AdService{repository: inMemoryAdRepository}
		t.Run(name, func(t *testing.T) {
			ad := service.post(tt.title, tt.description, tt.price)
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
		service := AdService{repository: inMemoryAdRepository}
		t.Run(name, func(t *testing.T) {
			_, ok := service.find(tt.id)
			if ok != tt.ok {
				t.Errorf("Expected %v, got %v", tt.ok, ok)
			}
		})
	}
}

func TestFindValidAd(t *testing.T) {
	inMemoryAdRepository := infra.NewInMemoryAdRepository()
	service := AdService{repository: inMemoryAdRepository}
	ad, _ := domain.NewAd("t", "d", 1)
	inMemoryAdRepository.Save(ad)

	actualAd, _ := service.find(ad.Id.String())

	if !reflect.DeepEqual(ad, actualAd) {
		t.Errorf("Exepected %v, got %v", ad, actualAd)
	}
}
