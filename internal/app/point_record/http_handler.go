package point_record

import (
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"hackaton/internal/app/domain"
	"log"
	"net/http"
)

type HttpHandler struct {
	registerPointUseCase domain.RegisterPointUseCase
}

func NewHttpHandler(registerPointUseCase domain.RegisterPointUseCase) *HttpHandler {
	return &HttpHandler{registerPointUseCase: registerPointUseCase}
}

func (h *HttpHandler) RegisterPoint(w http.ResponseWriter, req bunrouter.Request) error {
	result, err := h.registerPointUseCase.Handle(domain.RegisterPointDTO{})
	if err != nil {
		log.Println("ERRO: ", err)
	}
	resultJson, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
	return nil
}
