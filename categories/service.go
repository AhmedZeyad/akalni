package categories

type categoriesService struct {
	CategoriesRepo
}

func NewCategoriesService(cr CategoriesRepo) categoriesService {
	return categoriesService{CategoriesRepo: cr}
}

func (s *categoriesService) Create(createdBy int, categories CategoriesRequest) error {

	return s.CategoriesRepo.Create(createdBy, categories.ToData())
}
func (s *categoriesService) Update(updatedBy int, categories CategoriesRequest) error {
	return s.CategoriesRepo.Update(updatedBy, categories.ToData())
}
func (s *categoriesService) Search(params CategoriesSearchParam) ([]CategoriesData, int, error) {
	return s.CategoriesRepo.Search(params)
}
func (s *categoriesService) Delete(id, deletedBy int) error {
	return s.CategoriesRepo.Delete(id, deletedBy)
}
func (s *categoriesService) UpdateStatus(updatedBy int, categories CategoriesRequest) error {
	return s.CategoriesRepo.UpdateStatus(updatedBy, categories.ToData())
}
