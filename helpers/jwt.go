package helpers

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GetSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return secretKey
}

func GenerateToken(id uint, email string) string {
	secretKey := GetSecretKey()
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1),
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return err.Error()
	}

	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	secretKey := GetSecretKey()

	headerToken := ctx.GetHeader("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errors.New("Sign in to proceed")
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Sign in to proceed")
		}
		return ([]byte(secretKey)), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("the token has no claims")
	}

	expClaims, exist := claims["exp"]
	if !exist {
		return nil, errors.New("expire claim is missing")
	}

	expString, ok := expClaims.(string)
	if !ok {
		return nil, errors.New("Expire claim is not a valid type")
	}

	expTime, err := time.Parse(time.RFC3339, expString)
	if err != nil {
		return nil, errors.New("Error parsing expiration time")
	}

	if time.Now().After(expTime) {
		return nil, errors.New("Token is expired")
	}
	return claims, nil
}
