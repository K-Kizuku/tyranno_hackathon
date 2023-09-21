package logic

import (
	"log"
	"strconv"
	"time"
	"tyranno/backend/utils/config"

	"github.com/dgrijalva/jwt-go"
)

// Claims is the payload data we'll be storing in our JWT token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTimeInMinutes, err := strconv.Atoi(config.TokenLifeTime)
	if err != nil {
		log.Println("TokenLifeTime is invalid")
	}
	expirationTime := time.Now().Add(time.Duration(expirationTimeInMinutes) * time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.SigningKey))
}
