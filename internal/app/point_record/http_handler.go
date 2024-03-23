package point_record

import (
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, err = w.Write([]byte(`{"error": "unauthorized"}`))
		return err
	}
	result, err := h.registerPointUseCase.RecordPointEvent(domain.RegisterPointDTO{UserID: userId})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error": "internal server error"}`))
		return err
	}
	log.Println("success")
	err = bunrouter.JSON(w, result)
	return err
}

func (h *HttpHandler) GetRegistersDay(w http.ResponseWriter, req bunrouter.Request) error {
	log.Println("request get daily report")
	userId, err := h.authentication.ExtractUserIDFromToken(req.Header.Get("Authorization"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, err = w.Write([]byte(`{"error": "unauthorized"}`))
		return err
	}
	result, err := h.registerPointUseCase.GetRegistersDay(userId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error": "internal server error"}`))
		return err
	}
	log.Println("success")
	err = bunrouter.JSON(w, result)
	return err
}

func (h *HttpHandler) GetMonthlyReport(w http.ResponseWriter, req bunrouter.Request) error {
	log.Println("request get monthly report")
	userId, err := h.authentication.ExtractUserIDFromToken(req.Header.Get("Authorization"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, err = w.Write([]byte(`{"error": "unauthorized"}`))
		return err
	}
	result, err := h.registerPointUseCase.GetMonthlyReport(userId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error": "internal server error"}`))
		return err
	}
	log.Println("success")
	message := domain.ReportSuccess{
		Message: "monthly report sent to " + result.Email,
	}
	err = bunrouter.JSON(w, message)
	return err
}
