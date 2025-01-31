package buyers

import (
	"encoding/json"
	"errors"
	"net/http"

	sbuy "github.com/almarino_meli/grupo-5-wave-15/internal/service/buyers"
	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/buyer"
	"github.com/go-playground/validator/v10"

	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// Estructura instancia
type BuyerDefault struct {
	sv sbuy.BuyerService
}

// Instancia handler
func NewBuyerDefault(sv sbuy.BuyerService) *BuyerDefault {
	return &BuyerDefault{sv: sv}
}

// Controladores

// GET - Obtener todos los buyers
func (h *BuyerDefault) GetAllBuyers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buyers, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Internal Server Error",
			})
			return
		}
		var finalData []buyer.BuyerDoc
		for _, value := range buyers {
			finalData = append(finalData, value.MapBuyerToDto())
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    finalData,
		})
	}
}

// GET - Obtener buyer por Id
func (h *BuyerDefault) GetBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}
		bySaved, errD := h.sv.FindById(idInt)
		if errD != nil {
			switch {
			case errors.Is(errD, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error retrieving buyer: Not Found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error retrieving buyer: Internal Server Error",
				})
			}
			return
		}
		finalData := bySaved.MapBuyerToDto()
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    finalData,
		})
	}
}

// POST - Crear un nuevo buyer
func (h *BuyerDefault) PostBuyer(validHelper *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buyerDoc buyer.BuyerDoc
		if err := json.NewDecoder(r.Body).Decode(&buyerDoc); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		if err := buyerDoc.Validate(validHelper); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Error on creation buyer: Invalid input data")
			return
		}

		bySaved, errD := h.sv.SetNewBuyer(buyerDoc.MapBuyerToModel())

		if errD != nil {
			switch {
			case errors.Is(errD, &customErrors.DuplicateError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error on creation buyer: Invalid Request",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error on creation buyer: Internal Server Error",
				})
			}
			return
		}

		finalData := bySaved.MapBuyerToDto()
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    finalData,
		})
	}
}

// PATCH - Modificar atributos en el buyer
func (h *BuyerDefault) PatchBuyer(validHelper *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}
		var buyerDoc buyer.BuyerDocPatched
		if err := json.NewDecoder(r.Body).Decode(&buyerDoc); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		if err := buyerDoc.Validate(validHelper); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "Error on patching buyer: Invalid input data")
			return
		}

		d, errD := h.sv.UpdateBuyer(idInt, buyerDoc)
		if errD != nil {
			switch {
			case errors.Is(errD, &customErrors.DuplicateError{}):
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": "Error on patching buyer: Invalid Request",
				})
			case errors.Is(errD, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error on patching buyer: Not Found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error on patching buyer: Internal Server Error",
				})
			}
			return
		}

		finalData := d.MapBuyerToDto()
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    finalData,
		})
	}
}

// DELETE - Eliminar un buyer
func (h *BuyerDefault) DeleteBuyer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}
		errD := h.sv.DeleteBuyer(idInt)
		if errD != nil {
			switch {
			case errors.Is(errD, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error deleting buyer: Not Found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Error deleting buyer: Internal Server Error",
				})
			}
			return
		}
		response.JSON(w, http.StatusNoContent, nil)
	}
}
