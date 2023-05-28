package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserID   string `json:"userid"`
	Email    string `json:"email"`
	Metadata string `json:"metadata"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(key_sign, user_id, email, metadata, issuer string, expired int) (string, error) {
	signingKey := []byte(key_sign)
	// Create the claims
	claims := CustomClaims{
		user_id,
		email,
		metadata,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expired) * time.Second)),
			Issuer:    issuer,
		},
	}
	//fmt.Printf("%+v\r\n",claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	res, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return res, nil
}

func TokenExpiredTime(key, token_string string) float64 {
	var claims CustomClaims
	_, err := jwt.ParseWithClaims(token_string, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err == nil {
		return 0
	}
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		//tm := time.Unix(claims.ExpiresAt, 0)
		return time.Now().Sub(claims.ExpiresAt.Time).Seconds()
	}
	return 0
}

func VerifyJWTToken(key, token_string string) (*CustomClaims, error) {
	if key == "" {
		return nil, errors.New("_KEY_IS_EMPTY_")
	}
	if token_string == "" {
		return nil, errors.New("_TOKEN_IS_EMPTY_")
	}
	token, err := jwt.ParseWithClaims(token_string, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if token == nil {
		return nil, errors.New("PARSE_TOKEN_ERROR")
	}
	claim, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claim, nil
	}
	return nil, err
}
