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
	params, _ := io.ReadAll(req.Body)
	var loginDto domain.LoginDTO
	err := json.Unmarshal(params, &loginDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	jwt, err := h.userUseCase.Login(loginDto)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}
	jwtJson, err := json.Marshal(jwt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jwtJson)
	return err
}
