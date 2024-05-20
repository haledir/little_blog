package services

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

type AuthService struct {
	DB *sql.DB
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) AuthenticateUser(username, password string) (*User, error) {
	var user User
	err := s.DB.QueryRow("select ID, USERNAME, PASSWORD from USERS where username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, errors.New("user and password combination does not exist")
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("user and password combination does not exist")
	}
	return &user, nil
}
