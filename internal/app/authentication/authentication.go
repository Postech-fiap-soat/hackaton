package authentication

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
)

type Authentication struct {
	secretKey []byte
}

func NewAuthentication(secretKey string) *Authentication {
	return &Authentication{secretKey: []byte(secretKey)}
}

func (a *Authentication) ExtractUserIDFromToken(bearerToken string) (int, error) {
	strToken := a.extractToken(bearerToken)
	token, err := jwt.Parse(strToken, a.getSecretKey)
	if err != nil {
		return 0, err
	}
	permissions, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}
	userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(userID), nil
}

func (a *Authentication) extractToken(bearerToken string) string {
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (a *Authentication) getSecretKey(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return a.secretKey, nil
}
