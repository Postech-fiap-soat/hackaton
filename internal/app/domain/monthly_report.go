package domain

type MonthlyReport struct {
	Month           string        `json:"month"`
	TotalHourWorked string        `json:"total_hour_worked"`
	DailyReports    []DailyReport `json:"daily_reports"`
}

type ReportSuccess struct {
	Message string `json:"message"`
}

func NewMonthlyReport(pointsRecordedToday []*PointRecord, month string) *MonthlyReport {
	daily := map[int][]*PointRecord{}
	for _, v := range pointsRecordedToday {
		day := v.CreatedAt.Day()
		daily[day] = append(daily[day], v)
	}
	var dailyReports []DailyReport
	for _, v := range daily {
		var dailyReport *DailyReport
		dailyReport = NewDailyReport(v, *v[0].CreatedAt)
		dailyReports = append(dailyReports, *dailyReport)
	}
	return &MonthlyReport{Month: month, DailyReports: dailyReports}
}
