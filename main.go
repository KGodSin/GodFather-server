package main

import (
	"GodFather-server/pkg/db"
	router "GodFather-server/routers"
	"fmt"
)

func main() {
	client, err := db.GetDB("team")
	if err != nil {
		panic(err)
	}

	fmt.Println(client)
	r := router.InitRouter()
	r.Run(":8080")
}
