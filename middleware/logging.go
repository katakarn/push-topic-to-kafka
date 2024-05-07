package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			// Read the request body
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": "Error reading request body"})
				return
			}

			// Restore the request body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			// Attempt to hide sensitive information if content is JSON
			var jsonData map[string]interface{}
			if err := json.Unmarshal(body, &jsonData); err == nil {
				sensitiveFields := []string{"password", "pin", "idcard", "phoneNumber"}
				sensitivePatterns := []string{"Token"} // Use patterns to catch fields containing "Token"

				// Hide sensitive fields
				for _, field := range sensitiveFields {
					if _, exists := jsonData[field]; exists {
						jsonData[field] = "****"
					}
				}

				// Hide fields containing sensitive patterns
				for key := range jsonData {
					for _, pattern := range sensitivePatterns {
						if strings.Contains(strings.ToLower(key), strings.ToLower(pattern)) {
							jsonData[key] = "****"
							break
						}
					}
				}

				// Convert the modified map back to JSON
				modifiedBody, err := json.Marshal(jsonData)
				if err != nil {
					log.Println("Error re-encoding JSON")
				} else {
					body = modifiedBody
				}
			}

			var logstr string
			email, emailExists := c.Get("email")
			studentCode, studentCodeExists := c.Get("studentCode")
			logstr += fmt.Sprintf("[%s] %s - ", c.Request.Method, c.Request.RequestURI)
			if emailExists {
				emailStr, _ := email.(string)
				logstr += fmt.Sprintf("Email: %s, ", emailStr)
			}
			if studentCodeExists {
				studentCodeStr, _ := studentCode.(string)
				logstr += fmt.Sprintf("Student Code: %s, ", studentCodeStr)
			}
			logstr += fmt.Sprintf("Request Body: %s", string(body))
			log.Print(logstr)
		} else {
			var logstr string
			email, emailExists := c.Get("email")
			logstr += fmt.Sprintf("[%s] %s - ", c.Request.Method, c.Request.RequestURI)
			if emailExists {
				emailStr, _ := email.(string)
				logstr += fmt.Sprintf("Email: %s", emailStr)
			}
			log.Print(logstr)
		}

		c.Next()
	}
}
