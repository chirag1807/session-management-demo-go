package repository

import (
	"context"
	"database/sql"
	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"sessionmanagement/error"
	"sessionmanagement/utils"

	"github.com/jackc/pgx/v5"
)

type AuthRepository interface {
	UserLogin(user request.User) (response.User, error)
}

type authRepository struct{
	pgx *pgx.Conn
}

func NewAuthRepo(pgx *pgx.Conn) AuthRepository {
	return authRepository{
		pgx: pgx,
	}
}

func (a authRepository) UserLogin(user request.User) (response.User, error) {
	var dbUser response.User
	row := a.pgx.QueryRow(context.Background(), `SELECT id, name, bio, email, password, image, isadmin FROM users WHERE email = $1`, user.Email)
	err := row.Scan(&dbUser.ID, &dbUser.Name, &dbUser.Bio, &dbUser.Email, &dbUser.Password, &dbUser.Image, &dbUser.IsAdmin)

	if err == sql.ErrNoRows {
		return response.User{}, errorhandling.NoUserFound
	}

	passwordMatched := utils.VerifyPassword(user.Password, dbUser.Password)
	if !passwordMatched {
		return response.User{}, errorhandling.PasswordNotMatch
	}

	return dbUser, nil
}