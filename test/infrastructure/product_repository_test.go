package infrastructure

import (
	"ProductApp/common/postgresql"
	"ProductApp/domain"
	"ProductApp/persistence"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		"localhost",
		"6432",
		"postgres",
		"postgres",
		"productapp",
		"10",
		"30s",
	})
	productRepository = persistence.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)

}
func setup(ctx context.Context, dbpool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbpool)
}
func clear(ctx context.Context, dbpool *pgxpool.Pool) {
	TruncateTestData(ctx, dbpool)
}
func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetAllProductsByStoreName(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}

	t.Run("GetAllProductsByStoreName", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestCreateProduct(t *testing.T) {
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "Kitap",
			Price:    10,
			Discount: 10.0,
			Store:    "Kumcu Kırtasiye",
		},
	}
	newProduct := domain.Product{
		Name:     "Kitap",
		Price:    10,
		Discount: 10.0,
		Store:    "Kumcu Kırtasiye",
	}
	t.Run("CreateProduct", func(t *testing.T) {
		productRepository.CreateProduct(newProduct)
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}
func TestGetProductById(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetById(1)
		_, err := productRepository.GetById(5)
		assert.Equal(t, domain.Product{
			1,
			"AirFryer",
			3000.0,
			22.0,
			"ABC TECH",
		}, actualProduct)
		assert.Equal(t, "Product not found with id 5", err.Error())
	})
	clear(ctx, dbPool)
}

func TestDeleteProductById(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("DeleteProductById", func(t *testing.T) {
		productRepository.DeleteById(1)
		//allProducts := productRepository.GetAllProducts()
		//assert.Equal(t, 3, len(allProducts))
		_, err := productRepository.GetById(1)
		assert.Equal(t, "Product not found with id 1", err.Error())
	})
	clear(ctx, dbPool)
}
func TestUpdatePriceById(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("UpdatePriceById", func(t *testing.T) {
		productBeforeUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, float64(3000.0), productBeforeUpdate.Price)
		productRepository.UpdatePriceById(1, 2500.0)
		productAfterUpdate, _ := productRepository.GetById(1)
		assert.Equal(t, float64(2500.0), productAfterUpdate.Price)
	})
	clear(ctx, dbPool)
}
