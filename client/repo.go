package client

import (
	"database/sql"
	"errors"

	"github.com/ba7rIbrahim/Akalni/customeErrors"
	"github.com/jmoiron/sqlx"
)

type ClientRepo interface {
	// Define methods for interacting with the API
	GetByID(id int) (client Client, err error)
}
type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) ClientRepo {
	return &clientRepo{db: db}
}

func (r *clientRepo) GetByID(id int) (client Client, err error) {

	err = r.db.Get(&client, `
		select
	
			first_name,
			last_name,
			email
		--	phone_number
		from clients
		where id=$1
		`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return client, errors.New(customeErrors.AUTH_USER_NOT_FOUND)
		}
		return client, err
	}
	return
}
