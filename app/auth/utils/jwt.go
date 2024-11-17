package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

var Secret = []byte("tikmall")

// GenToken generates a jwt token with userID and role
// if duration is 0, the token will not expire, used for refresh token
func GenToken(userID int64, duration time.Duration) (string, error) {
	jwtCliams := new(jwt.StandardClaims)
	if duration != 0 {
		jwtCliams.ExpiresAt = time.Now().Add(duration).Unix()
	}

	c := Claims{
		userID,
		*jwtCliams,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

func ParseToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}
	return 0, err
}
