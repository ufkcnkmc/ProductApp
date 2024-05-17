package persistence

import (
	"ProductApp/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product //çoklu product dönebileceği için slice olarak tanımlıyoruz
	CreateProduct(product domain.Product) error
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePriceById(productId int64, price float64) error
}
type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository { //burada implamente ediyoruz
	return &ProductRepository{dbPool: dbPool}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "SELECT * FROM products")

	if err != nil {
		log.Error("Failed to get products from database %v", err)
		return []domain.Product{}
	}
	var products = []domain.Product{}
	var id int64
	var name string
	var price float64
	var discount float64
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{id, name, price, discount, store})
	}
	return products
}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNamesql := `SELECT * FROM products WHERE store = $1`
	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNamesql, storeName)

	if err != nil {
		log.Error("Failed to get products from database %v", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) CreateProduct(product domain.Product) error {
	ctx := context.Background()
	insert_sql := `Insert into products (name,price,discount,store) values ($1, $2, $3, $4)`
	newProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Error("Failed to insert product from database %v", err)
		return err
	}
	log.Info(fmt.Sprintf("Inserted product %v successfully", newProduct))
	return nil
}

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	getByIdSql := `Select * from products where id = $1`
	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)
	var id int64
	var name string
	var price float64
	var discount float64
	var store string
	scanErr := queryRow.Scan(&id, &name, &price, &discount, &store)
	if scanErr != nil && scanErr.Error() == "no rows in result set" {
		return domain.Product{}, errors.New(fmt.Sprintf("Product not found with id %d", productId))
	}
	if scanErr != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Failed to get product from database %v", scanErr))
	}
	return domain.Product{id, name, price, discount, store}, nil
}

func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()
	_, getErr := productRepository.GetById(productId)
	if getErr != nil {
		return errors.New(fmt.Sprintf("Failed to get product from database %v", getErr))
	}
	deleteSql := `Delete from products where id = $1`
	_, err := productRepository.dbPool.Exec(ctx, deleteSql, productId)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to deleteing products with id : %d", productId))
	}
	log.Info("Deleted product successfully")
	return nil
}

func (productRepository *ProductRepository) UpdatePriceById(productId int64, newPrice float64) error {
	ctx := context.Background()

	updateSql := `Update products set price = $1 where id = $2`

	_, err := productRepository.dbPool.Exec(ctx, updateSql, newPrice, productId)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while updating product with id : %d", productId))
	}
	log.Info("Product %d price updated with new price %v", productId, newPrice)
	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float64
	var discount float64
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{id, name, price, discount, store})
	}
	return products
}
