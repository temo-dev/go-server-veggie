package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"server-veggie/common"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	//load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func Protected(content *gin.Context, method string) error {
	if content.Request.Method != method {
		return common.ErrMethodNotAllowed(errors.New("method is not allowed"))
	}
	return nil
}
