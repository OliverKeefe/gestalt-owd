package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type JWTRequest struct {
	Alg    string `json:"alg"`
	Typ    string `json:"typ"`
	Sub    string `json:"sub"`
	Claims string `json:"name"`
	Admin  string `json:"admin"`
	Iat    string `json:"iat"`
}

type Claims struct {
	Claims interface{}
}

type OAuth2 interface {
	ValidateJWT(ctx context.Context, rawJWT string) (*Claims, error)
	ReissueJWT(ctx context.Context, refreshToken string) (string, error)
}

type UCAN interface {
	ValidateJWT(token jwt.Token) (bool, error)
}
