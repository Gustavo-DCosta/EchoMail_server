package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	// Load .env file (to access the secret)
	godotenv.Load()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not found in environment")
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret used to verify the token
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token: signature or claims invalid")
	}

	// Optional: print or extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		fmt.Println("Token verified ✅")
		fmt.Println("Claims:")
		for k, v := range claims {
			fmt.Printf("- %s: %v\n", k, v)
		}
	}

	return token, nil
}

func CreateJWT(username string, role string) string {
	// Load .env file (contains secret)
	godotenv.Load()

	// Grab the secret from environment
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not found in environment")
	}

	// Define claims (you can customize this)
	claims := jwt.MapClaims{
		"sub":  username,
		"role": role,
		"iat":  time.Now().Unix(),                     // Issued at
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	// Create the token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		panic("failed to sign JWT: " + err.Error())
	}

	fmt.Println("JWT signed successfully:", signedToken)
	return signedToken
}
