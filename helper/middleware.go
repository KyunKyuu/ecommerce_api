package helper

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(tokenUseCase TokenUseCase, allowedRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			token, err := ValidToken(authHeader)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}

			jwtToken, err := tokenUseCase.VerifyJWT(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			claims, ok := jwtToken.Claims.(*JwtCustomClaims)
			if !ok || !jwtToken.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token claims"})
			}

			fmt.Println("Role in token:", claims.Role) // Debug output

			if claims.Role != allowedRole {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "access forbidden: insufficient role"})
			}

			// Set claims in context for further use in handlers if needed
			c.Set("user", claims)
			return next(c)
		}
	}
}
