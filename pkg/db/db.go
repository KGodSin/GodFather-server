package db

import (
	"GodFather-server/ent"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB return *ent.Client
func GetDB(n string) (*ent.Client, error) {
	dsn := "root:admin1234@tcp(localhost:3306)/" + n + "?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
