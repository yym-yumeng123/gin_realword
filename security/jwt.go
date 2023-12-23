package security

import (
	"gin_realword/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GeneratorJWT 生成 jwt
func GeneratorJWT(username, email string) (string, error) {
	hmacSampleSecret := []byte(config.GetSecret())
	now := time.Now()
	tokenDuration := 24 * time.Hour
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]string{
			"username": username,
			"email":    email,
		},
		"iat": now.Unix(),
		"exp": now.Add(tokenDuration).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func VerifyJWT(token string) (*jwt.MapClaims, bool, error) {
	var claim jwt.MapClaims
	// Parse the token
	claims, err := jwt.ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetSecret()), nil
	})

	if err != nil {
		return nil, false, err
	}

	if claims.Valid {
		return &claim, true, nil
	}
	return nil, true, nil
}
