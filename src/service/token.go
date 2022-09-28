package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"auth/src/constanta/response"
	"auth/src/models"

	"github.com/dgrijalva/jwt-go"
)

type Token interface {
	GenerateToken(ctx context.Context, user *models.User) (*response.ResultResponse, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
}

func NewTokenUc() Token {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	return secretKey
}

func (u *jwtService) GenerateToken(ctx context.Context, user *models.User) (*response.ResultResponse, error) {

	now := time.Now()
	end := now.Add(time.Minute * 15)
	claims := &models.AccessCustomClaim{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := newToken.SignedString([]byte(u.secretKey))
	if err != nil {
		fmt.Println("3")
		return nil, err
	}

	res := &response.ResultResponse{
		Name:      user.Name,
		Token:     tokenStr,
		ExpiredAt: end.Format(time.RFC3339),
	}

	fmt.Println("4")

	return res, nil
}

func (u *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(u.secretKey), nil
	})
}
