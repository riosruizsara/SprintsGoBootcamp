package products

import (
	products_repo "github.com/almarino_meli/grupo-5-wave-15/internal/repository/products"
	"github.com/almarino_meli/grupo-5-wave-15/internal/repository/sellers"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
)

type ProductServiceDefault struct {
	rp          products_repo.ProductRepository
	sellersRepo sellers.SellerRepository
}

func NewProductServiceDefault(rp products_repo.ProductRepository, sellersRepo sellers.SellerRepository) ProductServiceDefault {
	return ProductServiceDefault{rp: rp, sellersRepo: sellersRepo}
}

func (ps ProductServiceDefault) GetAll() (products map[int]products.Product, err error) {
	products, err = ps.rp.GetAll()
	if err != nil {
		return nil, &errors.UnknownError{Message: "Unknown error fetching products"}
	}
	if len(products) == 0 {
		return nil, &errors.NotFoundError{Message: "No products were found"}
	}
	return
}

func (ps ProductServiceDefault) GetByID(id models.ID) (product products.Product, err error) {
	product, err = ps.rp.GetByID(id)
	if err != nil {
		return products.Product{}, err
	}
	return
}

func (ps ProductServiceDefault) Create(product *products.Product) (addedProduct products.Product, err error) {
	// Check if the seller_id exists (if provided)
	if product.SellerID.ID != 0 {
		_, err = ps.sellersRepo.GetById(product.SellerID.ID)
		if err != nil {
			return products.Product{}, &errors.NotFoundError{Message: "Seller not found"}
		}
	}

	// Generate new unique ID
	newID, err := ps.rp.GenerateNewID()
	if err != nil { // currently database in memory can't return an error
		return products.Product{}, &errors.UnknownError{Message: "Unknown error generating new ID"}
	}
	product.ID = newID

	addedProduct, err = ps.rp.Create(product)
	if err != nil {
		return products.Product{}, err
	}
	return
}

func (ps ProductServiceDefault) Update(id models.ID, patch products.ProductPatch) (updatedProduct products.Product, err error) {
	// Fetch the existing product
	existingProduct, err := ps.rp.GetByID(id)
	if err != nil {
		return products.Product{}, err
	}

	// Apply the patch to the existing product
	applyPatch(&existingProduct, patch)

	// Check if the seller_id exists (if it is being updated)
	if patch.SellerID != nil {
		_, err = ps.sellersRepo.GetById(*patch.SellerID)
		if err != nil {
			return products.Product{}, &errors.NotFoundError{Message: "Seller not found"}
		}
	}

	// Check for duplicate product_code
	if patch.ProductCode != nil && *patch.ProductCode != existingProduct.ProductCode {
		foundProducts, err := ps.rp.GetAll()
		if err != nil {
			return products.Product{}, err
		}
		for _, product := range foundProducts {
			if product.ProductCode == *patch.ProductCode {
				return products.Product{}, &errors.DuplicateError{Message: "Error updating product: ProductCode must be unique"}
			}
		}
	}

	// Update the product in the repository
	updatedProduct, err = ps.rp.Update(&existingProduct)
	if err != nil {
		return products.Product{}, err
	}

	return updatedProduct, nil
}

func (ps ProductServiceDefault) Delete(id models.ID) (err error) {
	err = ps.rp.Delete(id)
	if err != nil {
		return err
	}
	return
}
