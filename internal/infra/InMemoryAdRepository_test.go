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
	gotAgain, ok := suite.repository.Find(ad.Id)

	assert.Equal(suite.T(), ad, gotAgain)
	assert.True(suite.T(), ok)
}

func (suite *InMemoryAdRepositoryTestSuite) TestGivenAdWhenFindThenNotFound() {
	ad := domain.RandomAdFactory()
	_, ok := suite.repository.Find(ad.Id)

	assert.False(suite.T(), ok)
}

func (suite *InMemoryAdRepositoryTestSuite) TestGivenEmptyRepoWhenSearch() {
	actual, err := suite.repository.Search(5)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), actual, 0)
}

func (suite *InMemoryAdRepositoryTestSuite) TestGivenFullRepoWhenSearch() {
	for range [10]struct{}{} {
		suite.repository.Save(domain.RandomAdFactory())
	}
	ads, _ := suite.repository.Search(5)

	assert.Len(suite.T(), ads, 5)
}

func TestInMemoryAdRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(InMemoryAdRepositoryTestSuite))
}
