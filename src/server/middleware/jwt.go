package middleware

import (
	"log"
	"net/http"

	"auth/src/constanta/response"
	"auth/src/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(s service.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := response.BuildErrorResponse("Failed", "No token found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := s.ValidateToken(authHeader)
		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
		} else {
			log.Println(err)
			response := response.BuildErrorResponse("Token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
