package postgres

import (
	"backend/internal/database/postgres"
)

type Repository struct {
	client *postgres.Client
}

func NewRepository(client *postgres.Client) *Repository {
	return &Repository{client: client}
}
