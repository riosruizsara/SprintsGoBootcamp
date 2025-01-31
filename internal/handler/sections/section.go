package sections

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	sectionSv "github.com/almarino_meli/grupo-5-wave-15/internal/service/sections"
	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	section "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sections"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// NewSectionController crea una nueva instancia del controlador de secciones
func NewSectionController(sv sectionSv.SectionService) *SectionController {
	return &SectionController{sv: sv}
}

// SectionController maneja las peticiones HTTP relacionadas con secciones
type SectionController struct {
	sv sectionSv.SectionService
}

// GetAll obtiene todas las secciones
func (ct *SectionController) GetAllSections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener las secciones del servicio
		sections, err := ct.sv.GetAll()
		if err != nil {
			if errors.Is(err, &customErrors.NotFoundError{}) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Not Found",
				})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Internal Server Error",
			})
			return
		}

		// Convertir las secciones a DTOs y formatear la respuesta
		sectionResponse := section.ToSectionResponse(sections)

		// Retornar la respuesta en formato JSON
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    sectionResponse.Data,
			"total":   len(sectionResponse.Data),
		})
	}
}

// GetByID obtiene una sección por su ID
func (ct *SectionController) GetSectionsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// Parse the ID into an integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		// Obtener la sección por ID
		section, err := ct.sv.GetByID(models.ID{ID: idInt})
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
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
		// Convertir el modelo a DTO
		foundSectionDTO := section.ToDTO()

		// Retornar la sección en formato JSON
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    foundSectionDTO,
		})
	}
}

// Create crea una nueva sección
func (sc *SectionController) CreateSections() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sectionDTO section.SectionDTO

		// Decodificar el cuerpo de la solicitud en la estructura SectionDTO
		if err := json.NewDecoder(r.Body).Decode(&sectionDTO); err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error creating section: Invalid input data",
			})
			return
		}

		// Mapear DTO a modelo
		sectionModel, err := sectionDTO.ToModel()
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
				"message": "Error creating section: Invalid input data",
			})
			return
		}

		// Crear la sección en el servicio
		if err := sc.sv.Create(sectionModel); err != nil {
			switch {
			case errors.Is(err, &customErrors.ConflictError{}):
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": "Error creating section: Conflict",
				})
			case errors.Is(err, &customErrors.ValidationError{}):
				response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
					"message": "Error creating section: Invalid input data",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}

		// Convertir el modelo a DTO para la respuesta
		createdSectionDTO := sectionModel.ToDTO()

		// Retornar la sección creada en formato JSON
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": http.StatusText(http.StatusCreated),
			"data":    createdSectionDTO,
		})
	}
}

// Delete elimina una sección por su ID
func (sc *SectionController) DeleteSection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID de la URL
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		// Eliminar la sección usando el servicio
		if err := sc.sv.Delete(id); err != nil {
			if errors.Is(err, &customErrors.NotFoundError{}) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Section not found",
				})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Internal Server Error",
			})
			return
		}

		// Retornar una respuesta vacía con código 204 (No Content)
		w.WriteHeader(http.StatusNoContent)
	}
}

// Update actualiza una sección por su ID
func (sc *SectionController) UpdateSection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID de la URL
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad Request: Invalid input data",
			})
			return
		}

		// Decodificar el cuerpo de la solicitud en un SectionPatch
		var sectionPatch section.SectionPatch
		if err := json.NewDecoder(r.Body).Decode(&sectionPatch); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Error updating section: Invalid input data",
			})
			return
		}

		// Actualizar la sección usando el servicio
		updatedSection, err := sc.sv.Update(id, sectionPatch)
		if err != nil {
			switch {
			case errors.Is(err, &customErrors.NotFoundError{}):
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": "Section not found",
				})
			case errors.Is(err, &customErrors.ValidationError{}):
				response.JSON(w, http.StatusUnprocessableEntity, map[string]any{
					"message": "Error updating section: Invalid input data",
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "Internal Server Error",
				})
			}
			return
		}

		// Convertir el modelo actualizado a DTO para la respuesta
		updatedSectionDTO := updatedSection.ToDTO()

		// Retornar la sección actualizada en formato JSON
		response.JSON(w, http.StatusOK, map[string]any{
			"data": updatedSectionDTO,
		})
	}
}
