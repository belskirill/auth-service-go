package auth

import "proto-auth/internal/infra/repositories/postgres"

type Repository struct {
	db postgres.Executor
}

func NewRepository(db postgres.Executor) *Repository {
	return &Repository{
		db: db,
	}
}
