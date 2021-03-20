package v1handler

import (
	"GodFather-server/ent"
	"GodFather-server/pkg/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	fmt.Println("signup 들어옴")
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	client := c.MustGet("client").(*ent.Client)
	if !user.IsValid(client) {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "회원가입 성공!",
		})
	}
}

func Signin(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	client := c.MustGet("client").(*ent.Client)
	if user.Signin(client) {
		jwtToken, err := user.GetToken()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Authentication failed",
			})
			return
		}
		c.SetCookie("access-token", jwtToken, 1800, "", "", false, false)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "ok",
		})
	}
}

func TokenTest(c *gin.Context) {

}
