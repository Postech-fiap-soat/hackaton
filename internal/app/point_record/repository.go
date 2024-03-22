package point_record

import (
	"database/sql"
	"fmt"
	"hackaton/internal/app/domain"
	"time"
)

type PointRecordRepository struct {
	db *sql.DB
}

func NewPointRecordRepository(db *sql.DB) *PointRecordRepository {
	return &PointRecordRepository{db: db}
}

func (p *PointRecordRepository) RegisterPoint(pointRecord *domain.PointRecord) (*domain.PointRecord, error) {
	stmt, err := p.db.Prepare("insert into point_records (id, created_at, type, users_id) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(pointRecord.ID, pointRecord.CreatedAt, pointRecord.Type, pointRecord.UsersId)
	return pointRecord, nil
}

func (p *PointRecordRepository) GetPointsRecordedToday(userID int, dateToday time.Time) ([]*domain.PointRecord, error) {
	stmt, err := p.db.Prepare("select id, created_at, type, users_id from point_records where users_id = 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var pointRecords []*domain.PointRecord
	for rows.Next() {
		var pointRecord domain.PointRecord
		err = rows.Scan(&pointRecord.ID, &pointRecord.CreatedAt, &pointRecord.Type, &pointRecord.UsersId)
		if err != nil {
			fmt.Println("passou aqui erro", err.Error())
			return nil, err
		}
		pointRecords = append(pointRecords, &pointRecord)
	}
	return pointRecords, nil
}
