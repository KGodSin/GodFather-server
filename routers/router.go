package router

import (
	"GodFather-server/pkg/db"
	v1handler "GodFather-server/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func dbGetter() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := db.GetDB("team")
		if err != nil {
			panic(err)
		}
		c.Set("client", client)
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api").Use(dbGetter())
	{
		api.GET("/test", v1handler.Test)
		api.POST("/signup", v1handler.Signup)
	}
	return r
}
