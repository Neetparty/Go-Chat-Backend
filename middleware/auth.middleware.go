package middleware

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Role string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)

func Permit(allowedRoles []Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract token from request (replace with your logic)
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return echo.ErrUnauthorized
			}

			// Parse the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Replace with your secret key logic
				return []byte("secret_key"), nil
			})
			if err != nil {
				if errors.Is(err, jwt.ErrSignatureInvalid) {
					return echo.ErrUnauthorized
				}
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			// Extract user roles from claims (replace with your claims structure)
			claims := token.Claims.(jwt.MapClaims)
			userRoles, ok := claims["roles"].([]interface{})
			if !ok {
				return echo.ErrUnauthorized
			}

			var roles []Role
			for _, role := range userRoles {
				roles = append(roles, Role(role.(string)))
			}

			// Check if user has allowed role
			if !HasAllowedRole(roles, allowedRoles) {
				return echo.ErrUnauthorized
			}

			return next(c)
		}
	}
}

func HasAllowedRole(userRoles []Role, allowedRoles []Role) bool {
	for _, userRole := range userRoles {
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				return true
			}
		}
	}
	return false
}
