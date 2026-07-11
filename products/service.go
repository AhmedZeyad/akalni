package products

import "log/slog"

type ProductsService struct {
	ProductsRepo
}

func NewProductsService(repo ProductsRepo) *ProductsService {
	return &ProductsService{ProductsRepo: repo}
}
func (s *ProductsService) Create(createdBy int, product ProductsRequest) error {
	return s.ProductsRepo.Create(createdBy, product.ToData())
}
func (s *ProductsService) Update(updatedBy int, product ProductsRequest) error {
	return s.ProductsRepo.Update(updatedBy, product.ToData())
}
func (s *ProductsService) Search(params ProductsSearchParam) ([]ProductsData, int, error) {
	return s.ProductsRepo.Search(params)
}
func (s *ProductsService) Delete(id, deletedBy int) error {
	return s.ProductsRepo.Delete(id, deletedBy)
}
func (s *ProductsService) UpdateStatus(updatedBy int, product ProductsRequest) error {
	slog.Error("updating product status", "id", product.ID, "active", product.Active)

	return s.ProductsRepo.UpdateStatus(updatedBy, product.ToData())
}
func (s *ProductsService) GetProductByID(restID int64, ids ...int64) ([]ProductsData, error) {
	return nil, nil
}
