package repository

import (
	repository "eva/internal/repository/users"
	"eva/pkg/logging"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repositories contain all repos
type Repositories struct {
	UserRepo repository.UserRepository
}

// Constructor
func NewRepositories(pgxpool *pgxpool.Pool, logger *logging.Logger) *Repositories {
	return &Repositories{
		UserRepo: repository.NewUserRepository(pgxpool, logger),
	}
}
