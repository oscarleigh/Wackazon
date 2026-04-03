package store

import (
	"Server/internal/model"
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) Create(ctx context.Context, email, firstName, lastName, pwdHash string) (string, error) {
	id := uuid.New().String()
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO users (id, email, firstName, lastName, joinDate, pwdHash) VALUES (?, ?, ?, ?, ?, ?)",
		id, email, firstName, lastName, time.Now().Format("2006-01-02"), pwdHash,
	)
	return id, err
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := s.db.QueryRowContext(ctx,
		"SELECT id, email, firstName, lastName, joinDate, pwdHash, streetAddress, city, country, postCode FROM users WHERE email = ?",
		strings.ToLower(email),
	).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.JoinDate, &user.PwdHash, &user.StreetAddress, &user.City, &user.Country, &user.PostCode)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserStore) GetByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := s.db.QueryRowContext(ctx,
		"SELECT id, email, firstName, lastName, joinDate, pwdHash, streetAddress, city, country, postCode FROM users WHERE id = ?",
		strings.ToLower(id),
	).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.JoinDate, &user.PwdHash, &user.StreetAddress, &user.City, &user.Country, &user.PostCode)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserStore) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)",
		strings.ToLower(email),
	).Scan(&exists)
	return exists, err
}
