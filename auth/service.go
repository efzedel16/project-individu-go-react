package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {

}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	panic("implement me")
}

var SECRETKEY = []byte("5!l!h_43")

func NewService() *jwtService {
	return&jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRETKEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}