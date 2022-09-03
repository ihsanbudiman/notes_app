package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ihsanbudiman/notes_app/domain"
)

func GenerateJwt(user domain.User) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	if user.ID == 0 {
		return "", errors.New("user id cannot be empty")
	}

	// Create claims while leaving out some of the optional fields
	claims := domain.TokenClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1200 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ValidateJwt(tokenString string) (*domain.TokenClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.TokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
