package client

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"

	customErrors "github.com/AhmedZeyad/Akalni/customErrors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ClientRepo interface {
	// Define methods for interacting with the API
	// GetByID(id int64) (client Client, err error)
	Create(context context.Context, client *Client) (err error)
	GetByEmail(context context.Context, email string) (client Client, err error)
	GetByID(context context.Context, id int64) (client Client, err error)
}
type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) ClientRepo {
	return &clientRepo{db: db}
}

// func (r *clientRepo) GetByID(id int64) (client Client, err error) {

// 	err = r.db.Get(&client, `
// 		select

// 			first_name,
// 			last_name,
// 			email
// 		--	phone_number
// 		from clients
// 		where id=$1
// 		`, id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return client, errors.New(customErrors.AUTH_USER_NOT_FOUND)
// 		}
// 		return client, err
// 	}
// 	return
// }

// TODO : CREATE NEW CLIENT
func (ar *clientRepo) Create(context context.Context, client *Client) (err error) {
	rows, err := ar.db.NamedQuery(`
	INSERT INTO clients (email, password ,first_name, last_name,phone_number,created_at)
	VALUES (:email, :password, :first_name, :last_name, :phone_number,  NOW()) RETURNING id
	`, client)
	if err != nil {

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
		}
		if strings.Contains(err.Error(), "23505") && strings.Contains(err.Error(), "email_unique") {

			return errors.New(customErrors.VALIDATION_EMAIL_DUPLICATE)
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
func (ar *clientRepo) GetByEmail(context context.Context, email string) (client Client, err error) {
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
		// TODO return error
		if err == sql.ErrNoRows {

			return client, errors.New(customErrors.AUTH_USER_NOT_FOUND)
		}
		return
	}
	if client.EmailVerifiedAt == nil {
		return Client{}, errors.New(customErrors.AUTH_EMAIL_NOT_VERIFIED)
	}
	return
}

// TODO : GET CLIENT BY ID
func (ar *clientRepo) GetByID(context context.Context, id int64) (client Client, err error) {
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
			return client, errors.New(customErrors.AUTH_USER_NOT_FOUND)
		}
		return
	}
	return
}

// TODO : UPDATE CLIENT LOW PRIO
// TODO : DELETE CLIENT LOW PRIO
