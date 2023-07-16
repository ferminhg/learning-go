package application

import "testing"

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
		t.Run(name, func(t *testing.T) {
			ad := postAdService(tt.title, tt.description, tt.price)
			if ad.Title != tt.title {
				t.Errorf("Expected %s -> got %s", tt.title, ad.Title)
			}
		})
	}
}
