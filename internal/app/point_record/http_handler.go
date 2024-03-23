package point_record

import (
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"hackaton/internal/app/authentication"
	"hackaton/internal/app/domain"
	"log"
	"net/http"
)

type HttpHandler struct {
	registerPointUseCase domain.PointRecordUseCase
	authentication       *authentication.Authentication
}

func NewHttpHandler(registerPointUseCase domain.PointRecordUseCase, authentication *authentication.Authentication) *HttpHandler {
	return &HttpHandler{registerPointUseCase: registerPointUseCase, authentication: authentication}
}

func (h *HttpHandler) RegisterPoint(w http.ResponseWriter, req bunrouter.Request) error {
	log.Println("request in register point")
	userId, err := h.authentication.ExtractUserIDFromToken(req.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	result, err := h.registerPointUseCase.RecordPointEvent(domain.RegisterPointDTO{UserID: userId})
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
	userId, err := h.authentication.ExtractUserIDFromToken(req.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	result, err := h.registerPointUseCase.GetRegistersDay(userId)
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
	userId, err := h.authentication.ExtractUserIDFromToken(req.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	result, err := h.registerPointUseCase.GetMonthlyReport(userId)
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
