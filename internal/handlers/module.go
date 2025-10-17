package handlers

import "mp/internal/repository"

type HandlerModule struct {
	repo *repository.RepositoryModule
}

func HandlerModuleInit(r *repository.RepositoryModule) *HandlerModule {
	return &HandlerModule{repo: r}
}
