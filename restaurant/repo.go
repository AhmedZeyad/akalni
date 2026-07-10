package restaurant

import (
	"log/slog"

	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type restaurantRepo struct {
	db *sqlx.DB
}

func NewRestaurantRepo(db *sqlx.DB) RestaurantRepo {
	return &restaurantRepo{db: db}
}

type RestaurantRepo interface {
	// admin method
	Create(restaurant Restaurant, createdBy int) error
	Update(restaurant Restaurant, updatedBy int) error
	UpdateStatus(restaurant Restaurant, updatedBy int) error
	Search(query string, limit, offset int, status bool) (shared.PaginationResponse, error)
	GetResProducts(ids ...int) ([]Product, error)

	GetByID(id int64) (Restaurant, error)
	// GetActive(limit, offset int) (shared.PaginationResponse, error)
}

// NOTE :admin
func (r *restaurantRepo) Create(restaurant Restaurant, createdBy int) error {
	_, err := r.db.Exec(`
		INSERT INTO restaurants
		(
			name,
			lon,
			lat,
			address,
			created_by,
			updated_at
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			now()
		)
		`,
		restaurant.Name,
		restaurant.Lon,
		restaurant.Lat,
		restaurant.Address,
		createdBy)
	if err != nil {
		slog.Error("failed to create restaurant", "error", err)
		return err
	}
	return nil
}
func (r *restaurantRepo) Update(restaurant Restaurant, updatedBy int) error {
	_, err := r.db.Exec(`
		UPDATE restaurants
		SET
			name = $2,
			lon = $3,
			lat = $4,
			address = $5,
			updated_by = $6,
			updated_at = now()
		WHERE id = $1
		`,
		restaurant.ID,
		restaurant.Name,
		restaurant.Lon,
		restaurant.Lat,
		restaurant.Address,
		updatedBy,
	)
	if err != nil {
		slog.Error("failed to update restaurant", "error", err)
		return err
	}
	return nil
}
func (r *restaurantRepo) UpdateStatus(restaurant Restaurant, updatedBy int) error {
	_, err := r.db.Exec(`
		UPDATE restaurants
		SET
			status = $2,
			updated_by = $3,
			updated_at = now()
		WHERE id = $1
		`,
		restaurant.ID,
		restaurant.Status,
		updatedBy,
	)
	if err != nil {
		slog.Error("failed to update restaurant", "error", err)
		return err
	}
	return nil
}
func (r *restaurantRepo) Search(term string, limit, offset int, status bool) (shared.PaginationResponse, error) {
	var restaurants []Restaurant
	var count int
	err := r.db.Select(&restaurants, `
			SELECT
				id,
				name,
				status,
				lon,
				lat,
				address
			FROM restaurants
			WHERE name ILIKE $3
			and ($4 = false or status = $4)
			and deleted_at is null
			GROUP BY  id
			order by id desc
			limit $1 offset $2
	`, limit, offset, "%"+term+"%", status)
	if err != nil {
		slog.Error("failed to search on restaurant", "error", err)
		return shared.PaginationResponse{}, err
	}
	err = r.db.Get(&count, `
			SELECT
				count(true)
			FROM restaurants
			WHERE name ILIKE $1
			and deleted_at is null
			and ($2 = false or status = $2)
	`, "%"+term+"%", status)
	if err != nil {
		slog.Error("failed to search on restaurant", "error", err)
		return shared.PaginationResponse{}, err
	}
	return shared.PaginationResponse{
		Result: restaurants,
		Count:  count,
	}, nil
}

// NOTE: shared
func (r *restaurantRepo) GetByID(id int64) (Restaurant, error) {
	var restaurant Restaurant
	err := r.db.Get(&restaurant, `
	SELECT
		id,
		name,
		status,
		lon,
		lat,
		address
	FROM restaurants
	WHERE id = $1
	and deleted_at is null
	`, id)
	if err != nil {
		slog.Error("failed to get restaurant", "error", err)
		return Restaurant{}, err
	}

	return restaurant, nil
}
func (r *restaurantRepo) GetProducts(id int) ([]Product, error) {
	var products []Product
	err := r.db.Select(&products, `
		SELECT
				p.id,
				p.name,
				price,
				rest_id,
				c.name as category,
				p.active
		FROM products p
		join categories c on p.category_id = c.id
		WHERE rest_id = $1
		and p.deleted_at is null
		and p.active
		group by c.id, p.id
		order by id desc
	`, id)
	if err != nil {
		slog.Error("failed to get products", "error", err)
		return nil, err
	}
	return products, nil
}

// NOTE: client
//
//	func (r *restaurantRepo) GetActive(limit, offset int) (shared.PaginationResponse, error) {
//		var restaurants []Restaurant
//		var count int
//		err := r.db.Select(&restaurants, `
//		SELECT
//			id,
//			name,
//			status,
//			lon,
//			lat,
//			address
//		FROM restaurants
//		WHERE status = true
//		and deleted_at is null
//		order by id desc
//		limit $1 offset $2
//		`, limit, offset)
//		if err != nil {
//			slog.Error("failed to get active restaurants", "error", err)
//			return shared.PaginationResponse{}, err
//		}
//		err = r.db.Get(&count, `
//		SELECT
//			count(*)
//		FROM restaurants
//		WHERE status = true
//		and deleted_at is null
//		`)
//		if err != nil {
//			slog.Error("failed to get active restaurants", "error", err)
//			return shared.PaginationResponse{}, err
//		}
//		return shared.PaginationResponse{
//			Result: restaurants,
//			Count:  count,
//		}, nil
//	}
func (r *restaurantRepo) GetResProducts(ids ...int) ([]Product, error) {
	var products []Product
	err := r.db.Select(&products, `
		SELECT
			p.id,
			p.name,
			price,
			rest_id,
			c.name as category,
			p.active
		FROM products p
		join categories c on p.category_id = c.id
		WHERE rest_id = any($1)
		and p.deleted_at is null		group by c.id, p.id
		order by id desc
	`, pq.Array(ids))
	if err != nil {
		slog.Error("failed to get products", "error", err)
		return nil, err
	}
	return products, nil
}
