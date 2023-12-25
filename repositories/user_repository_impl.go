package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"midtrans-go/helper"
	"midtrans-go/model/domain"
	"midtrans-go/utils"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) GetByUsername(ctx context.Context, tx *sql.Tx, user domain.User) []domain.User {

	SQL := "SELECT id, username, email, phone, password, created_at, updated_at FROM user WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Username)

	helper.PanicIfError("Cannot find data user", err)

	defer rows.Close()

	var dataUser []domain.User
	for rows.Next() {
		var userRow domain.User
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&userRow.Id, &userRow.Username, &userRow.Email, &userRow.Phone, &userRow.Password, &createdAtStr, &updatedAtStr)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
		}

		userRow.CreatedAt = utils.ParseTimestamp(createdAtStr)
		userRow.UpdatedAt = utils.ParseTimestamp(updatedAtStr)

		dataUser = append(dataUser, userRow)
	}

	return dataUser
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(username, email, phone, password) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Phone, user.Password)
	helper.PanicIfError("Error to create data", err)

	id, err := result.LastInsertId()
	helper.PanicIfError("Error to catch last id", err)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select id, username, email, phone, password, created_at, updated_at from user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError("Error to get data", err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.Password, &createdAtStr, &updatedAtStr)
		if err != nil {
			helper.PanicIfError("Error to convert data", err)
		}

		user.CreatedAt = utils.ParseTimestamp(createdAtStr)
		user.UpdatedAt = utils.ParseTimestamp(updatedAtStr)

		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, username, email, phone, password, created_at, updated_at FROM user"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError("Error to get data", err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.Password, &createdAtStr, &updatedAtStr)

		if err != nil {
			helper.PanicIfError("Error to convert data", err)
		}

		user.CreatedAt = utils.ParseTimestamp(createdAtStr)
		user.UpdatedAt = utils.ParseTimestamp(updatedAtStr)

		users = append(users, user)
	}
	return users
}
