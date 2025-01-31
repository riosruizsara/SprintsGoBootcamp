package section

type SectionDTO struct {
	ID                 int     `json:"id,omitempty"`
	SectionNumber      int     `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

// Convierte SecciónDTO en modelo de Sección
func (dto *SectionDTO) ToModel() (Section, error) {
	return NewSection(
		dto.ID,
		dto.SectionNumber,
		dto.CurrentTemperature,
		dto.MinimumTemperature,
		dto.CurrentCapacity,
		dto.MinimumCapacity,
		dto.MaximumCapacity,
		dto.WarehouseID,
		dto.ProductTypeID,
	)
}

// Convierte el modelo de sección en SecciónDTO
func (s Section) ToDTO() SectionDTO {
	return SectionDTO{
		ID:                 s.ID,
		SectionNumber:      s.SectionNumber,
		CurrentTemperature: s.CurrentTemperature,
		MinimumTemperature: s.MinimumTemperature,
		CurrentCapacity:    s.CurrentCapacity,
		MinimumCapacity:    s.MinimumCapacity,
		MaximumCapacity:    s.MaximumCapacity,
		WarehouseID:        s.WarehouseID,
		ProductTypeID:      s.ProductTypeID,
	}
}

type SectionResponse struct {
	Data []SectionDTO `json:"data"`
}

// Convertir lista de modelos a respuesta JSON
func ToSectionResponse(sections []Section) SectionResponse {
	var dtos []SectionDTO
	for _, s := range sections {
		dtos = append(dtos, s.ToDTO())
	}
	return SectionResponse{Data: dtos}
}
