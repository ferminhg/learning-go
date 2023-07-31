package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"Description"`
	Price       float32   `json:"price"`
	CreatedDate time.Time `json:"postedAt"`
}

const descriptionMaxSize = 50

type invalidTitleError struct {
}

func (i invalidTitleError) Error() string {
	return "invalid title"
}

type InvalidDescriptionError struct {
}

func (i InvalidDescriptionError) Error() string {
	return "invalid Description"
}

type invalidPriceError struct {
}

func (i invalidPriceError) Error() string {
	return "invalid price"
}

func NewAd(title string, description string, price float32) (Ad, error) {
	if len(title) == 0 {
		return Ad{}, invalidTitleError{}
	}

	if len(description) == 0 || len(description) > descriptionMaxSize {
		return Ad{}, InvalidDescriptionError{}
	}

	if price <= 0 {
		return Ad{}, invalidPriceError{}
	}

	adId, err := uuid.NewRandom()
	if err != nil {
		return Ad{}, err
	}

	ad := Ad{
		Id:          adId,
		Title:       title,
		Description: description,
		Price:       price,
		CreatedDate: time.Now(),
	}
	return ad, nil
}
