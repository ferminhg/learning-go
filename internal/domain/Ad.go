package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	Id          uuid.UUID
	Title       string
	Description string
	Price       float32
	CreatedDate time.Time
}

type invalidTitleError struct {
}

func (i invalidTitleError) Error() string {
	return "invalid title"
}

type invalidDescriptionError struct {
}

func (i invalidDescriptionError) Error() string {
	return "invalid description"
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

	if len(description) == 0 {
		return Ad{}, invalidDescriptionError{}
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
