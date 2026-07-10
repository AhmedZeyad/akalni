package categories

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type CategoriesRepo interface {
	Create(createdBy int, categories CategoriesData) error
	Update(updatedBy int, categories CategoriesData) error
	Search(params CategoriesSearchParam) ([]CategoriesData, int, error)
	Delete(id, deletedBy int) error
	UpdateStatus(updatedBy int, categories CategoriesData) error
}

type categoriesRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) CategoriesRepo {
	return &categoriesRepo{db: db}
}
func (cr *categoriesRepo) Create(createdBy int, categories CategoriesData) error {
	_, err := cr.db.Exec(`
	INSERT INTO categories
		(
			name,
			created_by,
			created_at
		)
	values
		(
			$1,
			$2,
			now()
		)
			`, categories.Name, createdBy)
	if err != nil {
		slog.Error("failed to create categories", "error", err)
		return err
	}
	return nil
}
func (cr *categoriesRepo) Update(updatedBy int, categories CategoriesData) error {
	_, err := cr.db.Exec(`
		UPDATE categories
		set
			name = $1,
			updated_by = $2,
			updated_at = now()
		where id = $3
		and deleted_at is null
		`, categories.Name, updatedBy, categories.ID)
	if err != nil {
		slog.Error("failed to update categories", "error", err)
		return err
	}
	return nil
}
func (cr *categoriesRepo) UpdateStatus(updatedBy int, categories CategoriesData) error {
	_, err := cr.db.Exec(`
		UPDATE categories
		set
			active = $2,
			updated_by = $3,
			updated_at = now()
		where id = $1
		and deleted_at is null
		`, categories.ID, categories.Active, updatedBy)
	if err != nil {
		slog.Error("failed to update categories status", "error", err)
		return err
	}
	return nil
}
func (cr *categoriesRepo) Search(params CategoriesSearchParam) ([]CategoriesData, int, error) {
	var (
		result []CategoriesData
		count  int
	)
	err := cr.db.Select(&result, `
		SELECT
			id,
			name,
			active
		FROM categories
		WHERE deleted_at is null
		and ($1='' or name ilike $1)
		limit $2 offset $3
	`,
		"%"+params.Name+"%",
		params.PageSize,
		(params.PageNumber-1)*params.PageSize,
	)
	if err != nil {
		slog.Error("failed to search categories", "error", err, "name", params.Name)
		return nil, 0, err
	}
	err = cr.db.Get(&count, `
		SELECT
			count(true)
		FROM categories
		WHERE
			deleted_at is null
			and ($1='' or name ilike $1)
	`, "%"+params.Name+"%")
	if err != nil {
		slog.Error("failed to count categories", "error", err)
		return nil, 0, err
	}
	return result, count, nil
}
func (cr *categoriesRepo) Delete(id, deletedBy int) error {
	_, err := cr.db.Exec(`
		UPDATE categories
		set
			deleted_by = $2,
			deleted_at = now()
		where id = $1
		and deleted_at is null
		`, id, deletedBy)
	if err != nil {
		slog.Error("failed to delete categories", "error", err)
		return err
	}
	return nil
}
