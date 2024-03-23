package domain

import (
	"errors"
	"time"
)

type PointRecordRepository interface {
	RegisterPoint(pointRecord *PointRecord) (*PointRecord, error)
	GetPointsRecordedToday(userID int, dateNow, dateToday time.Time) ([]*PointRecord, error)
	GetPointsRecordedInMonth(userID int, initDate, finalDate time.Time) ([]*PointRecord, error)
}

type PointRecordUseCase interface {
	RecordPointEvent(registerPointDTO RegisterPointDTO) (*PointRecord, error)
	GetRegistersDay(userID int) (*DailyReport, error)
	GetMonthlyReport(userID int) (*User, error)
}

type PointRecordSender interface {
	SendMonthlyReport(*MonthlyReport, *User) error
}

type RegisterPointDTO struct {
	UserID int
}

type PointRecord struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	Type      int        `json:"type"`
	UsersId   int        `json:"users_id"`
}

const EnterPoint = 1
const GoIntervalPoint = 2
const BackIntervalPoint = 3
const ExitPoint = 4

func NewPointRecord(recordsTodayNumber int, timeNow time.Time, usersID int) (*PointRecord, error) {
	var typePoint int
	if recordsTodayNumber == 0 {
		typePoint = EnterPoint
	} else if recordsTodayNumber == 1 {
		typePoint = GoIntervalPoint
	} else if recordsTodayNumber == 2 {
		typePoint = BackIntervalPoint
	} else if recordsTodayNumber == 3 {
		typePoint = ExitPoint
	} else {
		return nil, errors.New("invalid point record")
	}
	return &PointRecord{
		CreatedAt: &timeNow,
		Type:      typePoint,
		UsersId:   usersID,
	}, nil
}
