package database

import (
	"fmt"
	"log"
	er "mp/internal/errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type MigrationModule struct {
	mg *migrate.Migrate
}

func Migration() (*MigrationModule, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, name)
	m, err := migrate.New("file://internal/migrations", dbURL)
	if err != nil {
		log.Println(er.MigrationInitErrr)
		return nil, err
	}
	err = m.Up()
	if err == migrate.ErrNoChange {
		log.Println("Migrations no change")
	}
	if err != nil && err != migrate.ErrNoChange {
		log.Println(er.MigrationUpErr)
		return nil, err
	}
	log.Println("Successfully migrations applied")
	return &MigrationModule{mg: m}, nil
}

func (m *MigrationModule) MigrationClose() error {
	err, dberr := m.mg.Close()
	if err != nil || dberr != nil {
		log.Println(er.MigrationCloserErr)
		err = fmt.Errorf(er.MigrationCloserErr+"source:%v,%v", err, dberr)
		return err
	}
	return nil
}

func (m *MigrationModule) DatabaseDrop() error {
	err := m.mg.Drop()
	if err != nil {
		log.Println(er.MigrationDropErr)
		return err
	}
	log.Println("Succesfully migration drop")
	return nil
}

func (m *MigrationModule) DatabaseDown() error {
	err := m.mg.Down()
	if err != nil {
		log.Println(er.MigrationDownErr)
		return err
	}
	return nil
}
