package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"tyranno/backend/adapter/middleware"
	"tyranno/backend/utils/config"

	"github.com/dgrijalva/jwt-go"
)

func createToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "username",
	})

	tokenString, _ := token.SignedString([]byte(config.SigningKey))
	return tokenString
}

func TestJWTMiddleware_ValidToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		t.Fatal(err)
	}

	token := createToken()
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	handler := middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected route"))
	}))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Protected route"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestJWTMiddleware_InvalidToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer invalidtoken")

	rr := httptest.NewRecorder()
	handler := middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected route"))
	}))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	if !strings.Contains(rr.Body.String(), "Invalid token") {
		t.Errorf("Expected 'Invalid token' in response but got: %v", rr.Body.String())
	}
}

func TestJWTMiddleware_NoToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/protected", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected route"))
	}))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	if !strings.Contains(rr.Body.String(), "Authorization header not found") {
		t.Errorf("Expected 'Authorization header not found' in response but got: %v", rr.Body.String())
	}
}
