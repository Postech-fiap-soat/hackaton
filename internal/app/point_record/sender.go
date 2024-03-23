package point_record

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"hackaton/internal/app/domain"
	"hackaton/internal/config"
	"html/template"
	"log"
	"strconv"
)

type PointRecordSender struct {
	cfg *config.Config
}

func NewPointRecordSender(cfg *config.Config) *PointRecordSender {
	return &PointRecordSender{
		cfg: cfg,
	}
}
func (p *PointRecordSender) SendMonthlyReport(monthlyReport *domain.MonthlyReport) error {
	t := template.New("template.html")
	var err error
	t, err = t.ParseFiles("template.html")
	if err != nil {
		log.Println(err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, monthlyReport); err != nil {
		log.Println(err)
	}
	result := tpl.String()
	port, err := strconv.Atoi(p.cfg.SMTPPort)
	if err != nil {
		return err
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", "from@gmail.com")
	msg.SetHeader("To", "to@gmail.com")
	msg.SetHeader("Subject", "Relatório Mensal")
	msg.SetBody("text/html", result)
	dialer := gomail.NewDialer(p.cfg.SMTPHost, port, p.cfg.SMTPUser, p.cfg.SMTPPassword)
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
