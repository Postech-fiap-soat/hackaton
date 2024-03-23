package user

import (
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"hackaton/internal/app/domain"
	"io"
	"log"
	"net/http"
)

type HttpHandler struct {
	userUseCase domain.UserUseCase
}

func NewHttpHandler(userUseCase domain.UserUseCase) *HttpHandler {
	return &HttpHandler{userUseCase: userUseCase}
}

func (h *HttpHandler) Login(w http.ResponseWriter, req bunrouter.Request) error {
	log.Println("request login")
	params, _ := io.ReadAll(req.Body)
	var loginDto domain.LoginDTO
	err := json.Unmarshal(params, &loginDto)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(`{"error": "bad request"}`))
		return err
	}
	jwt, err := h.userUseCase.Login(loginDto)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error": "internal server error"}`))
		return err
	}
	err = bunrouter.JSON(w, jwt)
	return err
}
