package generator

import (
	"context"
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

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Millisecond*350)
	defer cancel()

	for i := 0; i < 3; i++ {
		go f.generate(ctxTimeout, title, ch)
		select {
		case <-ctxTimeout.Done():
			fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
		case result := <-ch:
			descriptions = append(descriptions, result)
		}
	}

	return descriptions, nil
}

func (f FakerDescriptionGenerator) generate(
	ctx context.Context,
	_ string,
	ch chan domain.RandomDescription,
) {
	random := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(random))
	ch <- domain.NewRandomDescription(faker.Sentence(), float32(random)/1000.0)
}
