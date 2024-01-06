package middleware

import (
    "net/http"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt"
)

const (
    Bearer = len("Bearer ")
)

func AuthMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(http.StatusUnauthorized).SendString("Unauthorized")
    }

    bearerPrefix := "Bearer "
    if !strings.HasPrefix(authHeader, bearerPrefix) {
        return c.Status(http.StatusUnauthorized).SendString("Unauthorized: Invalid Bearer token")
    }

    tokenString := authHeader[len(bearerPrefix):]
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Anda perlu mengganti secretKey dengan kunci rahasia yang sesuai
        secretKey := []byte("secret")
        // Validate the token signing method and verify the secret key
        return secretKey, nil
    })

    if err != nil || !token.Valid {
        return c.Status(http.StatusForbidden).SendString("Forbidden")
    }

    return c.Next()
}
