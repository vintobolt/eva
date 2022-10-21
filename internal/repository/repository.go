package repository

import (
	repository "eva/internal/repository/users"
	"eva/pkg/logging"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	UserRepository *repository.UsersRepository
}

func NewRepository(pgxpool *pgxpool.Pool, logger *logging.Logger) *Repository {
	return &Repository{
		UserRepository: repository.NewUsersRepository(pgxpool),
	}
}
