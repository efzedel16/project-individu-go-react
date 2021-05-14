package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userId int, fullName string) (string, error)
}

type jwtService struct {
}

var SecretKey = []byte("silih_a3")

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
