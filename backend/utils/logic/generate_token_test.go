package logic_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"tyranno/backend/utils/config"
	"tyranno/backend/utils/logic"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateToken(t *testing.T) {
	config.LoadEnv()
	username := "testUser"

	token, err := logic.GenerateToken(username)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Error("Generated token is empty")
	}

	// Parsing and validating the token
	claims := &logic.Claims{}
	tokenObj, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SigningKey), nil
	})

	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	if !tokenObj.Valid {
		t.Error("Token is not valid")
	}

	// Validating the token content
	if claims.Username != username {
		t.Errorf("Expected username %v but got %v", username, claims.Username)
	}

	expirationTimeInMinutes, err := strconv.Atoi(config.TokenLifeTime)
	if err != nil {
		t.Errorf("TokenLifeTime is invalid %v", config.TokenLifeTime)
	}

	// The time.Duration(expirationTimeInMinutes) won't be exact due to the time it takes to execute, but it should be very close.
	if time.Until(time.Unix(claims.ExpiresAt, 0)) > time.Duration(expirationTimeInMinutes)+time.Second || time.Until(time.Unix(claims.ExpiresAt, 0)) < time.Duration(expirationTimeInMinutes)-time.Second {
		t.Errorf("Expected expiration to be about %v from now, but got %v", time.Duration(expirationTimeInMinutes), time.Until(time.Unix(claims.ExpiresAt, 0)))
	}
}
