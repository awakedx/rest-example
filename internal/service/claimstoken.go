package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ExtractClaims(ctx echo.Context) (jwt.MapClaims, error) {
	tokenStr, err := TokenStr(ctx.Cookie("token"))
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_ACCESS")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed extract claims")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, nil
}

func TokenStr(cookie *http.Cookie, err error) (string, error) {
	if cookie.Name != "token" {
		return "", fmt.Errorf("Unauthorized")
	}
	return cookie.Value, nil
}

func GetUserIdClaims(ctx echo.Context) (uuid.UUID, error) {
	c, err := ExtractClaims(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	userId, err := uuid.Parse(c["userId"].(string))
	return userId, nil
}
