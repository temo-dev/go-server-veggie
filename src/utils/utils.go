package utils

import (
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	//load .env file
	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// generate-token
var jwtSecret = []byte(GoDotEnvVariable("SECRET"))

func GenerateToken(name_account string) (string, error) {
	claims := jwt.MapClaims{
		"name_account": name_account,
		"exp":          time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func SendEmail(smtpHost string, smtpPort string, senderEmail string, senderPassword string, recipientEmail string, subject string, body string) error {
	//define header email
	headers := "From: " + senderEmail + "\r\n" +
		"To: " + recipientEmail + "\r\n" +
		"Subject: " + subject + "\r\n\r\n"
	// Gộp headers và body
	message := headers + body

	// Kết nối đến SMTP server
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
