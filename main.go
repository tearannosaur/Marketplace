package main

import (
	"log"
	db "mp/internal/database"
	"mp/internal/server"

	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	db, err := db.DBInit(user, name, password)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.DBClose()
	server.ServerInit()
}
