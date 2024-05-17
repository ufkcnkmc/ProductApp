package service

import (
	"ProductApp/domain"
	"ProductApp/persistence"
	"ProductApp/service/dto"
	"errors"
)

type IProductService interface {
	Add(productCreate dto.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float64) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) Add(productCreate dto.ProductCreate) error {
	validateErr := validateProductCreate(productCreate)
	if validateErr != nil {
		return validateErr
	}
	return productService.productRepository.CreateProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (productService *ProductService) DeleteById(productId int64) error {
	return productService.productRepository.DeleteById(productId)
}
func (productService *ProductService) GetById(productId int64) (domain.Product, error) {
	return productService.productRepository.GetById(productId)
}
func (productService *ProductService) UpdatePrice(productId int64, newPrice float64) error {
	return productService.productRepository.UpdatePriceById(productId, newPrice)
}
func (productService *ProductService) GetAllProducts() []domain.Product {
	return productService.productRepository.GetAllProducts()
}
func (productService *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	return productService.productRepository.GetAllProductsByStore(storeName)
}

func validateProductCreate(productCreate dto.ProductCreate) error {
	if productCreate.Discount > 70.0 {
		return errors.New("Discount can not be greater than 70")
	}
	return nil
}
