package section

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type Section struct {
	ID                 int     `json:"id"`                  // Identificador único del sector
	SectionNumber      int     `json:"section_number"`      // Un código de sector
	CurrentTemperature float64 `json:"current_temperature"` // Una temperatura actual
	MinimumTemperature float64 `json:"minimum_temperature"` // Temperatura mínima permitida
	CurrentCapacity    int     `json:"current_capacity"`    // Capacidad actual
	MinimumCapacity    int     `json:"minimum_capacity"`    // Capacidad mínima permitida
	MaximumCapacity    int     `json:"maximum_capacity"`    // Capacidad máxima permitida
	WarehouseID        int     `json:"warehouse_id"`        // Un depósito (warehouse) asociado
	ProductTypeID      int     `json:"product_type_id"`     // Un tipo de producto (ProductType) asociado
}

type SectionPatch struct {
	SectionNumber      *int     `json:"section_number" validate:"omitempty,numeric,min=1"`
	CurrentTemperature *float64 `json:"current_temperature" validate:"omitempty"`
	MinimumTemperature *float64 `json:"minimum_temperature" validate:"omitempty"`
	CurrentCapacity    *int     `json:"current_capacity" validate:"omitempty,numeric,min=0"`
	MinimumCapacity    *int     `json:"minimum_capacity" validate:"omitempty,numeric,min=0"`
	MaximumCapacity    *int     `json:"maximum_capacity" validate:"omitempty,numeric,min=0"`
	WarehouseID        *int     `json:"warehouse_id" validate:"omitempty,numeric,min=1"`
	ProductTypeID      *int     `json:"product_type_id" validate:"omitempty,numeric,min=1"`
}

func (s *SectionPatch) Validate(validate *validator.Validate) error {
	return validate.Struct(s)
}

func NewSection(id, sectionNumber int, currentTemperature, minimumTemperature float64, currentCapacity, minimumCapacity, maximumCapacity, warehouseID, productTypeID int) (Section, error) {
	if sectionNumber <= 0 {
		return Section{}, &errors.ValidationError{Message: "Section_number must be greater than 0"}
	}
	if minimumTemperature >= currentTemperature {
		return Section{}, &errors.ValidationError{Message: "currentTemperature must be higher than the minimum temperature"}
	}
	if minimumCapacity > currentCapacity {
		return Section{}, &errors.ValidationError{Message: "currentCapacity must be greater than the minimumCapacity"}
	}
	if currentCapacity > maximumCapacity {
		return Section{}, &errors.ValidationError{Message: "currentCapacity should not exceed the maximumCapacity"}
	}

	return Section{
		ID:                 id,
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}, nil
}
