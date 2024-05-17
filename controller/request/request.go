package request

import "ProductApp/service/dto"

type AddProductRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
	Store    string  `json:"store"`
}

func (addProductRequest AddProductRequest) ToModel() dto.ProductCreate {
	return dto.ProductCreate{
		Name:     addProductRequest.Name,
		Price:    addProductRequest.Price,
		Discount: addProductRequest.Discount,
		Store:    addProductRequest.Store,
	}
}
