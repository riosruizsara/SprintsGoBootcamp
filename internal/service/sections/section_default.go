package sections

import (
	sectionRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/sections"

	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	models "github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	section "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sections"
	"github.com/go-playground/validator/v10"
)

// SectionServiceDefault implementa las operaciones sobre las secciones
type SectionServiceDefault struct {
	rp sectionRp.SectionRepository
}

// NewSectionServiceDefault crea una nueva instancia del servicio de secciones
func NewSectionServiceDefault(rp sectionRp.SectionRepository) SectionServiceDefault {
	return SectionServiceDefault{rp: rp}
}

// GetAllSections obtiene todas las secciones
func (ss SectionServiceDefault) GetAll() ([]section.Section, error) {
	sections, err := ss.rp.GetAll()
	if err != nil {
		return nil, err
	}
	if len(sections) == 0 {
		return nil, &customErrors.NotFoundError{Message: "No sections were found"}
	}
	return sections, nil
}

// GetByID devuelve una sección por ID
func (ss SectionServiceDefault) GetByID(id models.ID) (sec section.Section, err error) {
	sec, err = ss.rp.GetByID(id)
	if err != nil {
		return section.Section{}, err
	}
	if sec.ID == 0 {
		return section.Section{}, &customErrors.NotFoundError{Message: "Section not found"}
	}
	return sec, nil
}

// Create crea una nueva sección
func (ss SectionServiceDefault) Create(sec section.Section) error {
	// Validar campos requeridos
	if sec.SectionNumber == 0 || sec.WarehouseID == 0 || sec.ProductTypeID == 0 {
		return &customErrors.ValidationError{Message: "required fields are missing"}
	}

	// Crear la sección en el repositorio
	return ss.rp.Create(sec)
}

// Delete elimina una sección por su ID
func (ss SectionServiceDefault) Delete(id int) error {
	// Eliminar la sección en el repositorio
	return ss.rp.Delete(id)
}

// Update actualiza una sección por su ID
func (ss SectionServiceDefault) Update(id int, sectionPatch section.SectionPatch) (section.Section, error) {
	// Validar el SectionPatch usando el validador
	validate := validator.New()
	if err := sectionPatch.Validate(validate); err != nil {
		return section.Section{}, &customErrors.ValidationError{Message: err.Error()}
	}
	// Convertir id a models.ID si es necesario
	modelID := models.ID{ID: id}

	// Obtener la sección existente
	existingSection, err := ss.rp.GetByID(modelID)
	if err != nil {
		return section.Section{}, err // Propagamos el error (por ejemplo, NotFoundError)
	}

	// Aplicar solo los valores enviados (no nil)
	if sectionPatch.SectionNumber != nil {
		existingSection.SectionNumber = *sectionPatch.SectionNumber
	}
	if sectionPatch.CurrentTemperature != nil {
		existingSection.CurrentTemperature = *sectionPatch.CurrentTemperature
	}
	if sectionPatch.MinimumTemperature != nil {
		existingSection.MinimumTemperature = *sectionPatch.MinimumTemperature
	}
	if sectionPatch.CurrentCapacity != nil {
		existingSection.CurrentCapacity = *sectionPatch.CurrentCapacity
	}
	if sectionPatch.MinimumCapacity != nil {
		existingSection.MinimumCapacity = *sectionPatch.MinimumCapacity
	}
	if sectionPatch.MaximumCapacity != nil {
		existingSection.MaximumCapacity = *sectionPatch.MaximumCapacity
	}
	if sectionPatch.WarehouseID != nil {
		existingSection.WarehouseID = *sectionPatch.WarehouseID
	}
	if sectionPatch.ProductTypeID != nil {
		existingSection.ProductTypeID = *sectionPatch.ProductTypeID
	}

	// Validar la sección actualizada
	if _, err := section.NewSection(
		existingSection.ID,
		existingSection.SectionNumber,
		existingSection.CurrentTemperature,
		existingSection.MinimumTemperature,
		existingSection.CurrentCapacity,
		existingSection.MinimumCapacity,
		existingSection.MaximumCapacity,
		existingSection.WarehouseID,
		existingSection.ProductTypeID,
	); err != nil {
		return section.Section{}, err // Propagamos el error de validación
	}

	// Guardar cambios en la BD
	return ss.rp.Update(id, existingSection)
}
