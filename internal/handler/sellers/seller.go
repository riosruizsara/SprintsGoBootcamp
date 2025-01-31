package sellers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	sellersSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/sellers"
	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/sellers"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewSellerDefault(sv sellersSv.SellerService) *SellerDefault {
	return &SellerDefault{sv: sv}
}

type SellerDefault struct {
	sv sellersSv.SellerService
}

func (h *SellerDefault) CreateSeller(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sDTO sellers.SellerDTO
		// Valida el decode
		err := json.NewDecoder(r.Body).Decode(&sDTO)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Error creating seller: Invalid input data",
			})
			return
		}
		// Map DTO a ValueObject
		s, err := sellers.MapDTOToValueObject(sDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Chequeo el ValueObject con la libreria
		if err := s.Validate(validate); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error updating seller: Invalid input data",
			})
			return
		}
		// llamo al service con el value object correcto
		res, err := h.sv.Create(*s)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.ConflictError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error creating seller: Conflict",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}
		var createdSeller sellers.SellerDTO = sellers.MapToSellerDTO(res)
		response.JSON(w, http.StatusCreated, map[string]any{"message": "success", "data": createdSeller})
		return
	}
}

func (h *SellerDefault) GetAllSellers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Internal Server Error",
			})
			return
		}
		// Crear slice de SellerDTO
		sellersDTO := make([]sellers.SellerDTO, 0, len(res))

		// Iterar y mapear cada seller
		for _, seller := range res {
			sellerDTO := sellers.MapToSellerDTO(seller)
			sellersDTO = append(sellersDTO, sellerDTO)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    sellersDTO,
		})
		return
	}
}

func (h *SellerDefault) GetSellerById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}
		seller, err := h.sv.GetById(id)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Seller Not Found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}

		sellerDTO := sellers.MapToSellerDTO(seller)
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    sellerDTO,
		})
		return
	}
}

func (h *SellerDefault) UpdateSeller(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}
		var sDTO sellers.SellerPatch
		err = json.NewDecoder(r.Body).Decode(&sDTO)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Error updating seller: Invalid input data",
			})
			return
		}
		if err := sDTO.Validate(validate); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error creating seller: Invalid input data",
			})
			return
		}
		res, err := h.sv.Update(sDTO, id)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.ConflictError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error updating seller: Conflict",
				})
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Error updating seller: Not Found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}
		var updatedSeller sellers.SellerDTO = sellers.MapToSellerDTO(res)
		response.JSON(w, http.StatusOK, map[string]any{"message": "success", "data": updatedSeller})
		return
	}
}

func (h *SellerDefault) DeleteSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}
		err = h.sv.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Seller Not Found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "success",
		})
		return
	}
}
