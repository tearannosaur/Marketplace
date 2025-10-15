package main

import (
	"log"
	db "mp/internal/database"
	er "mp/internal/errors"
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
	_ = repo
	defer mg.MigrationClose()
	defer database.DBClose()
	server.ServerInit()
}
