package sections

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	section "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sections"
)

// SectionService es una interfaz que representa el servicio para las secciones
type SectionService interface {
	// GetAll es un método que devuelve un mapa de todas las secciones
	GetAll() ([]section.Section, error)
	// GetByID es un método que devuelve una sección por ID
	GetByID(id models.ID) (sec section.Section, err error)
	// Create es un método que crea una nueva sección
	Create(sec section.Section) error
	// Delete es un método que elimina una sección por ID
	Delete(id int) error
	// Update es un método que actualiza una sección por ID
	Update(id int, sectionDTO section.SectionPatch) (section.Section, error)
}
