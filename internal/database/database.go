package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBModule struct {
	db *sqlx.DB
}

func DBInit(u, n, p string) (*DBModule, error) {
	user := u
	dbname := n
	password := p
	fmt.Println(user, dbname, password)
	connection := fmt.Sprintf("user=%s password=%s dbname=%s  host=localhost port=5432 sslmode=disable", user, password, dbname)
	fmt.Println(connection)
	database, err := sqlx.Connect("postgres", connection)
	if err != nil {
		return nil, err
	}
	return &DBModule{db: database}, nil
}
func (d *DBModule) DBClose() error {
	err := d.db.Close()
	if err != nil {
		return err
	}
	return nil
}
