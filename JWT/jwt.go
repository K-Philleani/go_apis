package JWT

import (
	"github.com/dgrijalva/jwt-go"
	"go_apis/models"
	"time"
)

type Claims struct {
	UserId string
	jwt.StandardClaims
}

var jwtKey = []byte("Kphilleani")


func ReleaseToken(user models.Account) (string, error){
	expirationTime := time.Now().Add(30 * time.Second)
	claims := &Claims{
		UserId: user.UserAccount,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "kphilleani.cn",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

