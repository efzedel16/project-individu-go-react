package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("silih_a3")

type AuthService interface {
	GenerateToken(userId int) (string, error)
	TokenValidation(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewAuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userId,
	}

	userTokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := userTokenData.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) TokenValidation(encodedToken string) (*jwt.Token, error) {
	userTokenData, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return SecretKey, nil
	})

	if err != nil {
		return userTokenData, err
	}

	return userTokenData, nil
}
