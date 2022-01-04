package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID, role string) (string, string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	SecretKey string
}

type SignedDetails struct {
	Role   string
	UserID string
	jwt.StandardClaims
}

func NewService(secret string) Service {
	return &jwtService{
		SecretKey: secret,
	}
}

func (s *jwtService) GenerateToken(userID, role string) (string, string, error) {
	claims := &SignedDetails{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(s.SecretKey), nil
	})
	if err != nil {
		return jwtToken, err
	}
	return jwtToken, nil
}
