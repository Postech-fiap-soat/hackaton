package domain

import "time"

type PointRecordRepository interface {
	RegisterPoint(pointRecord *PointRecord) (*PointRecord, error)
}

type RegisterPointUseCase interface {
	Handle(registerPointDTO RegisterPointDTO) (*PointRecord, error)
}

type RegisterPointDTO struct {
}

type PointRecord struct {
	User      User
	CreatedAt *time.Time
	Type      string
}
