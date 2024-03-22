package util

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Hash(data []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, 8)
}

func HashMatched(hash []byte, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, plain)
	return err == nil
}

func GenerateJWTToken(isAdmin bool, id uint) (string, error) {
	claims := jwt.MapClaims{
		"admin": isAdmin,
		"sub":   id,
		"exp":   time.Now().Add(5 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GetJWTClaims(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}

func GetSubFromClaims(claims map[string]any) (uint, error) {
	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid sub")
	}

	return uint(sub), nil
}
