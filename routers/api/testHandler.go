package testHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	fmt.Println("test")
	c.JSON(200, gin.H{
		"status": "success",
	})
}
