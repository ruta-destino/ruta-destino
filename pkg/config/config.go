package config

import (
	"fmt"
	"os"
)

func GetConnectionString() string {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	return fmt.Sprintf("user=%s password=%s dbname=%s host=db sslmode=disable", user, password, dbname)
}
