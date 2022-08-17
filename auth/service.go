package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int ) (string, error)
}

type jwtService struct {
}

var SECRETKEY = []byte("secretkeycrowdfunding") 

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {

	claim := jwt.MapClaims{}
	claim["user_id"] = userID //key user_id value nya adalah userID yang di dapat dari parameter
	

	// GENERATE TOKEN // BIKIN TOKEN  
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// SIGN TOKEN // TANDA TANGANI TOKEN
	signedToken, err := token.SignedString(SECRETKEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}