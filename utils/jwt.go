package utils

import (
	"go-cloud/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("cloud-main")

type Claims struct {
	UserID         string `json:"user_id"`
	Role           string `json:"role"`
	Email          string `json:"email"`
	UserType       string `json:"user_type"`
	OrganizationID string `json:"organization_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Role: user.Role,
		Email: user.Email,
		UserType: user.UserType,
		OrganizationID: user.Organization.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},

	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
