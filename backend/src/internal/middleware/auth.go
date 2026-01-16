package middleware

import (
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtRequest struct {
	alg    string `json:"alg"`
	typ    string `json:"typ"`
	sub    string `json:"sub"`
	claims string `json:"name"`
	admin  string `json:"admin"`
	iat    string `json:"iat"`
}

func ValidateJWT(req jwtRequest) (bool, error) {
	token, err := jwt.Parse(req)
	if err != nil {
		return false, fmt.Errorf("unable to parse jwt request: %v", err)
	}
