package generator

import (
	"fmt"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"time"
)

type FakerDescriptionGenerator struct {
}

func NewFakerDescriptionGenerator() *FakerDescriptionGenerator {
	faker.ResetUnique()
	return &FakerDescriptionGenerator{}
}

func (f FakerDescriptionGenerator) Run(title string) ([]domain.RandomDescription, error) {
	var descriptions []domain.RandomDescription

	ch := make(chan domain.RandomDescription)

	for i := 0; i < 3; i++ {
		go f.generate(title, ch)
		descriptions = append(descriptions, <-ch)
	}

	return descriptions, nil
}

func (f FakerDescriptionGenerator) generate(_ string, ch chan domain.RandomDescription) {
	random := rand.Intn(1000)
	fmt.Println(time.Now(), random)
	ch <- domain.NewRandomDescription(faker.Sentence(), float32(random)/1000.0)
}
