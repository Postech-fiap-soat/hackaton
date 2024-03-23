package domain

import (
	"strings"
	"time"
)

type DailyReport struct {
	Date                 string            `json:"date"`
	TotalHourWorked      string            `json:"total_hour_worked"`
	DailyReportRegisters []DailyReportItem `json:"recorded_points"`
}

type DailyReportItem struct {
	TimeEvent string `json:"time_event"`
	Type      string `json:"type"`
}

var typeEvents = map[int]string{
	1: "Entrada",
	2: "Pausa para intervalo",
	3: "Volta do intervalo",
	4: "Saída",
}

func NewDailyReport(pointsRecordedToday []*PointRecord, dateTimeReport time.Time) *DailyReport {
	var report DailyReport
	report.Date = dateTimeReport.Format("02/01/2006")
	var lastTimeEvent *time.Time
	var totalHourWorked time.Duration
	for _, v := range pointsRecordedToday {
		var item DailyReportItem
		item.TimeEvent = v.CreatedAt.Format("15:04:05")
		item.Type = typeEvents[v.Type]
		report.DailyReportRegisters = append(report.DailyReportRegisters, item)
		if v.Type == 1 || v.Type == 3 {
			lastTimeEvent = v.CreatedAt
			continue
		}
		totalHourWorked += v.CreatedAt.Sub(*lastTimeEvent)
		lastTimeEvent = v.CreatedAt
	}
	if len(pointsRecordedToday) == 1 || len(pointsRecordedToday) == 3 {
		totalHourWorked += time.Now().Sub(*pointsRecordedToday[len(pointsRecordedToday)-1].CreatedAt)
	}
	totalArr := strings.Split(totalHourWorked.String(), ".")
	report.TotalHourWorked = totalArr[0]
	return &report
}
