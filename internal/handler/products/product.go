package products

import (
	"encoding/json"
	"errors"
	"net/http"
	"sort"
	"strconv"

	products_service "github.com/almarino_meli/grupo-5-wave-15/internal/service/products"
	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// Creates a new instance of ProductController, and associates the passed service with it
func NewProductController(sv products_service.ProductService) *ProductController {
	return &ProductController{sv: sv}
}

type ProductController struct {
	sv products_service.ProductService
}

// Returns all the found products
func (ct *ProductController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productsFound, err := ct.sv.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error fetching products: No products were found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error fetching products: Server ran into an error",
				})
			}
			return
		}

		// Extract keys and sort them
		keys := make([]int, 0, len(productsFound))
		for key := range productsFound {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		// Map models to DTOs in sorted order
		productDTOs := make([]products.ProductDTO, 0, len(productsFound))
		for _, key := range keys {
			productDTOs = append(productDTOs, productsFound[key].ToDTO())
		}

		// Return the products as JSON
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    productDTOs,
		})
	}
}

// Returns a single product by its ID
func (ct *ProductController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		parsedID, err := parseID(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Error fetching product: Invalid input data",
			})
			return
		}

		product, err := ct.sv.GetByID(parsedID)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error fetching product: No product was found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error fetching product: Server ran into an error",
				})
			}
			return
		}

		// Map model to DTO
		foundProductDTO := product.ToDTO()

		// Return the product as JSON
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    foundProductDTO,
		})
	}
}

// Creates a new product
func (ct *ProductController) Create(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product products.ProductDTO
		// Decode the request body into the product struct
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error creating product: Invalid input data",
			})
			return
		}

		// Map DTO to model
		productToInsert, err := product.ToModel(validate)
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error creating product: Invalid input data",
			})
			return
		}
		// Validate the new product's fields
		if err := productToInsert.Validate(validate); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error creating product: Invalid input data",
			})
			return
		}

		// Create the product
		insertedProduct, err := ct.sv.Create(&productToInsert)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.DuplicateError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error creating product: Invalid input data",
				})
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error creating product: Not Found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error creating product: Server ran into an error",
				})
			}
			return
		}

		// Map model to DTO
		insertedProductDTO := insertedProduct.ToDTO()

		// Return the created product as JSON
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    insertedProductDTO,
		})
	}
}

// Updates a product by its ID
func (ct *ProductController) Update(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		parsedID, err := parseID(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Error updating product: Invalid input",
			})
			return
		}

		// Use json.Decoder to decode directly into the ProductPatch struct
		var patch products.ProductPatch
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&patch); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Error updating product: Invalid input data",
			})
			return
		}
		// Validate the patch fields
		if err := patch.Validate(validate); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error updating product: Invalid input data",
			})
			return
		}

		// Update the product
		updatedProduct, err := ct.sv.Update(parsedID, patch)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.ValidationError{}):
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": "Error updating product: Invalid input data",
				})
			case errors.Is(err, &customErrors.DuplicateError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error updating product: Invalid input data",
				})
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error updating product: Not found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error updating product: Server ran into an error",
				})
			}
			return
		}

		// Map model to DTO
		updatedProductDTO := updatedProduct.ToDTO()

		// Return the product as JSON
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    updatedProductDTO,
		})
	}
}

// Deletes a product by its ID
func (ct *ProductController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		parsedID, err := parseID(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Error deleting product: Invalid input data",
			})
			return
		}

		err = ct.sv.Delete(parsedID)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error deleting product: Product not found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error deleting product: Server ran into an error",
				})
			}
			return
		}
		// Return a 204 No Content status
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "success",
		})
	}
}

func parseID(id string) (parsedID models.ID, err error) {
	// Parse the ID into an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.ID{ID: 0}, err
	}
	parsedID = models.ID{ID: idInt}

	validate := validator.New()
	parsedID.Validate(validate)
	return
}
