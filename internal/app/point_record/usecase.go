package point_record

import (
	"hackaton/internal/app/domain"
	"time"
)

type PointRecordUseCase struct {
	recordPointRepository domain.PointRecordRepository
}

func NewRegisterPointUseCase(recordPointRepository domain.PointRecordRepository) *PointRecordUseCase {
	return &PointRecordUseCase{recordPointRepository: recordPointRepository}
}

func (r *PointRecordUseCase) RecordPointEvent(registerPointDTO domain.RegisterPointDTO) (*domain.PointRecord, error) {
	dateTimeNow := time.Now()
	dateToday, err := time.Parse("2006-01-02", dateTimeNow.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	pointsRecordedToday, err := r.recordPointRepository.GetPointsRecordedToday(registerPointDTO.UserID, dateToday)
	if err != nil {
		return nil, err
	}
	pointRecord, err := domain.NewPointRecord(len(pointsRecordedToday), dateTimeNow, registerPointDTO.UserID)
	if err != nil {
		return nil, err
	}
	point, err := r.recordPointRepository.RegisterPoint(pointRecord)
	if err != nil {
		return nil, err
	}
	return point, nil
}

func (r *PointRecordUseCase) GetRegistersDay(userID int) (*domain.DailyReport, error) {
	dateTimeNow := time.Now()
	dateToday, err := time.Parse("2006-01-02", dateTimeNow.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	pointsRecordedToday, err := r.recordPointRepository.GetPointsRecordedToday(userID, dateToday)
	if err != nil {
		return nil, err
	}
	dailyReport := domain.NewDailyReport(pointsRecordedToday)
	return dailyReport, nil
}
