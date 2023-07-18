package domain

import (
	"testing"
)

func TestInvalidAds(t *testing.T) {
	tests := map[string]struct {
		title       string
		description string
		price       float32
		expectedErr error
	}{
		"empty title": {
			expectedErr: invalidTitleError{},
		},
		"blank title": {
			title:       "",
			expectedErr: invalidTitleError{},
		},
		"empty description": {
			title:       "title",
			expectedErr: InvalidDescriptionError{},
		},
		"invalid price": {
			title:       "t",
			description: "d",
			expectedErr: invalidPriceError{},
		},
		"negative price": {
			title:       "t",
			description: "d",
			price:       -1,
			expectedErr: invalidPriceError{},
		},
		"zero price": {
			title:       "t",
			description: "d",
			price:       0,
			expectedErr: invalidPriceError{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewAd(tt.title, tt.description, tt.price)
			if err != tt.expectedErr {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
