package restaurant

import (
	"errors"
	"log/slog"

	"github.com/AhmedZeyad/Akalni/customErrors"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/AhmedZeyad/Akalni/utils"
)

type RestaurantService struct {
	repo RestaurantRepo
}

func NewRestaurantService(repo RestaurantRepo) RestaurantService {
	return RestaurantService{repo: repo}
}

func (s *RestaurantService) CreateRestaurant(req RestaurantRequest, createdBy int) error {
	err := req.Validate()
	if err != nil {
		slog.Error("validation failed", "error", err)
		return err
	}
	if err := s.repo.Create(req.FromRestaurantRequest(), createdBy); err != nil {
		slog.Error("failed to create restaurant", "error", err)
		return err
	}

	return nil
}

func (s *RestaurantService) UpdateRestaurant(req RestaurantRequest, updatedBy int) error {
	err := req.Validate()
	if err != nil {
		slog.Error("validation failed", "error", err)
		return err
	}
	if err := s.repo.Update(req.FromRestaurantRequest(), updatedBy); err != nil {
		slog.Error("failed to update restaurant", "error", err)
		return err
	}

	return nil
}
func (s *RestaurantService) UpdateRestaurantStatus(req RestaurantRequest, updatedBy int) error {
	if utils.IsEmpty(req.ID) {
		slog.Error("id is required")
		return errors.New(customErrors.VALIDATION_MISSING_REQUIRED_FIELD + ": id")
	}
	if err := s.repo.UpdateStatus(req.FromRestaurantRequest(), updatedBy); err != nil {
		slog.Error("failed to update restaurant", "error", err)
		return err
	}

	return nil
}
func (s *RestaurantService) SearchRestaurant(req SearchRequest) (Restaurant, error) {

	restaurants, err := s.repo.GetByID(req.ID)
	if err != nil {

		slog.Error("failed to update restaurant", "error", err)
		return Restaurant{}, err
	}

	restaurants.Products, err = s.repo.GetResProducts(int(restaurants.ID))
	if err != nil {
		slog.Error("failed to get products", "error", err)
		return Restaurant{}, err
	}
	return restaurants, nil
}
func (s *RestaurantService) GetActiveRestaurant(req SearchRequest) (shared.PaginationResponse, error) {

	restaurants, err := s.repo.Search(req.Term, req.Limit, req.Offset, req.Status)
	if err != nil {

		slog.Error("failed to update restaurant", "error", err)
		return shared.PaginationResponse{}, err
	}

	return restaurants, nil
}
func (s *RestaurantService) GetRestaurantByID(resId int, productsIDs []int) (Restaurant, error) {

	restaurants, err := s.repo.GetByID(int64(resId))
	if err != nil {

		slog.Error("failed to update restaurant", "error", err)
		return Restaurant{}, err
	}

	restaurants.Products, err = s.repo.GetProductsByID(resId, productsIDs)
	if err != nil {
		slog.Error("failed to get products", "error", err)
		return Restaurant{}, err
	}
	return restaurants, nil
}
