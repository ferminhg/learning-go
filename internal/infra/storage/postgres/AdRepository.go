package postgres

import (
	"database/sql"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/google/uuid"
	"log"

	_ "github.com/lib/pq"
)

type PostgresAdRepository struct {
	db *sql.DB
}

func NewPostgresAdRepository(db *sql.DB) *PostgresAdRepository {
	return &PostgresAdRepository{
		db: db,
	}
}

func (p PostgresAdRepository) Save(ad domain.Ad) error {
	sql := "INSERT INTO ads (id, title, description, price, createddate) VALUES ($1, $2, $3, $4, $5)"

	_, err := p.db.Exec(sql, ad.Id, ad.Title, ad.Description, ad.Price, ad.CreatedDate)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p PostgresAdRepository) Find(adId uuid.UUID) (domain.Ad, bool) {
	row, err := p.db.Query("SELECT id, title, description, price, createddate FROM ads WHERE id = $1", adId.String())
	defer row.Close()
	if err != nil {
		log.Fatalf("error trying to persist course on database: %v", err)
		return domain.Ad{}, false
	}

	ad := domain.Ad{}
	var sqlAdId string
	if !row.Next() {
		return domain.Ad{}, false
	}

	if err := row.Scan(&sqlAdId, &ad.Title, &ad.Description, &ad.Price, &ad.CreatedDate); err != nil {
		return domain.Ad{}, false
	}

	ad.Id, _ = uuid.Parse(sqlAdId)

	return ad, true
}

func (p PostgresAdRepository) Search(maxNumber int) ([]domain.Ad, error) {
	row, err := p.db.Query(
		"SELECT id, title, description, price, createddate FROM ads ORDER BY random() LIMIT $1 ",
		maxNumber)
	defer row.Close()

	var ads []domain.Ad
	if err != nil {
		log.Fatalf("error trying to persist course on database: %v", err)
		return ads, err
	}

	for row.Next() {
		var sqlAdId string
		ad := domain.Ad{}

		row.Scan(&sqlAdId, &ad.Title, &ad.Description, &ad.Price, &ad.CreatedDate)
		ad.Id, _ = uuid.Parse(sqlAdId)
		ads = append(ads, ad)
	}

	return ads, nil
}

func (p PostgresAdRepository) Delete(id uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}
