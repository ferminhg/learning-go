package application

import (
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
	"time"
)

//type PostAdService interface {
//
//}

func postAdService(title string, description string, price float32) domain.Ad {
	adId, _ := uuid.NewUUID()
	ad := domain.Ad{
		Id:          adId,
		Title:       title,
		Description: description,
		Price:       price,
		CreatedDate: time.Now(),
	}
	return ad
}
