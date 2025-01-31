package warehouses

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/bootcamp-go/web/response"

	warehouse1 "github.com/almarino_meli/grupo-5-wave-15/internal/service/warehouses"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/warehouse"
	"github.com/go-chi/chi/v5"
)

// NewWarehouseController creates a new instance of WarehouseController, and associates the passed service with it
func NewWarehouseController(sv warehouse1.WarehouseService) *WarehouseController {
	return &WarehouseController{sv: sv}
}

type WarehouseController struct {
	sv warehouse1.WarehouseService
}

// Estructura para respuesta de múltiples warehouses
type WarehousesResponse struct {
	Data []warehouse.WarehouseDTO `json:"data"`
}

// Estructura para respuesta de un único warehouse
type WarehouseResponse struct {
	Data warehouse.WarehouseDTO `json:"data"`
}

// GetAll returns all the found warehouses
func (ct *WarehouseController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := ct.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Internal Server Error",
			})
			return
		}

		// Transformar los warehouses al formato DTO
		warehousesDTO := make([]warehouse.WarehouseDTO, 0, len(warehouses))
		for _, wh := range warehouses {
			warehousesDTO = append(warehousesDTO, wh.ToDTO())
		}

		// Ordenar los DTOs por ID
		sort.Slice(warehousesDTO, func(i, j int) bool {
			return *warehousesDTO[i].ID < *warehousesDTO[j].ID
		})

		response := WarehousesResponse{Data: warehousesDTO}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// GetByID returns a single warehouse by its ID
func (ct *WarehouseController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener ID desde URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		// Obtener warehouse del servicio
		warehouse, err := ct.sv.GetByID(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "Warehouse not found",
			})
			return
		}

		// Crear respuesta usando ToDTO
		response := WarehouseResponse{
			Data: warehouse.ToDTO(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// Create crea un nuevo warehouse
func (ct *WarehouseController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto warehouse.WarehouseDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "Datos inválidos", http.StatusUnprocessableEntity)
			return
		}

		newWarehouse, err := ct.sv.Create(dto)
		if err != nil {
			switch err.(type) {
			case *errors.DuplicateError:
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error creating warehouse: Conflict",
				})
			default:
				response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
					"message": "Error creating warehouse: Invalid input data",
				})
			}
			return
		}

		response := WarehouseResponse{Data: newWarehouse.ToDTO()}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func (ct *WarehouseController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener ID del warehouse
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Decodificar el body
		var dto warehouse.WarehouseDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error creating warehouse: Invalid input data",
			})
			return
		}

		// Validar que el dto no esté vacío
		if dto == (warehouse.WarehouseDTO{}) {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		// Actualizar el warehouse
		updatedWarehouse, err := ct.sv.Update(id, dto)
		if err != nil {
			switch err.(type) {
			case *errors.NotFoundError:
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Warehouse not found",
				})
			case *errors.DuplicateError:
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error creating warehouse: Conflict",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}

		// Crear respuesta usando ToDTO
		response := WarehouseResponse{
			Data: updatedWarehouse.ToDTO(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (ct *WarehouseController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener ID del parametro URL
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		// Eliminar warehouse
		err = ct.sv.Delete(idInt)
		if err != nil {
			switch err.(type) {
			case *errors.NotFoundError:
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Section not found",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
