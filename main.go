package main

import (
	"log"
	db "mp/internal/database"
	er "mp/internal/errors"
	h "mp/internal/handlers"
	r "mp/internal/repository"
	"mp/internal/server"
)

func main() {

	database, err := db.DBInit()
	if err != nil {
		log.Fatalln(er.DataBaseConnectionErr, err)
	}
	log.Println("Successfully database connection")

	mg, err := db.Migration()
	if err != nil {
		log.Fatalln(err)
	}
	repo := r.RepositoryModuleInit(database)
	handler := h.HandlerModuleInit(repo)
	_ = repo
	defer mg.MigrationClose()
	defer database.DBClose()
	server.ServerInit(handler)
}
