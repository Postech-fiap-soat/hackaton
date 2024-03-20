package point_record

import (
	"database/sql"
	"hackaton/internal/app/domain"
	"log"
)

type PointRecordRepository struct {
	db *sql.DB
}

func NewPointRecordRepository(db *sql.DB) *PointRecordRepository {
	return &PointRecordRepository{db: db}
}

func (p *PointRecordRepository) RegisterPoint(pointRecord *domain.PointRecord) (*domain.PointRecord, error) {
	rows, err := p.db.Query("select id, name from users where id = 1")
	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	point := &domain.PointRecord{}
	pointRecord.User = user
	return point, nil
}
