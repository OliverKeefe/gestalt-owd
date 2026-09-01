package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type ctxStringKey string

type ctxKey int

const userIDKey ctxStringKey = "userID"
const claimsKey ctxKey = iota
const typedClaimsKey ctxKey = iota

type RealmAccess struct {
	Roles []string `json:"roles"`
}

type ClientAccess struct {
	Roles []string `json:"roles"`
}

type KeycloakClaims struct {
	jwt.RegisteredClaims
	Email             string                  `json:"email"`
	PreferredUsername string                  `json:"preferred_username"`
	Groups            []string                `json:"groups"`
	RealmAccess       RealmAccess             `json:"realm_access"`
	ResourceAccess    map[string]ClientAccess `json:"resource_access"`
}

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

func (k *Authenticator) ValidateJWT(ctx context.Context, jwtB64 string) (context.Context, bool, error) {
	claims := &KeycloakClaims{}
	kf := func(t *jwt.Token) (any, error) {
		return k.KeyFunc.Keyfunc(t)
	}
	token, err := jwt.ParseWithClaims(
		jwtB64,
		claims,
		kf,
		jwt.WithValidMethods([]string{"RS256"}),
	)
	if err != nil {
		log.Printf("failed to parse the JWT. %v", err)
		return ctx, false, err
	}

	if !token.Valid {
		log.Printf("invalid token.")
		return ctx, false, nil
	}

	if claims.Issuer != k.Issuer {
		return ctx, false, errors.New("invalid issuer")
	}

	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return ctx, false, fmt.Errorf("failed to marshal claims: %w", err)
	}
	claimsMap := make(map[string]interface{})
	if err := json.Unmarshal(claimsBytes, &claimsMap); err != nil {
		return ctx, false, fmt.Errorf("failed to unmarshal claims: %w", err)
	}

	newCtx := context.WithValue(ctx, userIDKey, claims.Subject)
	newCtx = context.WithValue(newCtx, claimsKey, claimsMap)
	newCtx = context.WithValue(newCtx, typedClaimsKey, claims)
	return newCtx, true, nil
}

func (k *Authenticator) InjectUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func (k *Authenticator) InjectClaims(ctx context.Context, claims map[string]interface{}) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

// InjectKeycloakClaims stores a fully-populated KeycloakClaims value in the
// context. Primarily useful for tests and for contexts built outside JWT
// validation.
func (k *Authenticator) InjectKeycloakClaims(ctx context.Context, claims *KeycloakClaims) context.Context {
	return context.WithValue(ctx, typedClaimsKey, claims)
}
