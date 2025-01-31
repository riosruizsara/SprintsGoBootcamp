package products

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
)

// ProductRepository is an interface that represents a repository that can be used to interact with products
type ProductRepository interface {
	GetAll() (products map[int]products.Product, err error)
	GetByID(id models.ID) (product products.Product, err error)
	Create(product *products.Product) (addedProduct products.Product, err error)
	Update(product *products.Product) (updatedProduct products.Product, err error)
	Delete(id models.ID) (err error)
	GenerateNewID() (id models.ID, err error)
}
