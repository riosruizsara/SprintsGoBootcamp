package sections

import (
	"sort"

	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	section "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sections"
)

// NewSectionMap crea una nueva instancia de SectionMap con una base de datos en memoria
func NewSectionMap(db map[int]section.Section) *SectionMap {
	// Base de datos por defecto en memoria
	defaultDb := make(map[int]section.Section)
	if db != nil {
		defaultDb = db
	}
	return &SectionMap{db: defaultDb}
}

// SectionMap representa un repositorio de secciones con implementación en memoria
type SectionMap struct {
	// db es un mapa de secciones
	db map[int]section.Section
}

// GetAll devuelve todas las secciones almacenadas en memoria, ordenadas por ID
func (sm *SectionMap) GetAll() ([]section.Section, error) {
	sections := make([]section.Section, 0, len(sm.db))
	for _, sec := range sm.db {
		sections = append(sections, sec)
	}

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].ID < sections[j].ID
	})

	return sections, nil
}

// GetByID devuelve una sección por ID
func (sm *SectionMap) GetByID(id models.ID) (sec section.Section, err error) {
	sec, ok := sm.db[id.ID]
	if !ok {
		err = &customErrors.NotFoundError{Message: "repository: Section does not exist"}
		return section.Section{}, err
	}
	return sec, nil
}

// Create crea una nueva sección en memoria
func (sm *SectionMap) Create(sec section.Section) error {
	// Verificar si el section_number ya existe
	for _, existingSec := range sm.db {
		if existingSec.SectionNumber == sec.SectionNumber {
			return &customErrors.ConflictError{Message: "repository: Section already exists"}
		}
	}

	// Generar un ID único si no se proporciona
	if sec.ID == 0 {
		sec.ID = len(sm.db) + 1
	}

	// Insertar la nueva sección
	sm.db[sec.ID] = sec
	return nil
}

// Delete elimina una sección por su ID
func (sm *SectionMap) Delete(id int) error {
	// Verificar si la sección existe
	if _, exists := sm.db[id]; !exists {
		return &customErrors.NotFoundError{Message: "repository: Section does not exist"}
	}

	// Eliminar la sección
	delete(sm.db, id)
	return nil
}

// Update actualiza una sección por su ID
func (sm *SectionMap) Update(id int, updatedSection section.Section) (section.Section, error) {
	// Verificar si la sección existe
	if _, exists := sm.db[id]; !exists {
		return section.Section{}, &customErrors.NotFoundError{Message: "repository: Section already exists"}
	}

	// Actualizar la sección
	updatedSection.ID = id // Asegurar que el ID no cambie
	sm.db[id] = updatedSection

	// Devolver la sección actualizada
	return updatedSection, nil
}
