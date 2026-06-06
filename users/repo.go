package users

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(User) error
	GetByEmail(string) (*User, error)
	ResetPassword(string, string, int64) error
	GetPassword(int64) (string, error)
	GetPasswordByEmail(string) (string, error)
}
type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}

}
func (u *userRepo) Create(user User) error {
	_, err := u.db.Exec(`
	insert into users
		(
		name,
		email,
		password,
		created_by,
		created_at
		)
		values
		(
		$1,
		$2,
		$3,
		$4,
		now()
		)
		`,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedBy,
	)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return err
	}
	return nil

}

func (u *userRepo) GetByEmail(email string) (*User, error) {
	var user User
	err := u.db.Get(&user, `
		select
			id,
			name,
			email,
			password,
			status,
			created_at,
			created_by
		from users
		where email = $1
		limit 1`,
		email)
	if err != nil {
		slog.Error("failed to get user by email", "error", err)
		return nil, err
	}
	return &user, nil

}

func (u *userRepo) ResetPassword(email, password string, updatedBy int64) error {
	_, err := u.db.Exec(`
		update users
			set password = $1,
			updated_at = now(),
			updated_by = $3
		where email = $2`,
		password,
		email,
		updatedBy,
	)
	if err != nil {
		slog.Error("failed to reset password", "error", err)
		return err
	}
	return nil
}
func (u *userRepo) GetPassword(id int64) (string, error) {
	var password string
	err := u.db.Get(&password, `
		select
			password
		from users
		where id = $1
		limit 1`,
		id)
	if err != nil {
		slog.Error("failed to get password", "error", err)
		return "", err
	}
	return password, nil
}

func (u *userRepo) GetPasswordByEmail(email string) (string, error) {
	var password string
	err := u.db.Get(&password, `
		select
			password
		from users
		where email = $1
		limit 1`,
		email)
	if err != nil {
		slog.Error("failed to get password", "error", err)
		return "", err
	}
	return password, nil
}
