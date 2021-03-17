package db

import (
	"GodFather-server/ent"
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB return *ent.Client
func GetDB(n string) (*ent.Client, error) {
	dsn := "root:admin1234@tcp(localhost:3306)/" + n + "?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}
