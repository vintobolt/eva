package users

import (
	"context"
	"eva/internal/models"
	"eva/pkg/logging"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepository interface
type UserRepository interface {
	GetExistUser(login string) (models.User, error)
	CreateUser(signUp models.SignUp) error
	UpdateUser(username, rolename string, active bool) error
	ActivateUser(username string) error
	Deactivate(username string) error
	FindCredsWithUsername(username string) (*models.SignIn, error)
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
	sql := fmt.Sprintf("SELECT passwd, rolename, fullname, active FROM users WHERE username='%s';", login)
	var passwd string
	var role string
	var fullname string
	var active bool
	err := r.dbPool.QueryRow(context.Background(), sql).Scan(&passwd, &role, &fullname, &active)
	if err != nil {
		r.logger.Error(err)
	}
	user := models.User{Login: login, Passwd: passwd, Role: role, Fullname: fullname}
	return user, nil
}

// TODO:
func (r *UserRepositoryImpl) CreateUser(signUp models.SignUp) error {
	//hashedPassword, err := utils.EncryptPassword(signUp.Password)
	//hashedPassword, err := hashPassword(signUp.Password)
	sql := fmt.Sprintf("INSERT INTO users (username, passwd, fullname) VALUES ('%s', '%s', '%s');", signUp.Username, signUp.Password, signUp.Fullname)
	r.logger.Debug(sql)
	_, err := r.dbPool.Exec(context.Background(), sql)
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
func (r *UserRepositoryImpl) ActivateUser(username string) error {
	return nil
}

// TODO:
func (r *UserRepositoryImpl) UpdateUser(username, rolename string, active bool) error {
	return nil
}

func (r *UserRepositoryImpl) FindCredsWithUsername(username string) (*models.SignIn, error) {
	sql := fmt.Sprintf("SELECT username, passwd FROM users WHERE username='%s';", username)
	var passwd string
	var user string
	err := r.dbPool.QueryRow(context.Background(), sql).Scan(&user, &passwd)
	if err != nil {
		//r.logger.Error(err)
		return &models.SignIn{}, err
	}
	userCreds := models.SignIn{Username: user, Password: passwd}
	return &userCreds, nil
}

/*
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
*/
