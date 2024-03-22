package domain

type MonthlyReport struct {
	Month           string
	TotalHourWorked string
	DailyReports    []DailyReport
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
