package restaurant

import (
	"errors"

	"github.com/AhmedZeyad/Akalni/customErrors"
)

type Restaurant struct {
	ID      int64   `db:"id"`
	Name    string  `db:"name"`
	Status  bool    `db:"status"`
	Lon     float64 `db:"lon"`
	Lat     float64 `db:"lat"`
	Address string  `db:"address"`
}

type RestaurantRequest struct {
	ID      int64   `json:"id" form:"id"`
	Name    string  `json:"name" form:"name"`
	Status  bool    `json:"status" form:"status"`
	Lon     float64 `json:"lon" form:"lon"`
	Lat     float64 `json:"lat" form:"lat"`
	Address string  `json:"address" form:"address"`
}

type SearchRequest struct {
	ID     int64  `json:"id" form:"id"`
	Term   string `json:"term" form:"term"`
	Limit  int    `json:"limit" form:"limit"`
	Offset int    `json:"offset" form:"offset"`
}
type PaginationResponse struct {
	Result any `json:"result"`
	Count  int `json:"count"`
}
type RestaurantResponse struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Status  bool    `json:"status"`
	Lon     float64 `json:"lon"`
	Lat     float64 `json:"lat"`
	Address string  `json:"address"`
}

func (r *RestaurantRequest) Validate() error {
	if r.Name == "" {
		return errors.New(customErrors.VALIDATION_NAME_REQUIRED)
	}

	if r.Lon == 0 {
		return errors.New(customErrors.VALIDATION_LON_REQUIRED)
	}
	if r.Lat == 0 {
		return errors.New(customErrors.VALIDATION_LAT_REQUIRED)
	}
	if r.Address == "" {
		return errors.New(customErrors.VALIDATION_ADDRESS_REQUIRED)
	}
	return nil
}

func (r *RestaurantRequest) FromRestaurantRequest() Restaurant {
	return Restaurant{
		Name:    r.Name,
		Status:  r.Status,
		Lon:     r.Lon,
		Lat:     r.Lat,
		Address: r.Address,
	}
}

func (r *RestaurantResponse) FromRestaurant(restaurant *Restaurant) {
	r.ID = restaurant.ID
	r.Name = restaurant.Name
	r.Status = restaurant.Status
	r.Lon = restaurant.Lon
	r.Lat = restaurant.Lat
	r.Address = restaurant.Address
}

type Product struct {
	ID           int64   `db:"id"`
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Price        float64 `db:"price"`
	RestaurantID int64   `db:"rest_id"`
	Category     string  `db:"category"`
	Active       bool    `db:"active"`
}

type ProductRequest struct {
	Name         string  `json:"name" form:"name"`
	Description  string  `json:"description" form:"description"`
	Price        float64 `json:"price" form:"price"`
	RestaurantID int64   `json:"restaurant_id" form:"restaurant_id"`
}
