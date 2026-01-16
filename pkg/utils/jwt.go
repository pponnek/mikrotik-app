package utils

import (
	"errors"
	"mikrotikapp/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var cfg = config.Load()

func GenerateToken(userID, tenantID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"tenant_id": tenantID,
		"role": role,
		"exp": time.Now().Add(time.Hour *24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	signed, err := token.SignedString([]byte(config.Load().JWTSecret))

	if err != nil {
		return "", err
	}

	return signed, nil
}

func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot parse claims")
	}

	return claims, nil
}

