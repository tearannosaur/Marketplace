package database

import (
	"fmt"
	"log"
	"os"

	er "mp/internal/errors"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBModule struct {
	db *sqlx.DB
}

func DBInit() (*DBModule, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(er.GodotenvDownloadErr)
		return nil, err
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	fmt.Println("user,password,name")
	connection := fmt.Sprintf("user=%s password=%s dbname=%s  host=localhost port=5432 sslmode=disable", user, password, name)
	database, err := sqlx.Connect("postgres", connection)
	if err != nil {
		log.Println(er.DataBaseConnectionErr)
		return nil, err
	}
	return &DBModule{db: database}, nil
}

func (d *DBModule) DBClose() error {
	err := d.db.Close()
	if err != nil {
		log.Println(er.DataBaseCloseErr)
		return err
	}
	return nil
}
