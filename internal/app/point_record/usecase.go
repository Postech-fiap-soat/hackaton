package point_record

import (
	"hackaton/internal/app/domain"
	"time"
)

type PointRecordUseCase struct {
	recordPointRepository domain.PointRecordRepository
	pointRecordSender     domain.PointRecordSender
	userRepository        domain.UserRepository
}

func NewRegisterPointUseCase(
	recordPointRepository domain.PointRecordRepository,
	pointRecordSender domain.PointRecordSender,
	userRepository domain.UserRepository) *PointRecordUseCase {
	return &PointRecordUseCase{
		recordPointRepository: recordPointRepository,
		pointRecordSender:     pointRecordSender,
		userRepository:        userRepository,
	}
}

func (r *PointRecordUseCase) RecordPointEvent(registerPointDTO domain.RegisterPointDTO) (*domain.PointRecord, error) {
	dateTimeNow := time.Now()
	dateToday, err := time.Parse("2006-01-02", dateTimeNow.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	pointsRecordedToday, err := r.recordPointRepository.GetPointsRecordedToday(registerPointDTO.UserID, dateTimeNow, dateToday)
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
	pointsRecordedToday, err := r.recordPointRepository.GetPointsRecordedToday(userID, dateTimeNow, dateToday)
	if err != nil {
		return nil, err
	}
	dailyReport := domain.NewDailyReport(pointsRecordedToday, dateTimeNow)
	return dailyReport, nil
}

func (r *PointRecordUseCase) GetMonthlyReport(userID int) (*domain.MonthlyReport, error) {
	lastMonth := time.Now().AddDate(0, -1, 0)
	initLastMonth := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, lastMonth.Location())
	finalLastMonth := initLastMonth.AddDate(0, 1, -1)
	pointsRecordedToday, err := r.recordPointRepository.GetPointsRecordedInMonth(userID, initLastMonth, finalLastMonth)
	if err != nil {
		return nil, err
	}
	monthlyReport := domain.NewMonthlyReport(pointsRecordedToday, lastMonth.Month().String())
	user, err := r.userRepository.GetUserById(userID)
	if err != nil {
		return nil, err
	}
	err = r.pointRecordSender.SendMonthlyReport(monthlyReport, user)
	if err != nil {
		return nil, err
	}
	return monthlyReport, nil
}
