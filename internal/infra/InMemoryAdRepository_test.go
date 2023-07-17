package infra

import (
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TestSuitefor InMemoryAdRepository
type InMemoryAdRepositoryTestSuite struct {
	suite.Suite
	repository domain.AdServiceRepository
}

func (suite *InMemoryAdRepositoryTestSuite) SetupTest() {
	suite.repository = NewInMemoryAdRepository()
}

func (suite *InMemoryAdRepositoryTestSuite) TestGivenValidAdWhenSave() {
	ad := domain.RandomAdFactory()
	suite.repository.Save(ad)

	got, _ := suite.repository.Find(ad.Id)
	assert.Equal(suite.T(), ad, got)

	suite.repository.Save(ad)
	gotAgain, _ := suite.repository.Find(ad.Id)

	assert.Equal(suite.T(), ad, gotAgain)
}

func TestInMemoryAdRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(InMemoryAdRepositoryTestSuite))
}
