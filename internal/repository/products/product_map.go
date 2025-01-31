package products

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
)

// NewProductMap is a function that returns a new instance of ProductMap
func NewProductMap(db map[int]products.Product) *ProductMap {
	// default db
	defaultDb := make(map[int]products.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductMap{db: defaultDb}
}

// ProductMap is a struct that represents a products repository with a memory map implementation
type ProductMap struct {
	// db is a map of products
	db map[int]products.Product
}

// GetAll is a method that returns all the products
func (pm *ProductMap) GetAll() (products map[int]products.Product, err error) {
	products = pm.db
	return
}

// GetByID is a method that returns a product by its ID
func (pm *ProductMap) GetByID(id models.ID) (product products.Product, err error) {
	product, ok := pm.db[id.ID]
	if !ok {
		err = &errors.NotFoundError{Message: "Product not found"}
		return products.Product{}, err
	}
	return
}

// Create is a method that creates a new product, if a product with a matching product_code already exists, an error is returned
func (pm *ProductMap) Create(product *products.Product) (addedProduct products.Product, err error) {
	// Check the database for a matching product code
	for _, p := range pm.db {
		if p.ProductCode == product.ProductCode {
			err = &errors.DuplicateError{Message: "Product already exists"}
			return products.Product{}, err
		}
	}
	// No matching product code found, add the product to the database
	pm.db[product.ID.ID] = *product
	addedProduct = *product
	return
}

// Update is a method that updates a product, if a product with a matching ID does not exist, an error is returned
// if the product code is updated and a product with a matching product code already exists, an error is returned
func (pm *ProductMap) Update(product *products.Product) (updatedProduct products.Product, err error) {
	// Check if the product exists
	_, ok := pm.db[product.ID.ID]
	if !ok {
		err = &errors.NotFoundError{Message: "Product not found"}
		return products.Product{}, err
	}
	// Check the database for a matching product code
	for _, p := range pm.db {
		if p.ProductCode == product.ProductCode && p.ID.ID != product.ID.ID {
			err = &errors.DuplicateError{Message: "Product already exists"}
			return products.Product{}, err
		}
	}
	// No matching product code found, update the product in the database
	pm.db[product.ID.ID] = *product
	updatedProduct = *product
	return
}

// Delete is a method that deletes a product by its ID, if a product with a matching ID does not exist, an error is returned
func (pm *ProductMap) Delete(id models.ID) (err error) {
	// Check if the product exists
	_, ok := pm.db[id.ID]
	if !ok {
		err = &errors.NotFoundError{Message: "Product not found"}
		return err
	}
	// Delete the product from the database
	delete(pm.db, id.ID)
	return
}

func (pm *ProductMap) GenerateNewID() (id models.ID, err error) {
	// Find the highest ID
	highestID := 0
	for _, p := range pm.db {
		if p.ID.ID > highestID {
			highestID = p.ID.ID
		}
	}
	// Increment the highest ID by 1
	id.ID = highestID + 1
	return
}
