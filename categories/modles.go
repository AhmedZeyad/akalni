package categories

type CategoriesData struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Active bool   `db:"active"`
}
type CategoriesRequest struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
type CategoriesStatus struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}
type CategoriesSearchParam struct {
	ID         int    `form:"id"`
	Name       string `form:"name"`
	PageNumber int    `form:"page_number"`
	PageSize   int    `form:"page_size"`
}
type CategoriesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (cd CategoriesRequest) ToData() CategoriesData {
	return CategoriesData{
		ID:     cd.ID,
		Name:   cd.Name,
		Active: cd.Active,
	}
}
