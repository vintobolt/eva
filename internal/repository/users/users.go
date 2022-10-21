package users

import (
	"context"
	"eva/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	dbPool *pgxpool.Pool
}

func NewUsersRepository(dbPool *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{
		dbPool: dbPool,
	}
}

func (r *UsersRepository) GetRoleByLogin(login string) (models.UserRole, error) {
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
}
