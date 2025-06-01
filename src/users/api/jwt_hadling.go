package api

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"path/filepath"
)

type UsernamePassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandlers struct {
	jwtPrivate *rsa.PrivateKey
	jwtPublic  *rsa.PublicKey
}

func InitAuthHandlers() *AuthHandlers {
	privateFile, err := filepath.Abs(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Error while reading private key: %v", err)
	}
	publicFile, err := filepath.Abs(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		log.Fatalf("Error while reading public key: %v", err)
	}

	private, err := os.ReadFile(privateFile)
	if err != nil {
		log.Fatalf("Error while reading private key: %v", err)
	}
	public, err := os.ReadFile(publicFile)
	if err != nil {
		log.Fatal(err)
	}

	jwtPrivate, err := jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		log.Fatal(err)
	}
	jwtPublic, err := jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		log.Fatal(err)
	}
	return &AuthHandlers{
		jwtPrivate: jwtPrivate,
		jwtPublic:  jwtPublic,
	}
}

func CreateToken(ah *AuthHandlers, username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(ah.jwtPrivate)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(ah *AuthHandlers, tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return ah.jwtPublic, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}
	return claims["username"].(string), nil
}
