package main

import (
	router "GodFather-server/routers"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
