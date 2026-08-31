package middleware

import (
	"net/http"
	"strings"

	"backend/internal/auth"
)

// AuthMiddleware returns middleware that validates a Keycloak JWT on every request.
func AuthMiddleware(a *auth.Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return Protect(a, next)
	}
}

// Protect wraps an http.Handler to require a valid Keycloak Bearer token before serving the request.
func Protect(a *auth.Authenticator, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(h, "Bearer ")
		newCtx, ok, err := a.ValidateJWT(r.Context(), token)
		if err != nil || !ok {
			http.Error(w, "unauthorized, invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

// RequireClientRole wraps an http.Handler to require the authenticated user to hold the specified Keycloak client role.
func RequireClientRole(a *auth.Authenticator, clientID, role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.HasClientRole(r.Context(), clientID, role) {
			http.Error(w, "forbidden: missing required role", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RequireRealmRole wraps an http.Handler to require the authenticated user to hold the specified Keycloak realm role.
func RequireRealmRole(a *auth.Authenticator, role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.HasRealmRole(r.Context(), role) {
			http.Error(w, "forbidden: missing required role", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RequireAnyRealmRole wraps an http.Handler to require the authenticated user to hold at least one of the specified Keycloak realm roles.
func RequireAnyRealmRole(a *auth.Authenticator, roles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, role := range roles {
			if auth.HasRealmRole(r.Context(), role) {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "forbidden: missing required role", http.StatusForbidden)
	})
}

// RequireGroup wraps an http.Handler to require the authenticated user to belong to the specified Keycloak group.
func RequireGroup(group string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.HasGroup(r.Context(), group) {
			http.Error(w, "forbidden: user is not a member of the required group", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RequireAnyGroup wraps an http.Handler to require the authenticated user to belong to at least one of the specified Keycloak groups.
func RequireAnyGroup(groups []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.HasAnyGroup(r.Context(), groups) {
			http.Error(w, "forbidden: user is not a member of any required group", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
