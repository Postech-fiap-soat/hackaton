package domain

import (
	"time"
)

type DailyReport struct {
	Date                 string
	TotalHourWorked      string
	DailyReportRegisters []DailyReportItem
}

type DailyReportItem struct {
	TimeEvent string
	Type      string
}

var typeEvents = map[int]string{
	1: "Entrada",
	2: "Pausa para intervalo",
	3: "Volta do intervalo",
	4: "Saída",
}

func NewDailyReport(pointsRecordedToday []*PointRecord) *DailyReport {
	var report DailyReport
	report.Date = time.Now().Format("02/01/2006")
	var lastTimeEvent *time.Time
	for _, v := range pointsRecordedToday {
		var item DailyReportItem
		item.TimeEvent = v.CreatedAt.Format("15:04:05")
		item.Type = typeEvents[v.Type]
		if lastTimeEvent != nil {
			lastTimeEvent = v.CreatedAt
			continue
		}
		report.DailyReportRegisters = append(report.DailyReportRegisters, item)
		//difference := v.CreatedAt.Sub(*lastTimeEvent)
		//fmt.Println(difference)
	}
	return &report
}
