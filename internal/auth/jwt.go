package auth

import (
	"os"
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

// Load envinronment variables from .env file
func init() {
	godotenv.Load()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Unable to retrieve JWT secret")
	}

	jwtSecret = []byte(secret)
}

// Generate and sign token using jwtSecret
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateTOken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return username, nil
}
