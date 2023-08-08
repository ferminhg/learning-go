package postgres

import (
	"database/sql"
	"fmt"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgresAdRepository_Save(t *testing.T) {
	connStr := "host=localhost port=5432 user=wopwop password=wopwop dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	defer db.Close()

	repo := NewPostgresAdRepository(db)

	t.Run("Given a valid params should insert into DB", func(t *testing.T) {
		ad := domain.RandomAdFactory()

		err := repo.Save(ad)
		require.NoError(t, err)

		findAd, ok := repo.Find(ad.Id)

		assert.True(t, ok)

		assert.Equal(t, ad.Id, findAd.Id)
		assert.Equal(t, ad.Title, findAd.Title)
		assert.Equal(t, ad.Description, findAd.Description)
		assert.Equal(t, ad.Price, findAd.Price)
		assert.Equal(t,
			ad.CreatedDate.Format(`%d/%m/%y %H:%m:%s`),
			findAd.CreatedDate.Format(`%d/%m/%y %H:%m:%s`))
	})
}

func TestPostgresAdRepository_Find(t *testing.T) {
	connStr := "host=localhost port=5432 user=wopwop password=wopwop dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	defer db.Close()

	repo := NewPostgresAdRepository(db)
	t.Run("Given a not valid AdId should return false", func(t *testing.T) {
		notExistUUID, _ := uuid.NewRandom()
		_, ok := repo.Find(notExistUUID)
		assert.False(t, ok)
	})

	t.Run("Given a valid AdId shoult return it", func(t *testing.T) {
		AdId, _ := uuid.Parse("24c95c22-2d2b-11ee-be56-0242ac120002")

		findAd, ok := repo.Find(AdId)
		assert.True(t, ok)
		assert.Equal(t, findAd.Id, AdId)
		fmt.Println(findAd)
	})
}

func TestPostgresAdRepository_FindRandom(t *testing.T) {
	connStr := "host=localhost port=5432 user=wopwop password=wopwop dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	defer db.Close()

	repo := NewPostgresAdRepository(db)
	repo.Save(domain.RandomAdFactory())
	repo.Save(domain.RandomAdFactory())
	repo.Save(domain.RandomAdFactory())
	repo.Save(domain.RandomAdFactory())
	repo.Save(domain.RandomAdFactory())

	t.Run("Given a repo should return find Ads", func(t *testing.T) {
		ads, err := repo.Search(5)
		require.NoError(t, err)

		assert.Equal(t, 5, len(ads))

	})
}

func TestPostgresAdRepository_Delete(t *testing.T) {
	connStr := "host=localhost port=5432 user=wopwop password=wopwop dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	defer db.Close()

	notExistId, _ := uuid.NewRandom()

	repo := NewPostgresAdRepository(db)
	assert.False(t, repo.Delete(notExistId))

	ad := domain.RandomAdFactory()
	repo.Save(ad)
	assert.True(t, repo.Delete(ad.Id))
}
