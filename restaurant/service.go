package restaurant

import (
	"errors"
	"log/slog"

	"github.com/AhmedZeyad/Akalni/customErrors"
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
func (s *restaurantService) SearchRestaurant(req SearchRequest) (PaginationResponse, error) {

	restaurants, err := s.repo.Search(req.Term, req.Limit, req.Offset)
	if err != nil {

		slog.Error("failed to update restaurant", "error", err)
		return PaginationResponse{}, err
	}

	return restaurants, nil
}
func (s *restaurantService) GetActiveRestaurant(req SearchRequest) (PaginationResponse, error) {

	restaurants, err := s.repo.GetActive(req.Limit, req.Offset)
	if err != nil {

		slog.Error("failed to update restaurant", "error", err)
		return PaginationResponse{}, err
	}

	return restaurants, nil
}
