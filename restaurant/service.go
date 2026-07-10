package restaurant

import (
	"errors"
	"log/slog"

	"github.com/AhmedZeyad/Akalni/customErrors"
	"github.com/AhmedZeyad/Akalni/shared"
	"github.com/AhmedZeyad/Akalni/utils"
)

type restaurantService struct {
	repo RestaurantRepo
}

func NewRestaurantService(repo RestaurantRepo) restaurantService {
	return restaurantService{repo: repo}
}

func (s *restaurantService) CreateRestaurant(req RestaurantRequest, createdBy int) error {
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

func (s *restaurantService) UpdateRestaurant(req RestaurantRequest, updatedBy int) error {
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
func (s *restaurantService) UpdateRestaurantStatus(req RestaurantRequest, updatedBy int) error {
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
func (s *restaurantService) SearchRestaurant(req SearchRequest) (Restaurant, error) {

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
func (s *restaurantService) GetActiveRestaurant(req SearchRequest) (shared.PaginationResponse, error) {

	restaurants, err := s.repo.Search(req.Term, req.Limit, req.Offset, req.Status)
	if err != nil {

		slog.Error("failed to update restaurant", "error", err)
		return shared.PaginationResponse{}, err
	}

	return restaurants, nil
}
