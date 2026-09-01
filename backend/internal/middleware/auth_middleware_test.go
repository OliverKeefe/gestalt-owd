package middleware

import (
	"backend/internal/auth"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newAuthenticator() *auth.Authenticator {
	return &auth.Authenticator{}
}

func injectGroups(t *testing.T, groups []string) context.Context {
	t.Helper()
	a := newAuthenticator()
	return a.InjectKeycloakClaims(context.Background(), &auth.KeycloakClaims{
		Groups: groups,
	})
}

func TestRequireGroup_Success(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireGroup("org-gestalt", next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(injectGroups(t, []string{"org-gestalt"}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestRequireGroup_Forbidden(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireGroup("org-other", next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(injectGroups(t, []string{"org-gestalt"}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", rec.Code)
	}
}

func TestRequireGroup_NoGroups(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireGroup("org-gestalt", next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(context.Background())

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", rec.Code)
	}
}

func TestRequireAnyGroup_Success(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireAnyGroup([]string{"missing", "org-gestalt"}, next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(injectGroups(t, []string{"org-gestalt"}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestRequireRealmRole_Success(t *testing.T) {
	a := newAuthenticator()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireRealmRole(a, "admin", next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(a.InjectKeycloakClaims(context.Background(), &auth.KeycloakClaims{
		RealmAccess: auth.RealmAccess{Roles: []string{"admin"}},
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestRequireClientRole_Success(t *testing.T) {
	a := newAuthenticator()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireClientRole(a, "open-web-drive", "files:upload", next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(a.InjectKeycloakClaims(context.Background(), &auth.KeycloakClaims{
		ResourceAccess: map[string]auth.ClientAccess{
			"open-web-drive": {Roles: []string{"files:upload"}},
		},
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestRequireClientRole_Forbidden(t *testing.T) {
	a := newAuthenticator()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireClientRole(a, "open-web-drive", "files:delete", next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(a.InjectKeycloakClaims(context.Background(), &auth.KeycloakClaims{
		ResourceAccess: map[string]auth.ClientAccess{
			"open-web-drive": {Roles: []string{"files:upload"}},
		},
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", rec.Code)
	}
}

func TestAnyRealmRole(t *testing.T) {
	a := newAuthenticator()
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := RequireAnyRealmRole(a, []string{"admin", "user"}, next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(a.InjectKeycloakClaims(context.Background(), &auth.KeycloakClaims{
		RealmAccess: auth.RealmAccess{Roles: []string{"user"}},
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
