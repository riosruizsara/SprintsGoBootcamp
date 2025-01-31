package sections

import (
	"encoding/json"

	"os"

	section "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sections"
)

// NewSectionJSONFile crea una nueva instancia de SectionJSONFile
func NewSectionJSONFile(path string) *SectionJSONFile {
	return &SectionJSONFile{
		path: path,
	}
}

// SectionJSONFile es una estructura que implementa la interfaz SectionLoader
type SectionJSONFile struct {
	// path es la ruta del archivo JSON que contiene las secciones
	path string
}

// Load carga las secciones desde un archivo JSON y devuelve un mapa de secciones
func (l *SectionJSONFile) Load() (map[int]section.Section, error) {
	// abrir archivo
	file, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decodificar archivo JSON
	var sectionsJSON []section.SectionDTO
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return nil, err
	}

	// serializar sections
	sectionsMap := make(map[int]section.Section)
	for _, sectionDTO := range sectionsJSON {
		sectionModel, err := sectionDTO.ToModel()
		if err != nil {
			// registrar el error pero continuar con otras secciones
			continue
		}
		sectionsMap[sectionModel.ID] = sectionModel
	}

	return sectionsMap, nil
}
