package gin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Struct to represent the request payload
type RequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FindSupplierFromK2(db *gorm.DB) gin.HandlerFunc {
	return func(content *gin.Context) {
		url := "http://veggieandfood.k2dc.cloud/k2apiv2/token/login"

		// Prepare the request payload
		payload := RequestPayload{
			Username: "NguyenD",
			Password: "Veggie@2024",
		}
		// Convert the payload to JSON
		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// Create a new HTTP client
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		// Make a GET request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", "Bearer YOUR_TOKEN_HERE")

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		log.Println("resp.Body", resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		log.Println("body", body)
		// Decode the response JSON into the ApiResponse struct

		content.JSON(http.StatusOK, gin.H{
			"data": body,
		})
	}
}
