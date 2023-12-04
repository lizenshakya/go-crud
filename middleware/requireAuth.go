package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lizenshakya/go-crud/initializers"
	"github.com/lizenshakya/go-crud/models"
)

func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	fmt.Println(authHeader)
	// Check if the header is empty
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString := authHeader

	// Replace with your secret key
	secretKey := []byte(os.Getenv("SECRET"))

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return nil, fmt.Errorf("Invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if token.Valid {
		fmt.Println("Token is valid")
		claims, ok := token.Claims.(jwt.MapClaims)
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("Failed to parse claims")
		} else {
			var user models.User
			initializers.DB.First(&user, claims["sub"])
			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Set("user", user)
			fmt.Println("User ID:", claims["sub"])
			// Add more claim validations here if needed
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("Token is invalid")
	}

	c.Next()
}
