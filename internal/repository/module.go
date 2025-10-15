package repository

import (
	db "mp/internal/database"
)

type RepositoryModule struct {
	database *db.DBModule
}

func RepositoryModuleInit(d *db.DBModule) *RepositoryModule {
	return &RepositoryModule{database: d}
}
