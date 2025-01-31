package sections

import section "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sections"

// SectionLoader es una interfaz que representa el cargador para los sectores (Sections).
type SectionLoader interface {
	// Load es un método que carga los sectores.

	Load() (v map[int]section.Section, err error)
}
