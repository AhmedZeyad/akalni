package repo

import (
	"context"

	"github.com/ba7rIbrahim/Akalni/Models"
	"github.com/jmoiron/sqlx"
)

type AuthRepo interface {
	Create(context context.Context, client *Models.Client) (err error)
}
type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepo {
	return &authRepo{db: db}
}

// TODO : CREATE NEW CLIENT
func (ar *authRepo) Create(context context.Context, client *Models.Client) (err error) {
	rows, err := ar.db.NamedQuery(`
	INSERT INTO clients (email, password ,first_name, last_name,phone_number,created_at)
	VALUES (:email, :password, :first_name, :last_name, :phone_number,  NOW()) RETURNING id
	`, client)
	if err != nil {
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

// TODO : CHECK IF CLIENT EXISTS
// TODO : GET CLIENT BY EMAIL
// TODO : GET CLIENT BY ID
//
// TODO : UPDATE CLIENT LOW PRIO
// TODO : DELETE CLIENT LOW PRIO
