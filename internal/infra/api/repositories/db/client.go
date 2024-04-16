package db

import (
	"errors"

	"github.com/CValier/PruebaGO/internal/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type authRepository struct {
	// TODO: Add any dependency or DB connection
	users map[string]*entity.User
}

func NewAuthRepository() *authRepository {
	// TODO: Initialize any DB connection or instance for client
	return &authRepository{
		users: make(map[string]*entity.User),
	}
}

func (ar *authRepository) RegisterUser(user *entity.User) error {
	// TODO: Add the logic to save the user in the DB
	_, exist := ar.users[user.Email]
	if exist {
		return errors.New("An account is already registered with this email.")
	}
	ar.users[user.Email] = user
	return nil
}

func (ar *authRepository) LoginUser(email, password string) (*entity.User, error) {
	user, exist := ar.users[email]
	if !exist {
		return nil, errors.New("User not found.")
	}

	// Comparing the password with the hash
	isValidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if isValidPassword != nil {
		return nil, errors.New("Your password and email do not match. Try Again.")
	}

	return user, nil
}
