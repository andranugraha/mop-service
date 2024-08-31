package strings

import (
	"fmt"
	"time"

	"github.com/empnefsi/mop-service/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID     uint64 `json:"user_id"`
	Email      string `json:"email"`
	MerchantID uint64 `json:"merchant_id"`
	jwt.RegisteredClaims
}

func GenerateToken(claims Claims) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.GetTokenExpiry()) * time.Second)

	claims.Issuer = "mop-service"
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.Subject = fmt.Sprintf("%d", claims.UserID)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	secret := []byte(config.GetTokenSecret())
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key
		return []byte(config.GetTokenSecret()), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
