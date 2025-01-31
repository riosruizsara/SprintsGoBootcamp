package products

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
)

// ProductService is an interface that represents a service that can be used to interact with products
type ProductService interface {
	GetAll() (products map[int]products.Product, err error)
	GetByID(id models.ID) (product products.Product, err error)
	Create(product *products.Product) (addedProduct products.Product, err error)
	Update(id models.ID, patch products.ProductPatch) (updatedProduct products.Product, err error)
	Delete(id models.ID) (err error)
}
