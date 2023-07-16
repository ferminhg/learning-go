package application

import (
	"github.com/ferminhg/learning-go/internal/infra"
	"testing"
)

func Test_postAddService(t *testing.T) {
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
