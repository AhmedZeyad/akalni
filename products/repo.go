package products

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type ProductsRepo interface {
	Create(createdBy int, product ProductsData) error
	Update(updatedBy int, product ProductsData) error
	Search(params ProductsSearchParam) ([]ProductsData, int, error) //by restaurants id
	Delete(id, deletedBy int) error
	UpdateStatus(updatedBy int, product ProductsData) error
}

type productsRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) ProductsRepo {
	return &productsRepo{db: db}
}
func (pr *productsRepo) Create(createdBy int, product ProductsData) error {
	_, err := pr.db.Exec(`
		INSERT INTO products
			(
			name,
			rest_id,
			price,
			category_id,
			created_by,
			created_at

			)
		values
			(
			$1,
			$2,
			$3,
			$4,
			$5,
			now()
			)
		`, product.Name, product.RestID, product.Price, product.Category_id, createdBy)
	if err != nil {
		slog.Error("failed to create products", "error", err)
		return err

	}

	return nil
}
func (pr *productsRepo) Update(updatedBy int, product ProductsData) error {
	_, err := pr.db.Exec(`
		UPDATE products
		set
			name = $1,
			updated_by = $2,
			updated_at = now()
		where id = $3
		and deleted_at is null
		`, product.Name, updatedBy, product.ID)
	if err != nil {
		slog.Error("failed to update products", "error", err)
		return err
	}

	return nil
}
func (pr *productsRepo) UpdateStatus(updatedBy int, product ProductsData) error {
	_, err := pr.db.Exec(`
		UPDATE products
		set
			active = $1,
			updated_by = $2,
			updated_at = now()
		where id = $3
		and deleted_at is null
		`, product.Active, updatedBy, product.ID)
	if err != nil {
		slog.Error("failed to update products status", "error", err)
		return err
	}

	return nil
}
func (pr *productsRepo) Search(params ProductsSearchParam) ([]ProductsData, int, error) {
	var result []ProductsData
	var count int
	err := pr.db.Select(&result, `
		SELECT
			id,
			name,
			rest_id,
			price,
			active,
			category_id
		FROM products
		where
			deleted_at is null
			and rest_id=$1
			and ($2='' or name like $2)
	`, params.RestID, params.Name)
	if err != nil {
		slog.Error("failed to search products", "error", err)
		return nil, 0, err
	}
	err = pr.db.Get(&count, `
		SELECT
			count(true)
		FROM products
		where
			deleted_at is null
			and rest_id=$1
			and ($2='' or name like $2)
	`, params.RestID, params.Name)
	if err != nil {
		slog.Error("failed to get products count", "error", err)
		return nil, 0, err
	}

	return result, count, nil
}
func (pr *productsRepo) Delete(id, deletedBy int) error {
	_, err := pr.db.Exec(`
		UPDATE products
		set
			deleted_by = $1,
			deleted_at = now()
		where id = $2

		`, id)
	if err != nil {
		slog.Error("failed to delete products", "error", err)
		return err
	}

	return nil
}
