package router

import (
	"GodFather-server/pkg/model"
	v1handler "GodFather-server/routers/api/v1"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("access-token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Authentication failed",
			})
			c.Abort()
			return
		}

		tokenStr := token.Value
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "token is None",
			})
			c.Abort()
			return
		}

		claims := &model.Claims{}
		_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  "token is expired",
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "Authentication failed",
			})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.POST("/signup", v1handler.Signup)
		api.POST("/signin", v1handler.Signin)
	}
	// test := r.Group("/test").Use(TokenChecker())
	// {
	// 	test.POST("/token_test", v1handler.TokenTest)
	// }

	return r
}
