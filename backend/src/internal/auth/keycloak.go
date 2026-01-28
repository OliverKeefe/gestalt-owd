package auth

import (
	"log"

var (
	_ = OAuth2(*Keycloak)(nil)
)

type Authenticator struct {
	Issuer  string
	KeyFunc keyfunc.Keyfunc
}

func New(issuer, jwksUrl string) (*Authenticator, error) {
	kf, err := keyfunc.NewDefault([]string{jwksUrl})
	if err != nil {
		log.Printf("Failed to create JWK Set from resource at the given URL, %v", err)
		return nil, err
	}

	return &Authenticator{
		Issuer:  issuer,
		KeyFunc: kf,
	}, nil
}

func (k *Authenticator) ValidateJWT(jwtB64 string) (bool, error) {
	log.Printf(jwtB64)
	kf := func(t *jwt.Token) (any, error) {
		return k.KeyFunc.Keyfunc(t)
	}
	token, err := jwt.Parse(jwtB64, kf)
	if err != nil {
		log.Printf("failed to parse the JWT. %v", err)
		return false, err
	}

	//TODO: map claims to struct and verify using jwt pkg.

	if !token.Valid {
		log.Printf("invalid token.")
		return false, nil
	}

	return true, nil
}

func (k *Authenticator) ReissueJWT() (jwt.Token, error) {
	var tkn jwt.Token
	return tkn, nil
}
