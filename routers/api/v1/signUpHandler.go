package v1handler

import (
	"GodFather-server/ent"
	"GodFather-server/pkg/model"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user model.User
	c.Bind(&user)

	client := c.MustGet("client").(*ent.Client)
	if !user.IsValid(client) {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "이미 존재하는 아이디입니다.",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "회원가입 성공!",
		})
	}
}
