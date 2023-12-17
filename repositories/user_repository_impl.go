package repositories

import (
	"context"
	"database/sql"
	"log"
	"midtrans-go/helper"
	"midtrans-go/model/domain"
	"midtrans-go/utils"
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

		userRow.Created_at = utils.ParseTimestamp(createdAtStr)
		userRow.Updated_at = utils.ParseTimestamp(updatedAtStr)

		dataUser = append(dataUser, userRow)
	}

	return dataUser
}
