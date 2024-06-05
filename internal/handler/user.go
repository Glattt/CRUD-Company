package handler

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id,omitempty" db:"id"`
	Login    string `json:"login" db:"login"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Sex      string `json:"sex" db:"sex"`
}

func Users(ctx context.Context, conn *pgx.Conn) ([]User, error) {
	rows, err := conn.Query(ctx, `select * from users`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Sex, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u User) SignUp(ctx context.Context, conn *pgx.Conn) error {
	rows := conn.QueryRow(ctx, `select id from users where email = $1 or login = $2`, u.Email, u.Login)
	if err := rows.Scan(&u.Id); err == nil {
		return errors.New("user already exists")
	}

	password, err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, execErr := conn.Exec(ctx, `insert into users (login, email, password, sex) values ($1, $2, $3, $4)`, u.Login, u.Email, password, u.Sex)
	if execErr != nil {
		return errors.New("user already exists " + execErr.Error())
	}

	return nil
}

func (u User) SignIn(ctx context.Context, conn *pgx.Conn) error {
	var password string

	rows := conn.QueryRow(ctx, `select id, password from users where email = $1 or login = $2`, u.Email, u.Login)
	if rowsErr := rows.Scan(&u.Id, &password); rowsErr != nil {
		return errors.New("user not found " + rowsErr.Error())
	}

	if isCorrect := CheckPasswordHash(u.Password, password); !isCorrect {
		return errors.New("invalid password")
	}

	return nil
}

func (u User) ResetPassword(ctx context.Context, conn *pgx.Conn) error {
	password, err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, execErr := conn.Exec(ctx, `update users set password = $1 where (email = $2 or login = $3)`, password, u.Email, u.Login)
	if execErr != nil {
		return errors.New("user not found " + execErr.Error())
	}

	rows := conn.QueryRow(ctx, `select * from users where login = $1 or email = $2`, u.Login, u.Email)
	if scanErr := rows.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.Sex); scanErr != nil {
		return errors.New("user not found " + scanErr.Error())
	}

	return nil
}

func (u User) Update(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `update users set login = $1, email = $2, password = $3, sex = $4 where id = $5`, u.Login, u.Email, u.Password, u.Sex, u.Id)
	if err != nil {
		return errors.New("user not found")
	}

	rows := conn.QueryRow(ctx, `select * from users where id = $1`, u.Id)
	if scanErr := rows.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.Sex); scanErr != nil {
		return errors.New("user not found " + scanErr.Error())
	}

	return nil
}

func (u User) Delete(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `delete from users where id = $1`, u.Id)
	if err != nil {
		return errors.New("user not found " + err.Error())
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return err == nil
}
