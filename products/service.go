package products

import "log/slog"

type productsService struct {
	ProductsRepo
}

func NewProductsService(repo ProductsRepo) *productsService {
	return &productsService{ProductsRepo: repo}
}
func (s *productsService) Create(createdBy int, product ProductsRequest) error {
	return s.ProductsRepo.Create(createdBy, product.ToData())
}
func (s *productsService) Update(updatedBy int, product ProductsRequest) error {
	return s.ProductsRepo.Update(updatedBy, product.ToData())
}
func (s *productsService) Search(params ProductsSearchParam) ([]ProductsData, int, error) {
	return s.ProductsRepo.Search(params)
}
func (s *productsService) Delete(id, deletedBy int) error {
	return s.ProductsRepo.Delete(id, deletedBy)
}
func (s *productsService) UpdateStatus(updatedBy int, product ProductsRequest) error {
	slog.Error("updating product status", "id", product.ID, "active", product.Active)

	return s.ProductsRepo.UpdateStatus(updatedBy, product.ToData())
}
