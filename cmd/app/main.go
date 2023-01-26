package main

import (
	"fmt"
	"log"
	config "test-module/internal/configs"
	"test-module/internal/database/psql"
)

func main() {
	config := config.Config{}
	config.LoadConfig("D:\\projects\\GO  learning\\vs go tests\\configs\\config.yml")

	db, err := psql.Connect(config)
	if err != nil {
		log.Printf("db.Connect err   #%v ", err)
	}
	rows, _ := db.Queryx("SELECT * FROM cars")

	s := rows
	fmt.Printf("%i", s)
}
