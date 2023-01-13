package Utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token string
}

func GenerateToken(user_id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userid"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenstr Token) (string, error) {
	token, err := jwt.Parse(tokenstr.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte("secret"), nil

	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id := claims["userid"].(string)
		return user_id, nil
	}
	return "", nil
}
