package products

import (
	"encoding/json"
	"os"

	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
)

// NewProductJSONFile is a function that returns a new instance of ProductJSONFile
func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{
		path: path,
	}
}

// ProductJSONFile is a struct that implements the ProductLoader interface
type ProductJSONFile struct {
	// path is the path to the file that contains the Products in JSON format
	path string
}

// Load is a method that loads the products from a JSON file and returns a map of products
func (l *ProductJSONFile) Load() (v map[int]products.Product, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var productsJSON []products.ProductDTO
	err = json.NewDecoder(file).Decode(&productsJSON)
	if err != nil {
		if err.Error() == "EOF" {
			// Handle empty file by returning an empty map
			return make(map[int]products.Product), nil
		}
		return
	}

	// serialize products
	v = make(map[int]products.Product)
	for _, product := range productsJSON {
		productModel, err := product.ToModel(nil)
		if err != nil {
			return nil, err
		}
		v[productModel.ID.ID] = productModel
	}

	return
}
