package services

import (
	"dating-app/config"
	"dating-app/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret")

func CreateUser(username, email, password string) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)
	return user, result.Error
}

func AuthenticateUser(email, password string) (string, error) {
	var user models.User
	config.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid email or password")
	}

	status := false
	if user.Premium {
		status = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"status": status,
	})

	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

func ValidateToken(tokenString string) (uint, bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, false, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["userID"].(float64))
	status := claims["status"].(bool)
	return userID, status, nil
}
