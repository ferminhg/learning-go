package application

import (
	"fmt"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/ferminhg/learning-go/internal/infra/eventHandler"
	"github.com/ferminhg/learning-go/internal/infra/generator"
	"github.com/ferminhg/learning-go/internal/infra/storage/inmemory"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

func TestPostAd(t *testing.T) {
	inMemoryAdRepository := inmemory.NewInMemoryAdRepository()
	sp := eventHandler.NewMockEventHandler(t)

	service := AdService{
		repository:   inMemoryAdRepository,
		eventHandler: &sp,
	}

	sp.MockSP.ExpectSendMessageAndSucceed()

	title := "t1"
	description := "d1"
	price := float32(15.1)
	ad, _ := service.Post(title, description, price)
	if ad.Title != title {
		t.Errorf("Expected %s -> got %s", title, ad.Title)
	}
	if ad.Description != description {
		t.Errorf("Expected %s -> got %s", title, ad.Title)
	}
	if ad.Price != price {
		t.Errorf("Expected %s -> got %s", title, ad.Title)
	}

	_, ok := inMemoryAdRepository.Find(ad.Id)
	if !ok {
		t.Errorf("Ad {%s} not found on repository", ad.Id)
	}
}

func TestGivenNotValidId(t *testing.T) {
	inMemoryAdRepository := inmemory.NewInMemoryAdRepository()
	service := AdService{repository: inMemoryAdRepository}
	_, ok := service.Find("not valid")
	assert.False(t, ok)
}

func TestFindValidAd(t *testing.T) {
	inMemoryAdRepository := inmemory.NewInMemoryAdRepository()
	service := AdService{repository: inMemoryAdRepository}
	ad, _ := domain.NewAd("t", "d", 1)
	inMemoryAdRepository.Save(ad)

	actualAd, _ := service.Find(ad.Id.String())

	if !reflect.DeepEqual(ad, actualAd) {
		t.Errorf("Expected %v, got %v", ad, actualAd)
	}
}

func TestFindRandomAds(t *testing.T) {
	inMemoryAdRepository := inmemory.NewInMemoryAdRepository()
	service := AdService{repository: inMemoryAdRepository}
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

// Testing AdService using mocking
type AdServiceTestSuite struct {
	suite.Suite
	repository domain.AdServiceRepository
	service    AdService
}

func (suite *AdServiceTestSuite) SetupTest() {
	fmt.Println("⚒️ Setup Test")
	suite.repository = newMockRepository()
	suite.service = AdService{repository: suite.repository}
}

func TestAdServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AdServiceTestSuite))
}

func newMockRepository() *mockRepository {
	return &mockRepository{}
}

func (suite *AdServiceTestSuite) TestGivenAdThenPost() {
	repository := newMockRepository()
	mockEventHandler := eventHandler.NewMockEventHandler(suite.T())
	mockEventHandler.MockSP.ExpectSendMessageAndSucceed()

	service := AdService{repository: repository, eventHandler: &mockEventHandler}

	repository.On("Save", mock.Anything).Return(nil)

	ad, err := service.Post("t", "d", 1)

	assert.IsType(suite.T(), domain.Ad{}, ad)
	assert.Nil(suite.T(), err)
}

func (suite *AdServiceTestSuite) TestGivenAdWhenNotFound() {
	repository := newMockRepository()
	service := AdService{repository: repository}
	_, ok := service.Find("Not Valid UUID")
	assert.False(suite.T(), ok)
}

func (suite *AdServiceTestSuite) TestGivenAdWhenFound() {
	repository := newMockRepository()
	service := AdService{repository: repository}

	randomId, _ := uuid.NewRandom()
	repository.On("Find", randomId).Return(true)

	ad, ok := service.Find(randomId.String())

	assert.True(suite.T(), ok)
	assert.IsType(suite.T(), domain.Ad{}, ad)
	assert.Equal(suite.T(), randomId, ad.Id)
}

func (suite *AdServiceTestSuite) TestGivenAdSWhenSearch() {
	repository := newMockRepository()
	service := AdService{repository: repository}

	repository.On("Search", 5).Return([5]domain.Ad{})

	ads, err := service.FindRandom()

	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), ads, 5)
}

func (suite *AdServiceTestSuite) TestGivenAdWithLongDesWhenPostThenError() {
	repository := newMockRepository()
	service := AdService{repository: repository}
	repository.On("Save", mock.Anything).Return(nil)

	var longDescription = faker.Paragraph()
	for len(longDescription) < 50 {
		faker.ResetUnique()
		longDescription = faker.Paragraph()
	}
	longDescription = longDescription[:51]
	_, err := service.Post("t", longDescription, 1)
	assert.IsType(suite.T(), err, domain.InvalidDescriptionError{})
}

type mockRepository struct{ mock.Mock }

func (m mockRepository) Save(ad domain.Ad) error {
	args := m.Called(ad)
	return args.Error(0)
}

func (m mockRepository) Find(uuid uuid.UUID) (domain.Ad, bool) {
	args := m.Called(uuid)
	return domain.Ad{Id: uuid}, args.Bool(0)
}

func (m mockRepository) Search(maxNumber int) ([]domain.Ad, error) {
	m.Called(5)
	ads := make([]domain.Ad, 5)
	return ads, nil
}

func (m mockRepository) Delete(uuid uuid.UUID) bool {
	return false
}
func TestAdService_DescriptionGenerator(t *testing.T) {
	service := AdService{
		generator: generator.FakerDescriptionGenerator{},
	}

	t.Run("should return at least 1 descriptions", func(t *testing.T) {
		title := "t1"
		got, err := service.DescriptionGenerator(title)
		require.NoError(t, err)

		assert.IsType(t, domain.RandomDescription{}, got)
	})
}
