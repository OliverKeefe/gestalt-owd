package auth

import "github.com/golang-jwt/jwt/v5"

var (
	_ = OAuth2(*Keycloak)(nil)
)

type Keycloak struct {
	Issuer   string
	ClientID string
	PubKey   interface{}
}

func (k *Keycloak) ValidateJWT(token jwt.Token) (bool, error) {
	if !token.Valid {
		return false, nil
	}
	return true, nil
}

func (k *Keycloak) ReissueJWT() (jwt.Token, error) {
	var tkn jwt.Token
	return tkn, nil
}
