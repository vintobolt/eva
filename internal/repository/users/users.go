package users

import (
	"context"
	"eva/internal/models"
	"eva/pkg/logging"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	dbPool *pgxpool.Pool
	logger *logging.Logger
}

func NewUserRepository(dbPool *pgxpool.Pool, logger *logging.Logger) *UserRepository {
	return &UserRepository{
		dbPool: dbPool,
		logger: logger,
	}
}

/*
func (r *UserRepository) GetRoleByLogin(login string) (models.UserRole, error) {
	sql := fmt.Sprintf("SELECT rolename, fullname FROM users WHERE login='%s';", login)
	var rolename string
	var fullname string
	err := r.dbPool.QueryRow(context.Background(), sql).Scan(&rolename, &fullname)
	if err != nil {
		fmt.Println(err)
	}
	//user := models.UserModel{Login: login, Role: rolename, Fullname: fullname}
	user := models.UserRole{Login: login, Role: rolename}
	fmt.Printf("%+v\n", user)
	return user, nil
} */

// TODO:
func (r *UserRepository) GetExistUser(login string) (models.User, error) {
	sql := fmt.Sprintf("SELECT passwd, role, fullname FROM users WHERE login='%s';", login)
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
func (r *UserRepository) CreateUser() {

}

// TODO:
func (r *UserRepository) DeleteUser() {

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
