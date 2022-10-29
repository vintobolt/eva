package users

import (
	"context"
	"eva/internal/models"
	"eva/pkg/logging"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository interface
type UserRepository interface {
	GetExistUser(login string) (models.User, error)
	CreateUser(signUp models.SignUp) error
	UpdateUser(username, rolename string, active bool) error
	Deactivate(username string) error
}

// UserRepositoryImpl implements UserRepository interface
type UserRepositoryImpl struct {
	dbPool *pgxpool.Pool
	logger *logging.Logger
}

func NewUserRepository(dbPool *pgxpool.Pool, logger *logging.Logger) UserRepository {
	return &UserRepositoryImpl{
		dbPool: dbPool,
		logger: logger,
	}
}

func (r *UserRepositoryImpl) GetExistUser(login string) (models.User, error) {
	sql := fmt.Sprintf("SELECT passwd, rolename, fullname FROM users WHERE username='%s';", login)
	var passwd string
	var role string
	var fullname string
	err := r.dbPool.QueryRow(context.Background(), sql).Scan(&passwd, &role, &fullname)
	if err != nil {
		r.logger.Error(err)
	}
	user := models.User{Login: login, Passwd: passwd, Role: role, Fullname: fullname}
	return user, nil
}

// TODO:
func (r *UserRepositoryImpl) CreateUser(signUp models.SignUp) error {
	hashedPassword, err := hashPassword(signUp.Password)
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("INSERT INTO users (username, passwd, fullname) VALUES ('%s', '%s', '%s');", signUp.Username, hashedPassword, signUp.Fullname)
	r.logger.Debug(sql)
	_, err = r.dbPool.Exec(context.Background(), sql)
	if err != nil {
		r.logger.Error(err)
	}
	return nil
}

// TODO:
func (r *UserRepositoryImpl) Deactivate(username string) error {
	return nil
}

// TODO:
func (r *UserRepositoryImpl) UpdateUser(username, rolename string, active bool) error {
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
