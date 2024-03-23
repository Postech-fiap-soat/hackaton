package point_record

import (
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"hackaton/internal/app/domain"
	"log"
	"net/http"
)

type HttpHandler struct {
	registerPointUseCase domain.PointRecordUseCase
}

func NewHttpHandler(registerPointUseCase domain.PointRecordUseCase) *HttpHandler {
	return &HttpHandler{registerPointUseCase: registerPointUseCase}
}

func (h *HttpHandler) RegisterPoint(w http.ResponseWriter, req bunrouter.Request) error {
	log.Println("request in register point")
	result, err := h.registerPointUseCase.RecordPointEvent(domain.RegisterPointDTO{UserID: 1})
	if err != nil {
		log.Println("ERRO: ", err)
	}
	resultJson, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
	log.Println("success")
	return nil
}

func (h *HttpHandler) GetRegistersDay(w http.ResponseWriter, req bunrouter.Request) error {
	log.Println("request get daily report")
	result, err := h.registerPointUseCase.GetRegistersDay(1)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		log.Println("ERRO: ", err)
	}
	resultJson, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
	log.Println("success")
	return nil
}

func (h *HttpHandler) GetMonthlyReport(w http.ResponseWriter, req bunrouter.Request) error {
	log.Println("request get monthly report")
	result, err := h.registerPointUseCase.GetMonthlyReport(1)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		log.Println("ERRO: ", err)
	}
	resultJson, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
	log.Println("success")
	return nil
}
