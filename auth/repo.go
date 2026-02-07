package auth

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"

	"github.com/ba7rIbrahim/Akalni/customeErrors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AuthRepo interface {
	Create(context context.Context, client *Client) (err error)
	GetByEmail(context context.Context, email string) (client Client, err error)
	GetByID(context context.Context, id int) (client Client, err error)
}
type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepo {
	return &authRepo{db: db}
}

// TODO : CREATE NEW CLIENT
func (ar *authRepo) Create(context context.Context, client *Client) (err error) {
	rows, err := ar.db.NamedQuery(`
	INSERT INTO clients (email, password ,first_name, last_name,phone_number,created_at)
	VALUES (:email, :password, :first_name, :last_name, :phone_number,  NOW()) RETURNING id
	`, client)
	if err != nil {

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
		}
		if strings.Contains(err.Error(), "23505") && strings.Contains(err.Error(), "email_unique") {

			return errors.New(customeErrors.VALIDATION_EMAIL_DUPLICATE)
		}

		slog.Error("feild to add client", "error", err)
		return err

	}

	if rows.Next() {
		err = rows.Scan(&client.ID)
		if err != nil {
			return err
		}
	}
	defer rows.Close()

	return err
}

// DONE : CHECK IF CLIENT EXISTS
// TODO : GET CLIENT BY EMAIL
func (ar *authRepo) GetByEmail(context context.Context, email string) (client Client, err error) {
	err = ar.db.Get(&client,
		`
	select
		id,
		first_name,
		last_name,
		phone_number,
		email,
		is_email_verified,
		email_verified_at,
		password
	from  clients
	where email=$1
		`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return client, errors.New(customeErrors.AUTH_USER_NOT_FOUND)
		}
		return
	}
	return
}

// TODO : GET CLIENT BY ID
func (ar *authRepo) GetByID(context context.Context, id int) (client Client, err error) {
	err = ar.db.Get(&client,
		`
	select
		id,
		first_name,
		last_name,
		phone_number,
		email,
		is_email_verified,
		email_verified_at,
		password
	from  clients
	where id=$1
		`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return client, errors.New(customeErrors.AUTH_USER_NOT_FOUND)
		}
		return
	}
	return
}

// TODO : UPDATE CLIENT LOW PRIO
// TODO : DELETE CLIENT LOW PRIO
