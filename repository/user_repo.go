package repository

import (
	"database/sql"

	"eCommerceAPI/models"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) GetByEmail(email string) (models.User, error) {
	row := r.DB.QueryRow("SELECT id, email, password_hash, email_verification, created_at FROM users WHERE email = $1", email)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.EmailVerified, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepo) NewUser(user models.User) (models.User, error) {
	err := r.DB.QueryRow(
		"INSERT INTO users (email, password_hash, email_verification) VALUES ($1, $2, false) RETURNING id, created_at",
		user.Email, user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
