package service

import (
	"ProductApp/domain"
	"ProductApp/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func (fakeRepository *FakeProductRepository) CreateProduct(product domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func (fakeRepository *FakeProductRepository) UpdatePriceById(productId int64, price float64) error {
	//TODO implement me
	panic("implement me")
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}

func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	// todo: your turn
	return []domain.Product{}
}

func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}
func (fakeRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	// todo: your turn
	return domain.Product{}, nil
}
func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {
	return nil
}

func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float64) error {
	return nil
}
