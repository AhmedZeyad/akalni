package products

type ProductsData struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	RestID      int     `db:"rest_id"`
	Price       float64 `db:"price"`
	Active      bool    `db:"active"`
	Category_id int     `db:"category_id"`
}
type ProductsRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	RestID      int     `json:"rest_id"`
	Price       float64 `json:"price"`
	Category_id int     `json:"category_id"`
	Active      bool    `json:"active"`
}
type ProductsStatus struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}
type ProductsSearchParam struct {
	ID     int    `form:"id"`
	Name   string `form:"name"`
	RestID int    `form:"rest_id"`
}
type ProductsResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	RestID      int     `json:"rest_id"`
	Price       float64 `json:"price"`
	Active      bool    `json:"active"`
	Category_id int     `json:"category_id"`
}

func (p ProductsRequest) ToData() ProductsData {
	return ProductsData{
		ID:          p.ID,
		Name:        p.Name,
		RestID:      p.RestID,
		Price:       p.Price,
		Category_id: p.Category_id,
		Active:      p.Active,
	}
}
