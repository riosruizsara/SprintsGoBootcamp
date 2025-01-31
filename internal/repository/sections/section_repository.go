package sections

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	section "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sections"
)

// SectionRepository  interfaz que define los métodos para interactuar con los sectores
type SectionRepository interface {
	// GetAll  método que retorna un mapa de todas las secciones
	GetAll() ([]section.Section, error)
	// GetByID  método que retorna una sección por ID
	GetByID(id models.ID) (sec section.Section, err error)
	// Create  método que crea una nueva sección
	Create(sec section.Section) error
	// Delete  método que elimina una sección por ID
	Delete(id int) error
	// Update  método que actualiza una sección por ID
	Update(id int, updatedSection section.Section) (section.Section, error)
}
