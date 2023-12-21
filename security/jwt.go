package security

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GeneratorJWT 生成 jwt
func GeneratorJWT(username, email string) (string, error) {
	hmacSampleSecret := []byte("secret")
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
