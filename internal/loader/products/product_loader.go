package products

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
)

// ProductLoader is an interface that represents the loader for products
type ProductLoader interface {
	// Load is a method that loads the products
	Load() (v map[int]products.Product, err error)
}
