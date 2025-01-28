package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Not authorized",
			})
		}
		tokenStr := cookie.Value
		if tokenStr == "" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Not authorized",
			})
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Invalid token signature method")
			}
			key := []byte(os.Getenv("SECRET_ACCESS"))
			return key, nil
		})
		if token.Claims.Valid() != nil {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid token ,expired",
			})
		}
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid token",
			})
		}
		return next(ctx)
	}
}
