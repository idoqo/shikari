package db

import (
	"database/sql"
	"gitlab.com/idoko/shikari/models"
)

func (db Database) SaveUser(user *models.User) error {
	var id int
	query := "INSERT INTO users(email, password) VALUES($1, $2) RETURNING id"
	err := db.Conn.QueryRow(query, user.Email, user.Password).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (db Database) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT id, email, password FROM users WHERE email = $1"
	row := db.Conn.QueryRow(query, email)
	switch err := row.Scan(&user.ID, &user.Email, &user.Password); err {
	case sql.ErrNoRows:
		return user, ErrNoRecord
	default:
		return user, err
	}
}