package models

import (
	db "mp/internal/database"
	"sync"
)

type MainModule struct {
	DataBase *db.DBModule
	mux      sync.RWMutex
}
