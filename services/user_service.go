package services

import (
	"database/sql"
	"fmt"
	"net/mail"

	"eCommerceAPI/models"
	"eCommerceAPI/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepo
}

func (s *UserService) RegisterUser(email string, password string) (models.User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid email address")
	}

	_, err = s.Repo.GetByEmail(email)
	if err == nil {
		return models.User{}, fmt.Errorf("email already registered")
	} else if err != sql.ErrNoRows {
		return models.User{}, fmt.Errorf("database error checking email: %v", err)
	}

	if len(password) <= 8 {
		return models.User{}, fmt.Errorf("password must be longer than 8 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:         email,
		PasswordHash:  string(hashedPassword),
		EmailVerified: false,
	}

	createdUser, err := s.Repo.NewUser(user)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}
