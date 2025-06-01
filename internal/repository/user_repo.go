package repository

import (
	"database/sql"
	"restApiCase/internal/models"
)

func CreateUser(db *sql.DB, name string) (models.User, error) {
	var user models.User
	err := db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id, name", name).Scan(&user.ID, &user.Name)
	return user, err
}

func GetUser(db *sql.DB, id int) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	return user, err
}
